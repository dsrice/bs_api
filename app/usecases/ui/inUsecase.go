package ui

import "go.uber.org/dig"

type InUsecase struct {
	dig.In
	Login LoginUsecase
}