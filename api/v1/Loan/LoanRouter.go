package loan

import (
	"github.com/gin-gonic/gin"
)

func InitializeLoan(router *gin.Engine) {

	loanService := NewLoanService()
	loanController := NewLoanController(loanService)

	InitializeLoanRouters(router, loanController)
}

func InitializeLoanRouters(router *gin.Engine, controller *LoanController) {
	loanGroup := router.Group("/api/v1/")
	{
		loanGroup.GET("loan", controller.TestAPIGet)
		loanGroup.POST("loan", controller.TestAPIPost)
	}
}
