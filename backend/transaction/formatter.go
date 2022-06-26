package transaction

import "time"

type DonationTransactionFormatter struct {
	ID        int       `json:"id"`
	Nama      string    `json:"nama"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatDonationTransaction(transaction Transaction) DonationTransactionFormatter {
	formatter := DonationTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Nama = transaction.User.Nama
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt
	return formatter
}

func FormatDonationTransactions(transactions []Transaction) []DonationTransactionFormatter {
	if len(transactions) == 0 {
		return []DonationTransactionFormatter{}
	}

	var transactionsFormatter []DonationTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatDonationTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}

type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	VANumber  string            `json:"va_number"`
	Bank      string            `json:"bank"`
	CreatedAt time.Time         `json:"created_at"`
	Donation  DonationFormatter `json:"donation"`
}

type DonationFormatter struct {
	Nama     string `json:"nama"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.VANumber = transaction.VANumber
	formatter.Bank = transaction.Bank
	formatter.CreatedAt = transaction.CreatedAt

	donationFormatter := DonationFormatter{}
	donationFormatter.Nama = transaction.Donation.Nama
	donationFormatter.ImageURL = ""

	if len(transaction.Donation.DonationImages) > 0 {
		donationFormatter.ImageURL = transaction.Donation.DonationImages[0].Filename
	}

	formatter.Donation = donationFormatter

	return formatter
}

func FormateUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var transactionsFormatter []UserTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}

type TransactionFormatter struct {
	ID         int    `json:"id"`
	DonationID int    `json:"campaign_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	VANumber   string `json:"va_number"`
	Bank       string `json:"bank"`
	PaymentURL string `json:"payment_url"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.DonationID = transaction.DonationID
	formatter.UserID = transaction.UserID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.VANumber = transaction.VANumber
	formatter.Bank = transaction.Bank
	formatter.PaymentURL = transaction.PaymentURL
	return formatter
}
