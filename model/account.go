package model

type User struct {
	Username string `json:"username" camel:"username"`
}

type Account struct {
	User          *User  `json:"user" camel:"user"`
	ProfileImgURL string `json:"profile_img_url" camel:"profileImgURL"`
	Address       string `json:"address" camel:"address"`
	Config        string `json:"config" camel:"config"`
}
