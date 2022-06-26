package handler

import (
	"doakan/donasi"
	"doakan/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type donationHandler struct {
	donationService donasi.Service
	userService     user.Service
}

func NewDonationHandler(donationService donasi.Service, userService user.Service) *donationHandler {
	return &donationHandler{donationService, userService}
}

func (h *donationHandler) Index(c *gin.Context) {
	donations, err := h.donationService.GetDonasi(0)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "donation_index.html", gin.H{"donations": donations})
}

func (h *donationHandler) New(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := donasi.FormCreateDonationInput{}
	input.Users = users

	c.HTML(http.StatusOK, "donation_new.html", input)
}

func (h *donationHandler) Create(c *gin.Context) {
	var input donasi.FormCreateDonationInput

	err := c.ShouldBind(&input)
	if err != nil {
		users, e := h.userService.GetAllUsers()
		if e != nil {
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}

		input.Users = users
		input.Error = err

		c.HTML(http.StatusOK, "donation_new.html", input)
		return
	}

	user, err := h.userService.GetUserByID(input.UserID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	createDonationInput := donasi.CreateDonationInput{}
	createDonationInput.Nama = input.Nama
	createDonationInput.DeskripsiSingkat = input.DeskripsiSingkat
	createDonationInput.Deskripsi = input.Deskripsi
	createDonationInput.GoalAmount = input.GoalAmount
	createDonationInput.Perks = input.Perks
	createDonationInput.User = user

	_, err = h.donationService.CreateDonation(createDonationInput)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/donations")
}

func (h *donationHandler) NewImage(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	c.HTML(http.StatusOK, "donation_image.html", gin.H{"ID": id})
}

func (h *donationHandler) CreateImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	existingDonation, err := h.donationService.GetDonasiByID(donasi.GetDonationDetailInput{ID: id})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	userID := existingDonation.UserID

	path := fmt.Sprintf("donation-images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	createDonationImageInput := donasi.CreateDonationImageInput{}
	createDonationImageInput.DonationID = id
	createDonationImageInput.IsPrimary = true

	userDonation, err := h.userService.GetUserByID(userID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	createDonationImageInput.User = userDonation

	_, err = h.donationService.SaveDonationImage(createDonationImageInput, path)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/donations")
}

func (h *donationHandler) Edit(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	existingDonation, err := h.donationService.GetDonasiByID(donasi.GetDonationDetailInput{ID: id})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	input := donasi.FormUpdateDonationInput{}
	input.ID = existingDonation.ID
	input.Nama = existingDonation.Nama
	input.DeskripsiSingkat = existingDonation.DeskripsiSingkat
	input.Deskripsi = existingDonation.Deskripsi
	input.GoalAmount = existingDonation.GoalAmount
	input.Perks = existingDonation.Perks

	c.HTML(http.StatusOK, "donation_edit.html", input)
}

func (h *donationHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var input donasi.FormUpdateDonationInput

	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = err
		input.ID = id
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	existingDonation, err := h.donationService.GetDonasiByID(donasi.GetDonationDetailInput{ID: id})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	userID := existingDonation.UserID

	userDonation, err := h.userService.GetUserByID(userID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	updateInput := donasi.CreateDonationInput{}
	updateInput.Nama = input.Nama
	updateInput.DeskripsiSingkat = input.DeskripsiSingkat
	updateInput.Deskripsi = input.Deskripsi
	updateInput.GoalAmount = input.GoalAmount
	updateInput.Perks = input.Perks
	updateInput.User = userDonation

	_, err = h.donationService.UpdateDonation(donasi.GetDonationDetailInput{ID: id}, updateInput)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/donations")
}

func (h *donationHandler) Show(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	existingDonation, err := h.donationService.GetDonasiByID(donasi.GetDonationDetailInput{ID: id})
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "donation_show.html", existingDonation)
}
