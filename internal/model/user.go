package model

import "github.com/RivGames/my-knowledge-base/internal/request"

type User struct {
	Model
	request.RegisterUserRequest
}
