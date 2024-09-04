package main

import (
	"basiccrud/internal/config"
	"basiccrud/internal/di"
	"log"
)

func main() {
	cnf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load environements")
	}
	server := di.Init(cnf)
	server.StartServer()
}
