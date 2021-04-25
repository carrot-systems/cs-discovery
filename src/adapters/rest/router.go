package rest

import "github.com/carrot-systems/cs-discovery/src/core/usecases"

type RoutesHandler struct {
	Usecases usecases.Usecases
}

func NewRouter(ucHandler usecases.Usecases) RoutesHandler {
	return RoutesHandler{Usecases: ucHandler}
}
