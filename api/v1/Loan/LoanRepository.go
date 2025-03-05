package loan

import (
	"amartha-loan-service/models"
	utilsConn "amartha-loan-service/utils/Connections"
	"database/sql"
	"fmt"
	"log"
)

type LoanRepository struct {
	connPSQL *utilsConn.ConnectionPSQL
}

func NewLoanRepository(conn_arg *utilsConn.ConnectionPSQL) ILoanRepository {
	return &LoanRepository{
		connPSQL: conn_arg,
	}
}

type ILoanRepository interface {
	CreateLoanRequest(loanRequest *models.LoanRequest) (int, error)
	GetLoanRequestById(loanRequestID string) (*models.LoanRequest, error)
	ApproveLoan(loanRequestID string, loanApproval *models.LoanApproval) error
	CreateLoanInvestment(loanInvestment *models.LoanInvestment) (int, error)

	GetLoanInvestments(loanRequestID string) ([]models.LoanInvestment, error)

	CreateInvestor(investor *models.Investor) (int, error)
}

// CreateLoanRequest implements ILoanRepository.
func (lr *LoanRepository) CreateLoanRequest(loanRequest *models.LoanRequest) (int, error) {

	// connect to DB
	db, err := lr.connPSQL.ConnectionOpenPSQL()

	if err != nil {
		return 0, err
	}

	defer db.Close()

	// query to insert loan request
	query := "INSERT INTO loan_request (loan_request_id, borrower_id, principal_amount, rate, roi, state, created_date)"
	query += " VALUES ($1, $2, $3, $4, $5, $6, $7)"

	// execute query and get result
	result, err := db.Exec(query,
		loanRequest.LoanRequestID,
		loanRequest.BorrowerID,
		loanRequest.PrincipalAmount,
		loanRequest.Rate,
		loanRequest.ROI,
		loanRequest.State,
		loanRequest.CreatedDate,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("could not fetch rows affected: %v", err)
	}

	return int(rowsAffected), err
}

// GetLoanRequestById implements ILoanRepository.
func (lr *LoanRepository) GetLoanRequestById(loanRequestID string) (*models.LoanRequest, error) {

	// connect to DB
	db, err := lr.connPSQL.ConnectionOpenPSQL()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	// query to get loan request by id
	query := "SELECT loan_request_id, borrower_id, principal_amount, rate, roi, state FROM loan_request "
	query += "WHERE loan_request_id = $1"

	rows, err := db.Query(query, loanRequestID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// scan rows for queried columns
	var loanRequest models.LoanRequest
	for rows.Next() {
		if err := rows.Scan(
			&loanRequest.LoanRequestID,
			&loanRequest.BorrowerID,
			&loanRequest.PrincipalAmount,
			&loanRequest.Rate,
			&loanRequest.ROI,
			&loanRequest.State); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			log.Fatal(err)
		}
	}

	return &loanRequest, err
}

// ApproveLoan implements ILoanRepository.
func (lr *LoanRepository) ApproveLoan(loanRequestID string, loanApproval *models.LoanApproval) error {

	// connect to DB
	db, err := lr.connPSQL.ConnectionOpenPSQL()

	if err != nil {
		return err
	}

	defer db.Close()

	// start transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// update the loan request state

	updateLoanQuery := "UPDATE loan_request SET state = $1 WHERE loan_request_id = $2"
	fmt.Println("updateLoanQuery: ", updateLoanQuery)

	_, err = tx.Exec(updateLoanQuery, models.StateApproved, loanRequestID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// insert loan request approval details

	insertLoanApprovalQuery := "INSERT INTO loan_approval (loan_approval_id, loan_request_id, field_validator_proof, employee_id, approved_date) "
	insertLoanApprovalQuery += "VALUES ($1, $2, $3, $4, $5)"
	fmt.Println("insertLoanApprovalQuery: ", insertLoanApprovalQuery)

	_, err = tx.Exec(insertLoanApprovalQuery,
		loanApproval.LoanApprovalID, loanRequestID, loanApproval.FieldValidatorProof, loanApproval.EmployeeID, loanApproval.ApprovedDate)
	if err != nil {

		fmt.Println("ERR: ", err)
		tx.Rollback()
		return err
	}

	// commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return nil

}

// LoanInvestment implements ILoanRepository.
func (lr *LoanRepository) CreateLoanInvestment(loanInvestment *models.LoanInvestment) (int, error) {

	// connect to DB
	db, err := lr.connPSQL.ConnectionOpenPSQL()

	if err != nil {
		return 0, err
	}

	defer db.Close()

	// query to insert loan request
	query := "INSERT INTO investment_list (investment_id, loan_request_id, investor_id, amount, created_date)"
	query += " VALUES ($1, $2, $3, $4, $5)"

	// execute query and get result
	result, err := db.Exec(query,
		loanInvestment.InvestmentID,
		loanInvestment.LoanRequestID,
		loanInvestment.InvestorID,
		loanInvestment.Amount,
		loanInvestment.CreatedDate,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("could not fetch rows affected: %v", err)
	}

	return int(rowsAffected), err
}

// GetLoanInvestments implements ILoanRepository.
func (lr *LoanRepository) GetLoanInvestments(loanRequestID string) ([]models.LoanInvestment, error) {

	// connect to DB
	db, err := lr.connPSQL.ConnectionOpenPSQL()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	query := "SELECT il.investment_id, il.investor_id, il.amount FROM investment_list il "
	query += "INNER JOIN loan_request lr ON il.loan_request_id = lr.loan_request_id "
	query += "WHERE lr.loan_request_id = $1"

	// query := "SELECT cr.prompt_template FROM chatbot_role cr "
	// query += "INNER JOIN chatbot_sessions cs ON cr.chatbot_role_id = cs.chatbot_role_id "
	// query += "WHERE cs.session_id = $1"

	rows, err := db.Query(query, loanRequestID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var loanInvestments []models.LoanInvestment
	for rows.Next() {
		var loanInvestment models.LoanInvestment
		if err := rows.Scan(
			&loanInvestment.InvestmentID,
			&loanInvestment.InvestorID,
			&loanInvestment.Amount); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		loanInvestments = append(loanInvestments, loanInvestment)
	}

	return loanInvestments, err

}

// CreateInvestor implements ILoanRepository.
func (lr *LoanRepository) CreateInvestor(investor *models.Investor) (int, error) {

	db, err := lr.connPSQL.ConnectionOpenPSQL()

	if err != nil {
		return 0, err
	}

	defer db.Close()

	query := "INSERT INTO investor_list (investor_id, investor_name, investor_email, created_date)"
	query += " VALUES ($1, $2, $3, $4)"

	result, err := db.Exec(query,
		investor.InvestorID,
		investor.InvestorName,
		investor.InvestorEmail,
		investor.CreatedDate,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("could not fetch rows affected: %v", err)
	}

	return int(rowsAffected), err
}
