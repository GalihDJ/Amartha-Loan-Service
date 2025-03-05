package models

import "time"

type LoanState string

// constants for loan reqeust states
const (
	StateProposed  LoanState = "PROPOSED"
	StateApproved  LoanState = "APPROVED"
	StateInvested  LoanState = "INVESTED"
	StateDisbursed LoanState = "DISBURSED"
)

// loan request model
type LoanRequest struct {
	LoanRequestID       string    `json:"loan_request_id"`
	BorrowerID          string    `json:"borrower_id"`
	PrincipalAmount     float64   `json:"principal_amount"`
	Rate                float64   `json:"rate"`
	ROI                 float64   `json:"roi"`
	State               LoanState `json:"state"`
	CreatedDate         time.Time `json:"created_date" format:"date-time"`
	AgreementLetterLink string    `json:"agreement_letter_link"`
	// LoanApproval        *LoanApproval     `json:"loan_approval"`
	// LoanInvestments     []LoanInvestment  `json:"loan_investments"`
	// LoanDisbursement    *LoanDisbursement `json:"loan_disbursement"`
}

// loan request model for swagger
type LoanRequestSwagger struct {
	BorrowerID      string    `json:"borrower_id"`
	PrincipalAmount float64   `json:"principal_amount"`
	Rate            float64   `json:"rate"`
	ROI             float64   `json:"roi"`
	State           LoanState `json:"state"`
}

// loan request approval model
type LoanApproval struct {
	LoanApprovalID      string    `json:"loan_approval_id"`
	LoanRequestID       string    `json:"loan_request_id"`
	FieldValidatorProof string    `json:"field_validator_proof" default:"https://picsum.photos/200"`
	EmployeeID          string    `json:"employee_id"`
	ApprovedDate        time.Time `json:"approved_date"`
}

// loan request approval model for swagger
type LoanApprovalSwagger struct {
	LoanRequestID       string `json:"loan_request_id"`
	FieldValidatorProof string `json:"field_validator_proof" default:"https://picsum.photos/200"`
	EmployeeID          string `json:"employee_id"`
}

// add investment for a loan request
type LoanInvestment struct {
	InvestmentID  string    `json:"investment_id"`
	LoanRequestID string    `json:"loan_request_id"`
	InvestorID    string    `json:"investor_id"`
	Amount        float64   `json:"amount"`
	CreatedDate   time.Time `json:"created_date" format:"date-time"`
}

// loan investment model for swagger
type LoanInvestmentSwagger struct {
	LoanRequestID string  `json:"loan_request_id"`
	InvestorID    string  `json:"investor_id"`
	Amount        float64 `json:"amount"`
}

// disbursement of loan
type LoanDisbursement struct {
	DisbursementID     string    `json:"disbursement_id"`
	LoanRequestID      string    `json:"loan_request_id"`
	AgreementLetterURL string    `json:"agreement_letter_url" default:"https://picsum.photos/200/300"`
	EmployeeID         string    `json:"employee_id"`
	DisbursementDate   time.Time `json:"disbursed_date"`
}

// disbursement of loan model for swagger
type LoanDisbursementSwagger struct {
	LoanRequestID      string `json:"loan_request_id"`
	AgreementLetterURL string `json:"agreement_letter_url" default:"https://picsum.photos/200/300"`
	EmployeeID         string `json:"employee_id"`
}

type Investor struct {
	InvestorID    string    `json:"investor_id"`
	InvestorName  string    `json:"investor_name"`
	InvestorEmail string    `json:"investor_email"`
	CreatedDate   time.Time `json:"created_date" format:"date-time"`
}

type InvestorSwagger struct {
	InvestorName  string `json:"investor_name"`
	InvestorEmail string `json:"investor_email"`
}
