// main.go
package main

import (
	// "github.com/gin-gonic/gin"
	"Rest-Api/db"
	"Rest-Api/routes"
)

func main() {
	db.InitDB()
	r := routes.SetupRouter()
	r.Run(":8081")
}
