package user

type UserFormatter struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Pekerjaan string `json:"pekerjaan"`
	Email     string `json:"email"`
	Token     string `json:"token"`
	ImageURL  string `json:"image_url"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:        user.ID,
		Nama:      user.Nama,
		Pekerjaan: user.Pekerjaan,
		Email:     user.Email,
		Token:     token,
		ImageURL:  user.ImageProfile,
	}

	return formatter
}
