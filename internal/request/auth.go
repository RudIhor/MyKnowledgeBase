package request

type RegisterUserRequest struct {
	FirstName   string `gorm:"default:null" json:"first_name,omitempty" validate:"omitempty,min=3,max=100"`
	LastName    string `gorm:"default:null" json:"last_name,omitempty" validate:"omitempty,min=3,max=100"`
	Email       string `gorm:"unique" json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8"`
	AccessToken string `gorm:"-" json:"access_token,omitempty"`
}

type LoginUserRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
