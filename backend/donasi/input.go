package donasi

import "doakan/user"

type GetDonationDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateDonationInput struct {
	Nama             string `json:"nama" binding:"required"`
	DeskripsiSingkat string `json:"deskripsi_singkat" binding:"required"`
	Deskripsi        string `json:"deskripsi" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             user.User
}

type CreateDonationImageInput struct {
	DonationID int  `form:"donation_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"`
	User       user.User
}

type FormCreateDonationInput struct {
	Nama             string `form:"nama" binding:"required"`
	DeskripsiSingkat string `form:"deskripsi_singkat" binding:"required"`
	Deskripsi        string `form:"deskripsi" binding:"required"`
	GoalAmount       int    `form:"goal_amount" binding:"required"`
	Perks            string `form:"perks" binding:"required"`
	UserID           int    `form:"user_id" binding:"required"`
	Users            []user.User
	Error            error
}

type FormUpdateDonationInput struct {
	ID               int
	Nama             string `form:"nama" binding:"required"`
	DeskripsiSingkat string `form:"deskripsi_singkat" binding:"required"`
	Deskripsi        string `form:"deskripsi" binding:"required"`
	GoalAmount       int    `form:"goal_amount" binding:"required"`
	Perks            string `form:"perks" binding:"required"`
	Error            error
	User             user.User
}
