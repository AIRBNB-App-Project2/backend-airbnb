package user

import "be/entities"

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type GetUserResponseFormat struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    entities.User `json:"data"`
}
