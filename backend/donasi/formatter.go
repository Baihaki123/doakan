package donasi

import "strings"

type DonationFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Nama             string `json:"nama"`
	DeskripsiSingkat string `json:"deskripsi_singkat"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatDonation(donation Donation) DonationFormatter {
	donationFormatter := DonationFormatter{}
	donationFormatter.ID = donation.ID
	donationFormatter.UserID = donation.UserID
	donationFormatter.Nama = donation.Nama
	donationFormatter.DeskripsiSingkat = donation.DeskripsiSingkat
	donationFormatter.GoalAmount = donation.GoalAmount
	donationFormatter.CurrentAmount = donation.CurrentAmount
	donationFormatter.Slug = donation.Slug
	donationFormatter.ImageURL = ""

	if len(donation.DonationImages) > 0 {
		donationFormatter.ImageURL = donation.DonationImages[0].Filename
	}

	return donationFormatter
}

func FormatDonations(donations []Donation) []DonationFormatter {
	donationsFormatter := []DonationFormatter{}

	for _, donation := range donations {
		donationFormatter := FormatDonation(donation)
		donationsFormatter = append(donationsFormatter, donationFormatter)
	}

	return donationsFormatter
}

type DonationDetailFormatter struct {
	ID               int                      `json:"id"`
	Nama             string                   `json:"nama"`
	DeskripsiSingkat string                   `json:"deskripsi_singkat"`
	Deskripsi        string                   `json:"deskripsi"`
	ImageURL         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	DonaturCount     int                      `json:"donatur_count"`
	UserID           int                      `json:"user_id"`
	Slug             string                   `json:"slug"`
	Perks            []string                 `json:"perks"`
	User             DonationUserFormatter    `json:"user"`
	Images           []DonationImageFormatter `json:"images"`
}

type DonationUserFormatter struct {
	Nama     string `json:"nama"`
	ImageURL string `json:"image_url"`
}

type DonationImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatDonationDetail(donation Donation) DonationDetailFormatter {
	donationDetailFormatter := DonationDetailFormatter{}
	donationDetailFormatter.ID = donation.ID
	donationDetailFormatter.Nama = donation.Nama
	donationDetailFormatter.DeskripsiSingkat = donation.DeskripsiSingkat
	donationDetailFormatter.Deskripsi = donation.Deskripsi
	donationDetailFormatter.GoalAmount = donation.GoalAmount
	donationDetailFormatter.CurrentAmount = donation.CurrentAmount
	donationDetailFormatter.DonaturCount = donation.DonaturCount
	donationDetailFormatter.UserID = donation.UserID
	donationDetailFormatter.Slug = donation.Slug
	donationDetailFormatter.ImageURL = ""

	if len(donation.DonationImages) > 0 {
		donationDetailFormatter.ImageURL = donation.DonationImages[0].Filename
	}

	var perks []string
	for _, perk := range strings.Split(donation.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	donationDetailFormatter.Perks = perks

	user := donation.User

	donationUserFormatter := DonationUserFormatter{}
	donationUserFormatter.Nama = user.Nama
	donationUserFormatter.ImageURL = user.ImageProfile

	donationDetailFormatter.User = donationUserFormatter

	images := []DonationImageFormatter{}

	for _, image := range donation.DonationImages {
		donationImageFormatter := DonationImageFormatter{}
		donationImageFormatter.ImageURL = image.Filename

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		donationImageFormatter.IsPrimary = isPrimary

		images = append(images, donationImageFormatter)
	}

	donationDetailFormatter.Images = images

	return donationDetailFormatter
}
