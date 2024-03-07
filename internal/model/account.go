package model

type Account struct {
	ID             uint32 `json:"id"`
	DocumentNumber string `json:"document_number"`
}

func NewAccount(documentNumner string) *Account {
	return &Account{
		DocumentNumber: documentNumner,
	}
}
