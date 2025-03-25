package domain

import "github.com/google/uuid"

type Transaction struct {
	Id     string  `json:"tx_id"`
	From   Account `json:"from"`
	To     Account `json:"to"`
	Amount int     `json:"amount"`
}

func NewTransaction(uuid uuid.UUID, from Account, to Account, amount int) *Transaction {
	return &Transaction{
		Id:     uuid.String(),
		From:   from,
		To:     to,
		Amount: amount,
	}
}
