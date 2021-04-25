package rest

import "github.com/gin-gonic/gin"

func SetRoutes(r *gin.Engine, routesHandler RoutesHandler) {
	r.POST("/services", routesHandler.RegisterService)
	r.GET("/services", routesHandler.GetAllServices)
	r.GET("/services/:appname", routesHandler.GetService)
}
