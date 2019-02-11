package structs

type Token struct {
	UserID      uint   `json:"userID"`
	UserName    string `json:"userName"`
	AccessToken string `json:"accessToken"`
}
