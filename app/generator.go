package main

import (
	"app/controllers/cg"
	"app/infra/genarator"
	"app/repositories/rg"
	"app/usecases/ug"
	"log"
	"os"
)

const (
	controller = "1"
	usecase    = "2"
	repo       = "3"
)

func main() {
	if len(os.Args) < 3 {
		log.Println("引数がたりません")
		os.Exit(9)
	}

	cgs := genarator.CreateGenerator{
		Fn: os.Args[2],
		In: os.Args[3],
	}

	ts := os.Args[1]
	var err error

	switch ts {
	case controller:
		err = cg.CreateController(&cgs)
	case usecase:
		err = ug.CreateUsecase(&cgs)
	case repo:
		err = rg.CreateRepository(&cgs)
	}

	if err != nil {
		log.Fatal(err.Error())
	}
}