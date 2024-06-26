package customer

type (
	RegisterReq struct {
		Email    string `json:"email"`
		UserName string `json:"username"`
		Password string `json:"password" `
	}

	LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	CustomerProfileRes struct {
		Status           string `json:"status"`
		*CustomerProfile `json:"user"`
	}

	CustomerProfile struct {
		Id         string         `json:"_id"`
		Email      string         `json:"email"`
		UserName   string         `json:"user_name"`
		ImageUrl   string         `json:"image_url"`
		Created_At string         `json:"created_at"`
		Updated_At string         `json:"updated_at"`
		Credential *CredentailRes `json:"credential"`
	}

	CredentailRes struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	// TokenRes struct {
	// 	AccessToken  string `json:"access_token"`
	// 	RefreshToken string `json:"refresh_token"`
	// }
)
