package handlers

import "github.com/gin-gonic/gin"

type Person struct {
	Firstname string `json:"firstname,omitempty"`
}

type NewPerson struct {
	Name string `json:"name" binding:"required"`
}

type Error struct {
	Message string `json:"message"`
}

var People = []Person{{Firstname: "test"}}

func AddPeople(c *gin.Context) {
	var data NewPerson
	c.BindJSON(&data)
	if data.Name == "" {
		c.JSON(400, Error{
			Message: "get out of here dude",
		})
		return
	}
	People = append(People, Person{Firstname: data.Name})
	c.JSON(200, People)
	return
}
