package usecases

import (
	"github.com/carrot-systems/cs-discovery/src/core/domain"
)

func (i interactor) RegisterService(registration domain.ServiceRegistration) error {
	if registration.Name == "" {
		return domain.ErrServiceNameIsEmpty
	}

	if registration.ExternalUrl == "" {
		return domain.ErrExternalUrlsEmpty
	}

	return i.discoveryCacheRepo.RegisterService(registration)
}

func (i interactor) GetAllServices() ([]domain.Service, error) {
	return i.discoveryCacheRepo.FindAllServices()
}

func (i interactor) GetService(name string) (*domain.Service, error) {
	return i.discoveryCacheRepo.FindService(name)
}
