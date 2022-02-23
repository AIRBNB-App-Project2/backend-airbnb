package user

type UserCreateResponse struct {
	User_uid string `json:"user_uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
