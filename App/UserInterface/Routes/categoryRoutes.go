package routes

import (
	"github.com/gin-gonic/gin"
)

var RegisterCategoryRouter = func(router *gin.Engine) {
	router.GET("/api/health", func(w *gin.Context) {
		w.JSON(200, map[string]string{"message": "hello world"})
	})
	//router.HandleFunc("/book/test", controllers.Test).Methods("POST")
}
