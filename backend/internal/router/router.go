// Package router wires up HTTP routes and middleware for the application.
package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"backend/internal/handler"
)

// New builds the Gin engine with all routes and middleware registered.
func New(taskHandler *handler.TaskHandler) *gin.Engine {
	r := gin.Default()

	// No reverse proxy in front by default, so don't trust X-Forwarded-For headers.
	// If nginx/ALB is later placed in front of this app, trust its address here
	// (e.g. r.SetTrustedProxies([]string{"127.0.0.1"})) so ClientIP() stays accurate.
	_ = r.SetTrustedProxies(nil)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
	{
		api.GET("/health", taskHandler.Health)

		tasks := api.Group("/tasks")
		{
			tasks.POST("", taskHandler.Create)
			tasks.GET("", taskHandler.List)
			tasks.GET("/:id", taskHandler.Get)
			tasks.PUT("/:id", taskHandler.Update)
			tasks.DELETE("/:id", taskHandler.Delete)
		}
	}

	return r
}
