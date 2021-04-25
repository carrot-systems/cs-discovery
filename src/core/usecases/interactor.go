package usecases

import "github.com/carrot-systems/cs-discovery/src/core/domain"

type DiscoveryCacheRepo interface {
	RegisterService(registration domain.ServiceRegistration) error
	FindAllServices() ([]domain.Service, error)
	FindService(name string) (*domain.Service, error)
}

type interactor struct {
	discoveryCacheRepo DiscoveryCacheRepo
}

func NewInteractor(dcR DiscoveryCacheRepo) interactor {
	return interactor{discoveryCacheRepo: dcR}
}
