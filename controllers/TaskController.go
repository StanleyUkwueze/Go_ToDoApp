package controllers

import (
	"log"
	"todo/initializers"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func AddTask(c *gin.Context) {
	var task struct {
		Title       string
		IsCompleted bool
	}

	c.Bind(&task)

	taskToAdd := models.TaskModel{Title: task.Title, IsCompleted: task.IsCompleted}

	response := initializers.DB.Create(&taskToAdd)

	if response.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"task": taskToAdd,
	})
}

func FetchAllTasks(c *gin.Context) {
	var tasks []models.TaskModel

	initializers.DB.Find(&tasks)

	c.JSON(200, gin.H{
		"tasks": tasks,
	})
}

func FetchTaskById(c *gin.Context) {
	id := c.Param("id")
	var task models.TaskModel

	initializers.DB.Find(&task, id)

	c.JSON(200, gin.H{
		"task": task,
	})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var taskObj struct {
		Title string
	}

	c.Bind(&taskObj)

	var taskToUpdate models.TaskModel

	initializers.DB.Find(&taskToUpdate, id)

	initializers.DB.Model(&taskToUpdate).Updates(models.TaskModel{
		Title: taskObj.Title,
	})

	c.JSON(200, gin.H{
		"task": taskToUpdate,
	})
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.Find(&models.TaskModel{}, id)

	if result.RowsAffected == 0 {
		log.Fatal("no record found")
		return
	}
	initializers.DB.Delete(&models.TaskModel{}, id)
	if result.Error != nil {
		log.Fatal("Error occurred")
		c.Status(500)
		return
	}

	c.Status(200)

}

func CompleteTask(c *gin.Context) {
	id := c.Param("id")

	var taskObj struct {
		IsCompleted bool
	}

	c.Bind(&taskObj)

	var taskToUpdate models.TaskModel

	initializers.DB.Find(&taskToUpdate, id)

	initializers.DB.Model(&taskToUpdate).Updates(models.TaskModel{
		IsCompleted: taskObj.IsCompleted,
	})

	c.JSON(200, gin.H{
		"task": taskToUpdate,
	})
}

func FetchAllCompletedTasks(c *gin.Context) {
	var tasks []models.TaskModel

	initializers.DB.Where(&models.TaskModel{IsCompleted: true}).Find(&tasks)

	c.JSON(200, gin.H{
		"tasks": tasks,
	})
}
