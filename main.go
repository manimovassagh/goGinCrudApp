package main

import (
	"app/controllers"
	"app/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.FindSinglePost)
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run() // listen and serve on 127.0.0.1:3000 or http://localhost:3000

}
