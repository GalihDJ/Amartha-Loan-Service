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

	var loanRequest models.LoanRequest

	if err := c.ShouldBindJSON(&loanRequest); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := lc.service.CreateLoanRequest(&loanRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"rows_affected": rowsAffected})
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

	loanRequestID := c.Param("loanRequestID")

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

	var approveLoanRequest models.LoanApproval

	loanRequestID := c.Param("loanRequestID")

	if err := c.ShouldBindJSON(&approveLoanRequest); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}

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

	var loanInvestment models.LoanInvestment

	loanRequestID := c.Param("loanRequestID")

	if err := c.ShouldBindJSON(&loanInvestment); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}

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
// @Param models.LoanRequestSwagger body models.LoanRequestSwagger true "Loan Request parameters"
// @Success 200
// @Router /api/v1/loan/disbursement [post]
func (lc *LoanController) LoanDisbursement(c *gin.Context) {
	var loanRequest models.LoanRequest

	// userCred := sec.GetCurrentUser(c)
	// fmt.Println("userCred: ", userCred.ID)

	if err := c.ShouldBindJSON(&loanRequest); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}

	// documentLibrary.UserID = userCred.ID

	// loanRequestID, err := lc.service.CreateLoanRequest(&loanRequest)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"library id": loanRequestID})
}

// =============================================================================== TEST API ===============================================================================

// CreateInvestor godoc
// @Summary Create new investor
// @Description Create a new investor
// @Tags Loan
// @Accept json
// @Produce json
// @Param models.InvestorSwagger body models.InvestorSwagger true "Investor parameters"
// @Success 200
// @Router /api/v1/loan/investor [post]
func (lc *LoanController) CreateInvestor(c *gin.Context) {
	var investor models.Investor

	// userCred := sec.GetCurrentUser(c)
	// fmt.Println("userCred: ", userCred.ID)

	if err := c.ShouldBindJSON(&investor); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}

	createdInvestorID, err := lc.service.CreateInvestor(&investor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"library id": createdInvestorID})
}
