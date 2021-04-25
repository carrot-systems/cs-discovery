package rest

import (
	"fmt"
	"github.com/carrot-systems/cs-discovery/src/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rH RoutesHandler) RegisterService(c *gin.Context) {
	var serviceRegistration *domain.ServiceRegistration
	err := c.ShouldBindJSON(&serviceRegistration)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.RegisterResponse{
			Status: domain.Status{
				Success: false,
				Message: err.Error(),
			},
		})
		return
	}

	err = rH.Usecases.RegisterService(*serviceRegistration)

	if err != nil {
		switch err {
		case domain.ErrServiceNameIsEmpty, domain.ErrExternalUrlsEmpty:
			c.AbortWithStatusJSON(http.StatusBadRequest, domain.RegisterResponse{
				Status: domain.Status{
					Success: false,
					Message: err.Error(),
				},
			})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, domain.RegisterResponse{
				Status: domain.Status{
					Success: false,
					Message: err.Error(),
				},
			})
		}
		return
	}

	c.JSON(http.StatusOK, domain.RegisterResponse{
		Status: domain.Status{
			Success: true,
			Message: fmt.Sprintf("service %s registered", serviceRegistration.Name),
		},
	})
}

func (rH RoutesHandler) GetAllServices(c *gin.Context) {
	services, err := rH.Usecases.GetAllServices()

	if err != nil {
		var code int

		switch err {
		case domain.ErrServiceNotFound:
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		c.AbortWithStatusJSON(code, domain.DiscoverResponse{
			Status: domain.Status{
				Success: false,
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, domain.DiscoverResponse{
		Status: domain.Status{
			Success: true,
		},
		Services: services,
	})
}

func (rH RoutesHandler) GetService(c *gin.Context) {
	appname := c.Param("appname")

	if appname == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.DiscoverResponse{
			Status: domain.Status{
				Success: false,
				Message: domain.ErrServiceNameIsEmpty.Error(),
			},
			Services: nil,
		})
		return
	}

	services, err := rH.Usecases.GetService(appname)

	if err != nil {
		var code int

		switch err {
		case domain.ErrServiceNotFound:
			code = http.StatusNotFound
		default:
			code = http.StatusInternalServerError
		}

		c.AbortWithStatusJSON(code, domain.DiscoverResponse{
			Status: domain.Status{
				Success: false,
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, domain.DiscoverResponse{
		Status: domain.Status{
			Success: true,
		},
		Services: []domain.Service{*services},
	})
}
