package main

import (
	"github.com/carrot-systems/cs-discovery/src/adapters/cache"
	"github.com/carrot-systems/cs-discovery/src/adapters/rest"
	"github.com/carrot-systems/cs-discovery/src/config"
	"github.com/carrot-systems/cs-discovery/src/core/usecases"
	env "github.com/carrot-systems/csl-env"
)

func main() {
	env.LoadEnv()

	ginConfiguration := config.LoadGinConfiguration()

	var discoveryCache usecases.DiscoveryCacheRepo
	discoveryCache = cache.NewCacheRepository()

	usecasesHandler := usecases.NewInteractor(discoveryCache)

	restServer := rest.NewServer(ginConfiguration)
	routesHandler := rest.NewRouter(usecasesHandler)

	rest.SetRoutes(restServer.Router, routesHandler)
	restServer.Start()
}
