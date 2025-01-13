package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

// GetAccountDetailsRequest represents the request structure for the GetAccountDetails API.
type GetAccountDetailsRequest struct {
	UserID        string `validate:"required,alphanum,min=10,max=32"`
	DeviceID      string `validate:"required,alphanum,min=10,max=32"`
	BankCode      string `validate:"required,len=3,alpha"`
	MerchantID    string `validate:"required,alphanum,min=5,max=10"`
	AccountNumber int    `validate:"required,accountNumber"`
	AccountType   string `validate:"required,accountType"`
}

// NewGetAccountDetailsRequestFromQuery creates a new GetAccountDetailsRequest from the URL query parameters.
func NewGetAccountDetailsRequestFromQuery(r *http.Request) (*GetAccountDetailsRequest, error) {
	req := &GetAccountDetailsRequest{}
	if err := decodeQuery(r.URL.Query(), req); err != nil {
		return nil, err
	}
	return req, validate.Struct(req)
}

func decodeQuery(query map[string][]string, req *GetAccountDetailsRequest) error {
	for key, values := range query {
		value := strings.Join(values, ",")
		switch key {
		case "userID":
			req.UserID = value
		case "deviceID":
			req.DeviceID = value
		case "bankCode":
			req.BankCode = value
		case "merchantId":
			req.MerchantID = value
		case "accountNumber":
			req.AccountNumber, _ = strconv.Atoi(value)
		case "accountType":
			req.AccountType = value
		}
	}
	return nil
}

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("accountNumber", validateAccountNumber)
	validate.RegisterValidation("accountType", validateAccountType)
}

func validateAccountNumber(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().Int()

	if isStartsWith100(fieldValue) && fieldValue < 10e9 {
		return true
	}

	return false
}

func validateAccountType(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()
	switch strings.ToLower(fieldValue) {
	case "saving", "current":
		return true
	default:
		return false
	}
}

func isStartsWith100(number int64) bool {
	digits := []int64{}
	for number > 9 {
		r := number % 10
		digits = append(digits, r)
		number /= 10
	}
	digits = append(digits, number)

	lastThreeDigits := digits[len(digits)-3:]

	first := lastThreeDigits[2]
	second := lastThreeDigits[1]
	third := lastThreeDigits[0]

	if first != 1 || second != 0 || third != 0 {
		return false
	}

	return true
}

func main() {
	r, _ := http.NewRequest("GET", "/?userID=abcd1234567890&deviceID=abcdefghij123456&bankCode=XYZ&merchantId=pqrst&accountNumber=100332211&accountType=Saving", nil)
	req, err := NewGetAccountDetailsRequestFromQuery(r)
	if err != nil {
		fmt.Println("Validation error:", err)
		return
	}
	fmt.Println("Valid request:", req)
}
