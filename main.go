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

	r.POST("/posts", controllers.PostsCreate)

	r.Run() // listen and serve on 0.0.0.0:8080
}
