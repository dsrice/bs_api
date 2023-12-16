package main

import "app/infra/server"

func main() {
	s := server.NewServer()

	s.Start()
}
