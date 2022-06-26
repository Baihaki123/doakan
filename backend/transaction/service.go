package transaction

import (
	"doakan/donasi"
	"doakan/payment"
	"errors"
	"strconv"
)

type Service interface {
	GetTransactionsByDonationID(input GetDonationTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
	GetAllTransactions() ([]Transaction, error)
}

type service struct {
	repository         Repository
	donationRepository donasi.Repository
	paymentService     payment.Service
}

func NewService(repository Repository, donationRepository donasi.Repository, paymentService payment.Service) *service {
	return &service{repository, donationRepository, paymentService}
}

func (s *service) GetTransactionsByDonationID(input GetDonationTransactionsInput) ([]Transaction, error) {
	donation, err := s.donationRepository.GetByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if donation.UserID != input.User.ID {
		return []Transaction{}, errors.New("ini bukan donasi milik anda")
	}

	transactions, err := s.repository.GetByDonationID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionsByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.DonationID = input.DonationID
	transaction.Amount = input.Amount
	transaction.UserID = input.User.ID
	transaction.Status = "pending"
	transaction.VANumber = "0"
	transaction.Bank = "default"

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.Transaction{
		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL

	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (s *service) ProcessPayment(input TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.repository.GetByID(transaction_id)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	if input.VANumbers != nil {
		transaction.VANumber = input.VANumbers[0].VANumber
		transaction.Bank = input.VANumbers[0].Bank
	}

	if input.MaskedCard != "" {
		transaction.VANumber = input.MaskedCard
		transaction.Bank = input.Bank
	}

	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.donationRepository.GetByID(updatedTransaction.DonationID)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		campaign.DonaturCount = campaign.DonaturCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount

		_, err := s.donationRepository.Update(campaign)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) GetAllTransactions() ([]Transaction, error) {
	transactions, err := s.repository.FindAll()
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
