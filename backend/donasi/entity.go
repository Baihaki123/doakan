package donasi

import (
	"doakan/user"
	"time"

	"github.com/leekchan/accounting"
)

type Donation struct {
	ID               int
	UserID           int
	Nama             string
	DeskripsiSingkat string
	Deskripsi        string
	Perks            string
	DonaturCount     int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DonationImages   []DonationImage
	User             user.User
}

func (c Donation) GoalAmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp ", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(c.GoalAmount)
}

func (c Donation) CurrentAmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp ", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(c.CurrentAmount)
}

type DonationImage struct {
	ID         int
	DonationID int
	Filename   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
