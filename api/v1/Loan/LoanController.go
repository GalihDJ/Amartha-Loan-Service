package loan

import (
	"net/http"

	"amartha-loan-service/models"

	"github.com/gin-gonic/gin"
)

type LoanController struct {
	service ILoanService
}

func NewLoanController(service ILoanService) *LoanController {
	return &LoanController{
		service: service,
	}
}

// CreateLoanRequest godoc
// @Summary Create new loan request
// @Description Create a new loan request
// @Tags Loan
// @Accept json
// @Produce json
// @Param models.LoanRequestSwagger body models.LoanRequestSwagger true "Loan Request parameters"
// @Success 200
// @Router /api/v1/loan [post]
func (lc *LoanController) CreateLoanRequest(c *gin.Context) {

	// define variable to store struct
	var loanRequest models.LoanRequest

	// bind value from user input
	if err := c.ShouldBindJSON(&loanRequest); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}

	// call CreateLoanRequest service
	createdLoanRequest, err := lc.service.CreateLoanRequest(&loanRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"loan_request": createdLoanRequest})
}

// GetLoanRequestById godoc
// @Summary Get loan request by id
// @Description Get loan request by id
// @Tags Loan
// @Produce json
// @Param loanRequestID path string true "Loan Reuqest ID"
// @Success 200
// @Router /api/v1/loan/{loanRequestID} [get]
func (lc *LoanController) GetLoanRequestById(c *gin.Context) {

	// bind parameter from input
	loanRequestID := c.Param("loanRequestID")

	// call GetLoanRequestById service
	loanRequest, err := lc.service.GetLoanRequestById(loanRequestID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if loanRequest == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "loan request not found"})
		return
	}

	c.JSON(http.StatusOK, loanRequest)
}

// ApproveLoanRequest godoc
// @Summary Approve loan request
// @Description Approve loan request
// @Tags Loan
// @Produce json
// @Param loanRequestID path string true "Loan Request ID"
// @Param models.LoanApprovalSwagger body models.LoanApprovalSwagger true "Loan approval"
// @Success 200
// @Router /api/v1/loan/{loanRequestID} [put]
func (lc *LoanController) ApproveLoanRequest(c *gin.Context) {

	// define variable to store struct
	var approveLoanRequest models.LoanApproval

	// bind parameter from input
	loanRequestID := c.Param("loanRequestID")

	// bind value from user input
	if err := c.ShouldBindJSON(&approveLoanRequest); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}

	// call ApproveLoanRequest service
	err := lc.service.ApproveLoanRequest(loanRequestID, &approveLoanRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error",
			"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Loan request approved succesfully",
	})
}

// CreateLoanInvestment godoc
// @Summary Create new loan investment
// @Description Create a new loan investment
// @Tags Loan
// @Accept json
// @Produce json
// @Param loanRequestID path string true "Loan Request ID"
// @Param models.LoanInvestmentSwagger body models.LoanInvestmentSwagger true "Loan Request parameters"
// @Success 200
// @Router /api/v1/loan/{loanRequestID}/investment [post]
func (lc *LoanController) CreateLoanInvestment(c *gin.Context) {

	// define variable to store struct
	var loanInvestment models.LoanInvestment

	// bind parameter from input
	loanRequestID := c.Param("loanRequestID")

	// bind value from user input
	if err := c.ShouldBindJSON(&loanInvestment); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}

	// call CreateLoanInvestment service
	err := lc.service.CreateLoanInvestment(loanRequestID, &loanInvestment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error",
			"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Loan investment success",
	})
}

// LoanDisbursement godoc
// @Summary Create loan disbursement
// @Description Create loan disbursement
// @Tags Loan
// @Accept json
// @Produce json
// @Param loanRequestID path string true "Loan Request ID"
// @Param models.LoanDisbursementSwagger body models.LoanDisbursementSwagger true "Loan Disbursement parameters"
// @Success 200
// @Router /api/v1/loan/{loanRequestID}/disbursement [post]
func (lc *LoanController) CreateLoanDisbursement(c *gin.Context) {

	// define variable to store struct
	var loanDisbursement models.LoanDisbursement

	// bind parameter from input
	loanRequestID := c.Param("loanRequestID")

	// bind value from user input
	if err := c.ShouldBindJSON(&loanDisbursement); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}

	// call CreateLoanDisbursement service
	err := lc.service.CreateLoanDisbursement(loanRequestID, &loanDisbursement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error",
			"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Loan disbursement success",
	})
}
