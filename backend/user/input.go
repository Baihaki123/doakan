package user

type RegisterUserInput struct {
	Nama      string `json:"nama" binding:"required"`
	Pekerjaan string `json:"pekerjaan" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

// web CMS
type FormCreateUserInput struct {
	Nama      string `form:"nama" binding:"required"`
	Email     string `form:"email" binding:"required"`
	Pekerjaan string `form:"pekerjaan" binding:"required"`
	Password  string `form:"password" binding:"required"`
	Error     error
}

type FormUpdateUserInput struct {
	ID        int
	Nama      string `form:"nama" binding:"required"`
	Email     string `form:"email" binding:"required"`
	Pekerjaan string `form:"pekerjaan" binding:"required"`
	Error     error
}
