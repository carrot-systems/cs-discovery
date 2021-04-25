package usecases

import (
	"errors"
	"github.com/carrot-systems/cs-discovery/src/core/domain"
)

func (i interactor) RegisterService(registration domain.ServiceRegistration) error {
	if registration.Name == "" {
		return errors.New("name can't be empty")
	}

	if registration.ExternalUrl == "" {
		return errors.New("external_url can't be empty")
	}

	return i.discoveryCacheRepo.RegisterService(registration)
}

func (i interactor) GetAllServices() ([]domain.Service, error) {
	return i.discoveryCacheRepo.FindAllServices()
}

func (i interactor) GetService(name string) (*domain.Service, error) {
	return i.discoveryCacheRepo.FindService(name)
}
