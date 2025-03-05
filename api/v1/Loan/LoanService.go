package loan

import (
	"amartha-loan-service/models"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type LoanService struct {
	repo ILoanRepository
}

func NewLoanService(repo ILoanRepository) ILoanService {
	return &LoanService{
		repo: repo,
	}
}

type ILoanService interface {
	CreateLoanRequest(loanRequest *models.LoanRequest) (*models.LoanRequest, error)
	GetLoanRequestById(loanRequestID string) (*models.LoanRequest, error)
	ApproveLoanRequest(loanRequestID string, loanApproval *models.LoanApproval) error
	CreateLoanInvestment(loanRequestID string, loanInvestment *models.LoanInvestment) error
	CreateLoanDisbursement(loanRequestID string, loanDisbursement *models.LoanDisbursement) error
}

// CreateLoanRequest implements ILoanService.
func (ls *LoanService) CreateLoanRequest(loanRequest *models.LoanRequest) (*models.LoanRequest, error) {

	// define loan request id
	loanRequest.LoanRequestID = uuid.New().String()

	// define loan request state
	loanRequest.State = models.StateProposed

	// define created date
	loanRequest.CreatedDate = time.Now()

	_, err := ls.repo.CreateLoanRequest(loanRequest)
	if err != nil {
		return nil, err
	}

	return loanRequest, nil
}

// GetLoanRequestById implements ILoanService.
func (ls *LoanService) GetLoanRequestById(loanRequestID string) (*models.LoanRequest, error) {

	// call GetLoanRequestById
	loanRequest, err := ls.repo.GetLoanRequestById(loanRequestID)

	if err != nil {
		log.Println("Error: ", err)
		return nil, err
	}

	return loanRequest, err
}

// ApproveLoanRequest implements ILoanService.
func (ls *LoanService) ApproveLoanRequest(loanRequestID string, loanApproval *models.LoanApproval) error {

	// get the loan from the repository
	loanRequest, err := ls.repo.GetLoanRequestById(loanRequestID)
	if err != nil {
		return err
	}

	// check loan request state
	if loanRequest.State != models.StateProposed {
		return errors.New("only loan request with PROPOSED status can be approved")
	}

	// define loan approval id
	loanApproval.LoanApprovalID = uuid.New().String()

	// define current date for approval date
	loanApproval.ApprovedDate = time.Now()

	// Save the approval and update the loan state
	err = ls.repo.ApproveLoan(loanRequestID, loanApproval)
	if err != nil {
		return err
	}

	return nil
}

// LoanInvestment implements ILoanService.
func (ls *LoanService) CreateLoanInvestment(loanRequestID string, loanInvestment *models.LoanInvestment) error {

	// get the loan request from the repository
	loanRequest, err := ls.repo.GetLoanRequestById(loanRequestID)
	if err != nil {
		return err
	}

	// check loan request state
	if loanRequest.State != models.StateApproved {
		return errors.New("investments can only be added to APPROVED loan request")
	}

	// get investments on loan
	loanInvestments, err := ls.repo.GetLoanInvestments(loanRequestID)
	if err != nil {
		return err
	}

	// get total of investment
	var totalInvestment float64
	for _, loanInvestment := range loanInvestments {
		totalInvestment += loanInvestment.Amount
	}

	// compare total investment and investment amount with principal amount
	if totalInvestment+loanInvestment.Amount > loanRequest.PrincipalAmount {
		return errors.New("total investment exceeds the loan principal")
	}

	// define loan investment id
	loanInvestment.InvestmentID = uuid.New().String()

	// define current date for approval date
	loanInvestment.CreatedDate = time.Now()

	// create loan investment
	_, err = ls.repo.CreateLoanInvestment(loanInvestment)
	if err != nil {
		return err
	}

	// check if total investment and investment amount is equal to principal amount
	if totalInvestment+loanInvestment.Amount == loanRequest.PrincipalAmount {

		// update loan request to INVESTED
		err = ls.repo.UpdateLoanRequestToInvested(loanRequestID)
		if err != nil {
			return err
		}

		// obtain investor IDs from investments data
		loanInvestments, err := ls.repo.GetLoanInvestments(loanRequestID)
		if err != nil {
			return err
		}

		// get investor emails

		// define array to store emails
		var investorEmails []string

		// loop through loanInvestments
		for index := range loanInvestments {

			// call GetInvestorEmail repo to get investor email
			investorEmail, err := ls.repo.GetInvestorEmail(loanInvestments[index].InvestorID)
			if err != nil {
				log.Println("Error: ", err)
				return err
			}

			// append email to array
			investorEmails = append(investorEmails, investorEmail)
		}

		// mock send email
		MockSendEmail(investorEmails)
	}

	return nil
}

// CreateLoanDisbursement implements ILoanService.
func (ls *LoanService) CreateLoanDisbursement(loanRequestID string, loanDisbursement *models.LoanDisbursement) error {

	// define loan request id
	loanDisbursement.DisbursementID = uuid.New().String()

	// define created date
	loanDisbursement.DisbursementDate = time.Now()

	// call CreateLoanDisbursement repo
	err := ls.repo.CreateLoanDisbursement(loanRequestID, loanDisbursement)
	if err != nil {
		return err
	}

	return nil
}

// mock function to send email

func MockSendEmail(recipients []string) {

	const bodyEmail = "Here is the link to the loan agreement letter: https://picsum.photos/200/300"

	for _, email := range recipients {
		fmt.Println("---------------------------")
		fmt.Printf("Sending email to: %s\n", email)
		fmt.Printf("Email Body: %s\n", bodyEmail)
		fmt.Println("---------------------------")
	}
}
