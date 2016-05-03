package main

import "github.com/vjscjp/api4/core/server"

const (
	Port = "8888"
)
 
func main() {
	server := server.InitRoutes()
	server.Run(":" + Port)
}
