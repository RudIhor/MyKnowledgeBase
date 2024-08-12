package request

type CreateQuesitonRequest struct {
	Title      string `validate:"required,gt=5,lte=100" json:"title"`
	UserId     uint   `json:"user_id"`
	IsAnswered *bool  `validate:"boolean" gorm:"default:false" json:"is_answered"`
}

type UpdateQuestionRequest struct {
	Title string `validate:"gt=5,lte=100" json:"title"`
}
