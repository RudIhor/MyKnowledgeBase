package model

import "github.com/RivGames/my-knowledge-base/internal/request"

type Answer struct {
	Model
	request.CreateAnswerRequest
}
