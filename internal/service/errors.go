package service

import "errors"

var (
	ErrDocumentMustHave11Digits = errors.New("document number must have 11 digits")

	ErrAccountAlreadyExists = errors.New("account already exists")
	ErrAccountNotFound      = errors.New("account not found")

	ErrAccountBalanceNotFound  = errors.New("account balance not found")
	ErrAccountBalanceNotEnough = errors.New("not enough balance to account")

	ErrOperationTypeNotFound = errors.New("operation type not found")

	ErrTransactionAmountNotAllowed = errors.New("invalid amount for this operation type")
)
