package usecases

type DiscoveryCacheRepo interface {
}

type interactor struct {
	discoveryCacheRepo DiscoveryCacheRepo
}

func NewInteractor(dcR DiscoveryCacheRepo) interactor {
	return interactor{discoveryCacheRepo: dcR}
}
