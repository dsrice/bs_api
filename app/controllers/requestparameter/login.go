package requestparameter

type Login struct {
	LoginID  string `json:"login_id"`
	Password string `json:"password"`
}