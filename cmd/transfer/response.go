package main

import "github.com/118thmobius/dsql-serverless-sample/internal/domain"

type TransferResponse struct {
	Transaction domain.Transaction `json:"transaction"`
	Message     string             `json:"message"`
}

func NewTransferResponse(transaction domain.Transaction, message string) *TransferResponse {
	return &TransferResponse{
		Transaction: transaction,
		Message:     message,
	}
}
