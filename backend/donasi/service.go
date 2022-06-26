package donasi

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetDonasi(userID int) ([]Donation, error)
	GetDonasiByID(input GetDonationDetailInput) (Donation, error)
	CreateDonation(input CreateDonationInput) (Donation, error)
	UpdateDonation(inputID GetDonationDetailInput, inputData CreateDonationInput) (Donation, error)
	SaveDonationImage(input CreateDonationImageInput, fileLocation string) (DonationImage, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetDonasi(userID int) ([]Donation, error) {
	if userID != 0 {
		donations, err := s.repository.GetByUserID(userID)
		if err != nil {
			return donations, err
		}

		return donations, nil
	}

	donations, err := s.repository.GetAll()
	if err != nil {
		return donations, err
	}

	return donations, nil
}

func (s *service) GetDonasiByID(input GetDonationDetailInput) (Donation, error) {
	donation, err := s.repository.GetByID(input.ID)
	if err != nil {
		return donation, err
	}

	return donation, nil
}

func (s *service) CreateDonation(input CreateDonationInput) (Donation, error) {
	donation := Donation{}
	donation.Nama = input.Nama
	donation.DeskripsiSingkat = input.DeskripsiSingkat
	donation.Deskripsi = input.Deskripsi
	donation.GoalAmount = input.GoalAmount
	donation.Perks = input.Perks
	donation.UserID = input.User.ID

	// pembuatan slug
	slugCandidate := fmt.Sprintf("%s %d", input.Nama, input.User.ID)
	donation.Slug = slug.Make(slugCandidate)

	newDonation, err := s.repository.Save(donation)
	if err != nil {
		return newDonation, err
	}

	return newDonation, nil
}

func (s *service) UpdateDonation(inputID GetDonationDetailInput, inputData CreateDonationInput) (Donation, error) {
	donation, err := s.repository.GetByID(inputID.ID)
	if err != nil {
		return donation, err
	}

	if donation.UserID != inputData.User.ID {
		return donation, errors.New("ini bukan donasi milik anda")
	}

	donation.Nama = inputData.Nama
	donation.DeskripsiSingkat = inputData.DeskripsiSingkat
	donation.Deskripsi = inputData.Deskripsi
	donation.GoalAmount = inputData.GoalAmount
	donation.Perks = inputData.Perks

	updatedDonation, err := s.repository.Update(donation)
	if err != nil {
		return updatedDonation, err
	}

	return updatedDonation, nil
}

func (s *service) SaveDonationImage(input CreateDonationImageInput, fileLocation string) (DonationImage, error) {
	donation, err := s.repository.GetByID(input.DonationID)
	if err != nil {
		return DonationImage{}, nil
	}

	if donation.UserID != input.User.ID {
		return DonationImage{}, errors.New("ini bukan donasi milik anda")
	}

	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1

		_, err := s.repository.MarkAllImagesAsNonPrimary(input.DonationID)
		if err != nil {
			return DonationImage{}, err
		}
	}

	donationImage := DonationImage{}
	donationImage.DonationID = input.DonationID
	donationImage.IsPrimary = isPrimary
	donationImage.Filename = fileLocation

	newDonationImage, err := s.repository.CreateImage(donationImage)
	if err != nil {
		return newDonationImage, err
	}

	return newDonationImage, nil
}
