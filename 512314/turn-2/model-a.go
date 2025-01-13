package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

// GetAccountDetailsRequest represents the request structure for the GetAccountDetails API.
type GetAccountDetailsRequest struct {
	UserID     string `validate:"required,alphanum,min=10,max=32"`
	DeviceID   string `validate:"required,alphanum,min=10,max=32"`
	BankCode   string `validate:"required,alpha,uppercase,eq=3"`
	MerchantID string `validate:"required,alphanum,min=5,max=10"`
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
		}
	}
	return nil
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func main() {
	r, _ := http.NewRequest("GET", "/?userID=1234567890abcdefghijklmnopqrs&deviceID=xyz1234567890&bankCode=XYZ&merchantId=123456789", nil)
	req, err := NewGetAccountDetailsRequestFromQuery(r)
	if err != nil {
		fmt.Println("Validation error:", err)
		return
	}
	fmt.Println("Valid request:", req)
}
