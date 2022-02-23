package user

import (
	"be/entities"
)

//Request
type UserCreateRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"email"`
	Password string `json:"password" form:"password" validate:"required"`
}
type UserUpdateRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

//=======================================
//Response

type GetUserResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.User `json:"data"`
}
