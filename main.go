package main

import (
	"github.com/Ankcorn/hello/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", handlers.HydrateIndex)
	r.POST("/people", handlers.AddPeople)
	r.Run()
}
