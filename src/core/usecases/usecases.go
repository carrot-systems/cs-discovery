package usecases

import "github.com/carrot-systems/cs-discovery/src/core/domain"

type Usecases interface {
	RegisterService(registration domain.ServiceRegistration) error
	GetAllServices() ([]domain.Service, error)
	GetService(name string) (*domain.Service, error)
}
