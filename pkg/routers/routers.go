package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/suvam720/crud-api/pkg/controllers"
)

func Routes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/users", controllers.FindUser)
	r.POST("/users", controllers.CreateUser)
	r.DELETE("/users", controllers.DeleteAllUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.PATCH("/users/:id", controllers.UpdateUser)

	return r
}
