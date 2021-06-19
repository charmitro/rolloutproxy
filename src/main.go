package main

import (
	"github.com/charmitro/rolloutproxy/src/pkg/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	server.Init()
}
