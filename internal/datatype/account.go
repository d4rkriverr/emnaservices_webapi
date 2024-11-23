package datatype

type Account struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	AccessToken string `json:"access_token"`
}
