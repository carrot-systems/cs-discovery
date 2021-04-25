package domain

type ServiceRegistration struct {
	Name        string `json:"name"`
	ExternalUrl string `json:"external_url"`
}

type Service struct {
	Name        string `json:"name"`
	ExternalUrl string `json:"external_url"`
}

func RegistrationToService(registration ServiceRegistration) Service {
	return Service{
		Name:        registration.Name,
		ExternalUrl: registration.ExternalUrl,
	}
}
