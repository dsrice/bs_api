package ci

import "go.uber.org/dig"

type InController struct {
	dig.In
	Login LoginController
	UserController UserController
}
