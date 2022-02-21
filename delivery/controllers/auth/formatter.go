package auth

type Userlogin struct {
	Email    string `json:"email"  form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginRespFormat struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
