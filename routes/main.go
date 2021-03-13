package routes

import (
	"net/http"
	"struck/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())

	// client := r.Group("/category")
	// {

	// }

	r.POST("/login", controllers.Login)

	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	r.Run() // ":3000"
}
