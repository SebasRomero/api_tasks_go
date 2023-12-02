package tasks

import (
	"errors"
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

func getTaskById(id string) (*Task, error) {
	for i, task := range Tasks {
		if task.ID == id {
			return &Tasks[i], nil
		}
	}

	return nil, errors.New("Task not found")
}

func GetTask(context *gin.Context) {
	id := context.Param("id")
	task, err := getTaskById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, task)
}

func UpdateTask(context *gin.Context) {
	id := context.Param("id")
	task, err := getTaskById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	task.Completed = !task.Completed
	context.IndentedJSON(http.StatusOK, task)
}

func DeleteTask(context *gin.Context) {
	id := context.Param("id")
	task, err := getTaskById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	Tasks = removeById(Tasks, task.ID)

	context.IndentedJSON(http.StatusAccepted, Tasks)
}

func removeById(tasks []Task, id string) []Task {
	index := -1
	for i, element := range tasks {
		if element.ID == id {
			index = i
			break
		}
	}

	if index != -1 {
		tasks = append(tasks[:index], tasks[index+1:]...)
	}

	return tasks
}
