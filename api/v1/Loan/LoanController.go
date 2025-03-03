package loan

import (
	"net/http"

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

// TestAPIPost godoc
// @Summary Test POST API
// @Description Test POST API
// @Tags Loan
// @Accept json
// @Produce json
// @Param testingInput formData string true "Testing Input"
// @Success 200
// @Router /api/v1/testing [post]
func (lc *LoanController) TestAPIPost(c *gin.Context) {

	testingInput := c.PostForm("testingInput")
	if testingInput == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input is required"})
		return
	}

	test, err := lc.service.TestAPIPost(testingInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, test)
}

// TestAPIGet godoc
// @Security BearerAuth
// @Summary Test GET API
// @Description Test GET API
// @Tags Loan
// @Produce json
// @Success 200
// @Router /api/v1/testing [get]
func (lc *LoanController) TestAPIGet(c *gin.Context) {

	test, err := lc.service.TestAPIGet()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, test)
}
