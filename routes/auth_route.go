package router

import (
	"github.com/gin-gonic/gin"
	"github.com/devanfer02/gokers/controllers"
	"github.com/devanfer02/gokers/services"
)

func InitRouteAuth(router *gin.Engine) {
	authController := controllers.AuthController{
		Router: router,
		Service: services.AuthService{},
	}

	authStudent := authController.Router.Group("/auth/student")

	authStudent.POST("/register", authController.RegisterStudent)
	authStudent.POST("/login", authController.LoginStudent)
	authStudent.POST("/logout", authController.LogoutStudent)

	authLecturer := router.Group("/auth/lecturer")
	authLecturer.POST("/register", authController.RegisterLecturer)
}