package route

import (
	"test_sat/handler"
	"test_sat/pkg/midleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("login", handler.Login)
	r.GET("logout", handler.Logout)
	auth := r.Group("/v0")
	auth.Use(midleware.AuthMiddleware())
	{
		auth.GET("student", handler.FindStudent)
		auth.GET("student/:id", handler.FindStudentDetail)
		auth.POST("student", handler.CreateStudent)
		auth.PUT("student/:id", handler.UpdateStudent)
		auth.DELETE("student/:id", handler.DeleteStudent)

	}
	return r
}
