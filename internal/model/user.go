package model

import (
	requests "github.com/RivGames/my-knowledge-base/internal/requests/auth"
)

type User struct {
	Model
	requests.RegisterUserRequest
}
