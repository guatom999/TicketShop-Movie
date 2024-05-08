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
		*CustomerProfile
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
