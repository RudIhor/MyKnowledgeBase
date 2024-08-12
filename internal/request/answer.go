package request

type CreateAnswerRequest struct {
	Text       string `json:"text" validate:"required,gt=30,lt=3000"`
	QuestionId uint   `json:"question_id" valiate:"required"`
	UserId     uint   `json:"user_id"`
}

type UpdateAnswerRequest struct {
	Text string `json:"text" validate:"gt=30,lt=3000"`
}
