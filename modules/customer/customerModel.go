package customer

type (
	RegisterReq struct {
		Email    string `json:"email"`
		UserName string `json:"username"`
		Password string `json:"password" `
	}

	LoginReq struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	LogoutReq struct {
		CredentialId string `json:"credential_id" form:"credential_id"`
	}
	CustomerProfileRes struct {
		Status           string `json:"status"`
		*CustomerProfile `json:"user"`
	}

	CustomerProfile struct {
		Id         string         `json:"_id"`
		CustomerId string         `json:"customer_id"`
		Email      string         `json:"email"`
		UserName   string         `json:"user_name"`
		ImageUrl   string         `json:"image_url"`
		Created_At string         `json:"created_at"`
		Updated_At string         `json:"updated_at"`
		Credential *CredentailRes `json:"credential"`
	}

	FindAccessTokenReq struct {
		AccessToken string `json:"access_token" form:"access_token"`
	}

	CredentailRes struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	CustomerRefreshTokenReq struct {
		CredentialId string `json:"credential_id" form:"credential_id"`
		CustomerId   string `json:"customer_id" validate:"required,max=128"`
		RefreshToken string `json:"refresh_token" form:"refresh_token" `
	}

	// TokenRes struct {
	// 	AccessToken  string `json:"access_token"`
	// 	RefreshToken string `json:"refresh_token"`
	// }
)
