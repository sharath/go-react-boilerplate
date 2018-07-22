package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func V1(router *gin.RouterGroup) {
	router.GET("/users", func(context *gin.Context) {
		users := []string{"user1", "user2", "user3", "user4"}
		context.JSON(http.StatusOK, users)
	})
}
