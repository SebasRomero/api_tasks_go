package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, Tasks)
}
