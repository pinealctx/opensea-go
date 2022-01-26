package model

type User struct {
	Username string `json:"username"`
}

type Account struct {
	User          *User  `json:"user"`
	ProfileImgURL string `json:"profile_img_url"`
	Address       string `json:"address"`
	Config        string `json:"config"`
}
