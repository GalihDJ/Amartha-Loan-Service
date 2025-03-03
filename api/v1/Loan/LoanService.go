package loan

import "fmt"

type ILoanService interface {
	TestAPIPost(testingInput string) (string, error)
	TestAPIGet() (string, error)
}

type LoanService struct {
}

func NewLoanService() ILoanService {
	return &LoanService{}
}

// TestAPIGet implements ILoanService.
func (ls *LoanService) TestAPIGet() (string, error) {

	fmt.Println("In TestAPIGet LoanService...")

	return "Test API Get Successful", nil
}

// TestAPIPost implements ILoanService.
func (ls *LoanService) TestAPIPost(testingInput string) (string, error) {
	fmt.Println("In TestAPIGet LoanService...")

	output := "User input: " + testingInput

	return output, nil
}
