package handlers

type V2Input struct {
	Username  string `json:"username"`
	Upassword string `json:"password"`
	U         User   `json:"load"`
}
