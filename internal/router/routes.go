package router

import (
	"challenge/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *handlers.Handler) {
	r.POST("/upload-courses", h.UploadCourses)
	r.GET("/courses", h.GetCourses)
}
