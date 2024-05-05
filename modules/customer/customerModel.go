package customer

type (
	RegisterReq struct {
		Email    string `json:"email"`
		UserName string `json:"username"`
		Password string `json:"password" `
	}
)
