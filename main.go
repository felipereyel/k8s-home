package main

import (
	"k8s-home/internal/server"
)

func main() {
	if err := server.SetupAndListen(); err != nil {
		panic(err.Error())
	}
}
