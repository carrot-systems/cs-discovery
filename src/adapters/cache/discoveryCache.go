package cache

import (
	"github.com/carrot-systems/cs-discovery/src/core/domain"
	"github.com/carrot-systems/cs-discovery/src/core/usecases"
)

type CacheRepository struct {
	services map[string]domain.Service
}

func (repo CacheRepository) RegisterService(registration domain.ServiceRegistration) error {
	service := domain.RegistrationToService(registration)
	repo.services[service.Name] = service

	return nil
}

func (repo CacheRepository) FindAllServices() ([]domain.Service, error) {
	services := []domain.Service{}

	for _, service := range repo.services {
		services = append(services, service)
	}

	return services, nil
}

func (repo CacheRepository) FindService(name string) (*domain.Service, error) {
	if val, ok := repo.services[name]; ok {
		return &val, nil
	}

	return nil, domain.ErrServiceNotFound
}

func NewCacheRepository() usecases.DiscoveryCacheRepo {
	return CacheRepository{
		services: map[string]domain.Service{},
	}
}
