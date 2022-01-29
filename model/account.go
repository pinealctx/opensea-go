package model

type User struct {
	Username string `opensea:"username" json:"username"`
}

type Account struct {
	User          *User  `opensea:"user" json:"user"`
	ProfileImgURL string `opensea:"profile_img_url" json:"profileImgURL"`
	Address       string `opensea:"address" json:"address"`
	Config        string `opensea:"config" json:"config"`
}

func (a *Account) AccountAddress() string {
	if a == nil {
		return ""
	}

	return a.Address
}
