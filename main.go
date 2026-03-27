package main

import (
	"scaler/internal/server"
)

func main() {
	if err := server.SetupAndListen(); err != nil {
		panic(err.Error())
	}
}
