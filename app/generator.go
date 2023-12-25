package main

import (
	"app/infra/genarator"
	"app/repositories/rg"
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

	cg := genarator.CreateGenerator{
		Fn: os.Args[2],
		In: os.Args[3],
	}

	ts := os.Args[1]
	var err error

	switch ts {
	case controller:

	case usecase:

	case repo:
		err = rg.CreateRepository(&cg)
	}

	if err != nil {
		log.Fatal(err.Error())
	}
}