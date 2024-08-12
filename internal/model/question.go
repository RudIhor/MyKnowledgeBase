package model

import "github.com/RivGames/my-knowledge-base/internal/request"

type Question struct {
	Model
	request.CreateQuesitonRequest
}
