package main

import (
	"todo/controllers"
	"todo/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}
func main() {
	r := gin.Default()

	r.POST("/create", controllers.AddTask)

	r.GET("/fetch", controllers.FetchAllTasks)

	r.GET("/get/:id", controllers.FetchTaskById)

	r.PUT("/update/:id", controllers.UpdateTask)

	r.DELETE("/delete/:id", controllers.Delete)

	r.PUT("/done/:id", controllers.CompleteTask)

	r.GET("completed", controllers.FetchAllCompletedTasks)

	//User methods

	r.POST("signup", controllers.SignUp)
	r.POST("login", controllers.Login)

	r.Run()
}
