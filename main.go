package main

import (
	"todo/controllers"
	"todo/initializers"
	"todo/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}
func main() {
	r := gin.Default()

	r.POST("/create", middleware.RequireAuth, controllers.AddTask)

	r.GET("/fetch", controllers.FetchAllTasks)

	r.GET("/get/:id", controllers.FetchTaskById)

	r.PUT("/update/:id", middleware.RequireAuth, controllers.UpdateTask)

	r.DELETE("/delete/:id", middleware.RequireAuth, controllers.Delete)

	r.PUT("/done/:id", middleware.RequireAuth, controllers.CompleteTask)

	r.GET("completed", controllers.FetchAllCompletedTasks)

	//User Auth methods
	r.POST("signup", controllers.SignUp)
	r.POST("login", controllers.Login)
	r.GET("validate", middleware.RequireAuth, controllers.Validate)
	r.GET("getToken", controllers.GetTokenFromRequest)

	//User fetches
	r.GET("users", controllers.GetAllUsers)
	r.GET("mytasks", middleware.RequireAuth, controllers.GetUserTasks)

	r.Run()
}
