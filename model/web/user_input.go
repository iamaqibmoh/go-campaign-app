package web

//input for API

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

//input for cms admin

type FormCreateUserInput struct {
	Name       string `form:"name" binding:"required"`
	Occupation string `form:"occupation" binding:"required"`
	Email      string `form:"email" binding:"required"`
	Password   string `form:"password" binding:"required"`
	Error      error
}
