package domain

type StatefulAccount struct {
	Account
	Balance int `json:"balance"`
}
