package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, Tasks)
}

func AddTask(context *gin.Context) {
	var newTask Task

	if err := context.BindJSON(&newTask); err != nil {
		return
	}

	Tasks = append(Tasks, newTask)

	context.IndentedJSON(http.StatusCreated, newTask)
}
