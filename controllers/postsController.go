package controllers

import (
	"app/initializers"
	"app/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	var Body struct {
		Body  string
		Title string
	}

	c.Bind(&Body)

	post := models.Post{Title: Body.Title, Body: Body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}
	c.JSON(200, gin.H{
		"message": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func FindSinglePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var Body struct {
		Body  string
		Title string
	}

	c.Bind(&Body)

	var post models.Post
	initializers.DB.First(&post, id)
	initializers.DB.Model(&post).Updates(models.Post{Title: Body.Title, Body: Body.Body})

	c.JSON(200, gin.H{
		"post": post,
	})
}
