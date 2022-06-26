package donasi

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Donation, error)
	GetByUserID(userID int) ([]Donation, error)
	GetByID(ID int) (Donation, error)
	Save(donation Donation) (Donation, error)
	Update(donation Donation) (Donation, error)
	CreateImage(donationImage DonationImage) (DonationImage, error)
	MarkAllImagesAsNonPrimary(donationID int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Donation, error) {
	var donations []Donation
	err := r.db.Preload("DonationImages", "donation_images.is_primary = 1").Find(&donations).Error
	if err != nil {
		return donations, err
	}

	return donations, nil
}

func (r *repository) GetByUserID(userID int) ([]Donation, error) {
	var donations []Donation

	err := r.db.Where("user_id = ?", userID).Preload("DonationImages", "donation_images.is_primary = 1").Find(&donations).Error
	if err != nil {
		return donations, err
	}

	return donations, nil
}

func (r *repository) GetByID(ID int) (Donation, error) {
	var donasi Donation

	err := r.db.Preload("User").Preload("DonationImages").Where("id = ?", ID).Find(&donasi).Error
	if err != nil {
		return donasi, err
	}

	return donasi, nil
}

func (r *repository) Save(donation Donation) (Donation, error) {
	err := r.db.Create(&donation).Error
	if err != nil {
		return donation, err
	}

	return donation, nil
}

func (r *repository) Update(donation Donation) (Donation, error) {
	err := r.db.Save(&donation).Error
	if err != nil {
		return donation, err
	}

	return donation, nil
}

func (r *repository) CreateImage(donationImage DonationImage) (DonationImage, error) {
	err := r.db.Create(&donationImage).Error
	if err != nil {
		return donationImage, err
	}

	return donationImage, nil
}

func (r *repository) MarkAllImagesAsNonPrimary(donationID int) (bool, error) {
	err := r.db.Model(&DonationImage{}).Where("donation_id = ?", donationID).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
