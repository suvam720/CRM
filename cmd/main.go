package main

import "github.com/suvam720/crud-api/pkg/routers"

func main() {

	routes := routers.Routes()
	println("Server started on port 8080...")
	routes.Run("localhost:8080")
}
