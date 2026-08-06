// Package router wires up HTTP routes and middleware for the application.
package router

import (
	"regexp"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"backend/internal/handler"
)

// localDevOrigin matches any http://localhost:<port> or http://127.0.0.1:<port>
// origin, so the CORS allowlist doesn't break every time Vite's dev server
// picks a different port because its default is already in use.
var localDevOrigin = regexp.MustCompile(`^http://(localhost|127\.0\.0\.1):\d+$`)

// New builds the Gin engine with all routes and middleware registered.
func New(taskHandler *handler.TaskHandler) *gin.Engine {
	r := gin.Default()

	// No reverse proxy in front by default, so don't trust X-Forwarded-For headers.
	// If nginx/ALB is later placed in front of this app, trust its address here
	// (e.g. r.SetTrustedProxies([]string{"127.0.0.1"})) so ClientIP() stays accurate.
	_ = r.SetTrustedProxies(nil)

	r.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return localDevOrigin.MatchString(origin) },
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
