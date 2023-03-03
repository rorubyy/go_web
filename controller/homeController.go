package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/entity"
)

func HomePageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "test go"})
}

func Home(c *gin.Context) {
	data := new(entity.User)
	data.Age = 18
	data.Introduce = "Hi! I'm Iron Man!!!"
	data.Name = "Iron Man"
	c.HTML(http.StatusOK, "index.html", data)
}
