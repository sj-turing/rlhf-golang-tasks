package banking

import "errors"

var (
	ErrInvalidOwner            = errors.New("Owner name cannot be empty")
	ErrDepositAmountInvalid    = errors.New("Deposit amount must be greater than 0")
	ErrWithdrawalAmountInvalid = errors.New("Withdrawal amount must be greater than 0")
	ErrAccountNotFound         = errors.New("Account not found")
	ErrInsufficientBalance     = errors.New("Insufficient balance")
)
