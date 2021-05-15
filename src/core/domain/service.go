package domain

type ServiceRegistration struct {
	Name        string `json:"name"`
	ExternalUrl string `json:"external_url"`
}

//TODO: advertise if service is http(s)
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
