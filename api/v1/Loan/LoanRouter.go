package loan

import (
	"github.com/gin-gonic/gin"

	utils "amartha-loan-service/utils/Connections"
)

func InitializeLoan(router *gin.Engine, connPSQL *utils.ConnectionPSQL) {

	loanRepository := NewLoanRepository(connPSQL)
	loanService := NewLoanService(loanRepository)
	loanController := NewLoanController(loanService)

	InitializeLoanRouters(router, loanController)
}

func InitializeLoanRouters(router *gin.Engine, controller *LoanController) {
	loanGroup := router.Group("/api/v1/")
	{
		loanGroup.POST("loan", controller.CreateLoanRequest)
		loanGroup.GET("loan/:loanRequestID", controller.GetLoanRequestById)
		loanGroup.PUT("loan/:loanRequestID", controller.ApproveLoanRequest)
		loanGroup.POST("loan/:loanRequestID/investment", controller.CreateLoanInvestment)
		loanGroup.POST("loan/:loanRequestID/disbursement", controller.CreateLoanDisbursement)
	}
}
