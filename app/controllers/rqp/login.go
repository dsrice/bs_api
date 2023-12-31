package rqp

type Login struct {
	LoginID  string `json:"login_id" validate:"required"`
	Password string `json:"password"  validate:"required"`
}