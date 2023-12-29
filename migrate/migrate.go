package main

import (
	"todo/initializers"
	"todo/models"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.TaskModel{})
}
