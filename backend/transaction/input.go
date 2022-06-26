package transaction

import "doakan/user"

type GetDonationTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	Amount     int `json:"amount" binding:"required"`
	DonationID int `json:"donation_id" binding:"required"`
	User       user.User
}

type TransactionNotificationInput struct {
	TransactionStatus string             `json:"transaction_status"`
	OrderID           string             `json:"order_id"`
	PaymentType       string             `json:"payment_type"`
	FraudStatus       string             `json:"fraud_status"`
	VANumbers         []TransactionNotip `json:"va_numbers"`
	MaskedCard        string             `json:"masked_card"`
	Bank              string             `json:"bank"`
}

type TransactionNotip struct {
	VANumber string `json:"va_number"`
	Bank     string `json:"bank"`
}
