package transaction

import (
	"strings"
)

type RawTransaction struct {
	CreatedAt         string `csv:"Data transakcji"`
	BookedAt          string `csv:"Data ksi�gowania"`
	Contractor        string `csv:"Dane kontrahenta"`
	Title             string `csv:"Tytu�"`
	AccountNumber     string `csv:"Nr rachunku"`
	BankName          string `csv:"Nazwa banku"`
	Details           string `csv:"Szczeg�y"`
	TransactionNumber string `csv:"Nr transakcji"`
	Amount            string `csv:"Kwota transakcji (waluta rachunku)"`
	Currency          string `csv:"Waluta"`
	BlockedAmount     string `csv:"Kwota blokady/zwolnienie blokady"`
	Account           string `csv:"Konto"`
	Balance           string `csv:"Saldo po transakcji"`
}

type Transaction struct {
	CreatedAt         string `json:"cratedAt"`
	BookedAt          string `json:"bookedAt"`
	Contractor        string `json:"contractor"`
	Title             string `json:"title"`
	AccountNumber     string `json:"accountNumber"`
	BankName          string `json:"bankName"`
	Details           string `json:"details"`
	TransactionNumber string `json:"transactionNumber"`
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	BlockedAmount     string `json:"blockedAmount"`
	Account           string `json:"account"`
	Balance           string `json:"balance"`
}

func (t RawTransaction) ProcessTransaction() Transaction {
	transaction := Transaction{
		CreatedAt:         strings.TrimSpace(t.CreatedAt),
		BookedAt:          strings.TrimSpace(t.BookedAt),
		Contractor:        strings.TrimSpace(t.Contractor),
		Title:             strings.TrimSpace(t.Title),
		AccountNumber:     strings.TrimSpace(t.AccountNumber),
		BankName:          strings.TrimSpace(t.BankName),
		Details:           strings.TrimSpace(t.Details),
		TransactionNumber: strings.TrimSpace(t.TransactionNumber),
		Amount:            strings.TrimSpace(t.Amount),
		Currency:          strings.TrimSpace(t.Currency),
		BlockedAmount:     strings.TrimSpace(t.BlockedAmount),
		Account:           strings.TrimSpace(t.Account),
		Balance:           strings.TrimSpace(t.Balance),
	}

	return transaction
}
