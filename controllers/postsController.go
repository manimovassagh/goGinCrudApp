package controllers

import (
	"app/initializers"
	"app/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	post := models.Post{Title: "Mani", Body: "Test The Body"}

	result := initializers.DB.Create(&post) // pass pointer of data to Create
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}
	c.JSON(200, gin.H{
		"message": post,	
	})
}
