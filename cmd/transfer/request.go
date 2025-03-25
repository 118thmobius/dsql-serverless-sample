package main

type TransferRequest struct {
	FromId      string `json:"from_account_id"`
	ToId        string `json:"to_account_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

func NewTransferRequest(fromId, toId string, amount int, description string) *TransferRequest {
	return &TransferRequest{
		FromId:      fromId,
		ToId:        toId,
		Amount:      amount,
		Description: description,
	}
}
