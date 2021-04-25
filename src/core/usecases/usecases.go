package usecases

type Usecases interface {
	RegisterService(name string, url string)
}
