package main

import "github.com/CiscoCloud/shipped-utils/core/server"

const (
	Port = "8888"
)

func main() {
	server := server.InitRoutes()
	server.Run(":" + Port)
}
