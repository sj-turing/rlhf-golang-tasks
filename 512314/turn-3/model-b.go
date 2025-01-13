package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

// custom validation function for AccountNumber
func startsWith100(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	return strings.HasPrefix(field, "100") && len(field) <= 10
}

// custom validation function for AccountType
func validAccountType(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	return field == "Saving" || field == "Current"
}

// GetAccountDetailsRequest represents the request structure for the GetAccountDetails API.
type GetAccountDetailsRequest struct {
	UserID        string `validate:"required,alphanum,min=10,max=32"`
	DeviceID      string `validate:"required,alphanum,min=10,max=32"`
	BankCode      string `validate:"required,len=3,alpha"`
	MerchantID    string `validate:"required,alphanum,min=5,max=10"`
	AccountNumber int    `validate:"required,custom=startsWith100"`
	AccountType   string `validate:"required,custom=validAccountType"`
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
			req.AccountNumber = atoi(value)
		case "accountType":
			req.AccountType = value
		}
	}
	return nil
}

// atoi converts a string to an int and returns an error if conversion fails
func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err) // Handle error appropriately in your application
	}
	return n
}

var validate *validator.Validate

func init() {
	validate = validator.New()
	// Register custom validation functions
	validate.RegisterValidation("startsWith100", startsWith100)
	validate.RegisterValidation("validAccountType", validAccountType)
}

func main() {
	r, _ := http.NewRequest("GET", "/?userID=123&deviceID=abc&bankCode=xyz&merchantId=pqr&accountNumber=1001234567&accountType=Saving", nil)
	req, err := NewGetAccountDetailsRequestFromQuery(r)
	if err != nil {
		fmt.Println("Validation error:", err)
		return
	}
	fmt.Println("Valid request:", req)
}
