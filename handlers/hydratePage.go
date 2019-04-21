package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HydrateIndex(c *gin.Context) {
	var names []string
	for i := 0; i < len(People); i++ {
		names = append(names, People[i].Firstname)
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":  "Main website",
		"values": names,
	})
}
