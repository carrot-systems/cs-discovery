package rest

import (
	"github.com/carrot-systems/cs-discovery/src/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rH RoutesHandler) RegisterService(c *gin.Context) {
	var serviceRegistration *domain.ServiceRegistration
	err := c.ShouldBindJSON(&serviceRegistration)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = rH.Usecases.RegisterService(*serviceRegistration)

	//TODO: better error handling
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}

func (rH RoutesHandler) GetAllServices(c *gin.Context) {
	services, err := rH.Usecases.GetAllServices()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, services)
}

func (rH RoutesHandler) GetService(c *gin.Context) {
	appname := c.Param("appname")

	if appname == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "appname can't be empty"})
		return
	}

	services, err := rH.Usecases.GetService(appname)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, services)
}
