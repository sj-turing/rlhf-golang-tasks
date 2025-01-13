package main

import (
	"fmt"
	"net/http"
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
	validate.RegisterStructTagNameFunc(func(fld reflect.StructField) string {
		name := strings.Split(fld.Tag.Get("validate"), ",")[0]
		if name == "accountNumber" || name == "accountType" {
			return name
		}
		return ""
	})
	validate.RegisterValidator("accountNumber", new(AccountNumberValidator))
	validate.RegisterValidator("accountType", new(AccountTypeValidator))
}

// AccountNumberValidator implements the Validator interface for custom validation of AccountNumber.
type AccountNumberValidator struct{}

func (v *AccountNumberValidator) Validate(data interface{}) error {
	accountNumber, ok := data.(int)
	if !ok {
		return fmt.Errorf("account number must be an integer")
	}
	if accountNumber < 100 || accountNumber > 9999999999 {
		return fmt.Errorf("account number must start with 100 and have less than or equal to 10 digits")
	}
	return nil
}

// AccountTypeValidator implements the Validator interface for custom validation of AccountType.
type AccountTypeValidator struct{}

func (v *AccountTypeValidator) Validate(data interface{}) error {
	accountType, ok := data.(string)
	if !ok {
		return fmt.Errorf("account type must be a string")
	}
	switch accountType {
	case "Saving", "Current":
		return nil
	default:
		return fmt.Errorf("account type must be either 'Saving' or 'Current'")
	}
}

func main() {
	r, _ := http.NewRequest("GET", "/?userID=abcd1234567890&deviceID=abcdefghij123456&bankCode=XYZ&merchantId=pqrst&accountNumber=1000000000&accountType=Saving", nil)
	req, err := NewGetAccountDetailsRequestFromQuery(r)
	if err != nil {
		fmt.Println("Validation error:", err)
		return
	}
	fmt.Println("Valid request:", req)
}
