package main

import (
	"api/config"
)
import "api/server"

func main() {
	r := config.StartConfig()

	server.StartServer(r)
}
