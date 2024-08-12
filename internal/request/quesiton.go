package request

type CreateQuesitonRequest struct {
	Title      string `json:"title" validate:"required,gt=5,lte=100"`
	UserId     uint   `json:"user_id"`
	IsAnswered *bool  `json:"is_answered" validate:"boolean" gorm:"default:false"`
}

type UpdateQuestionRequest struct {
	Title string `json:"title" validate:"gt=5,lte=100"`
}
