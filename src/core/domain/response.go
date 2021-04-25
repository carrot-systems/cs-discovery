package domain

type Status struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type RegisterResponse struct {
	Status Status `json:"status"`
}

type DiscoverResponse struct {
	Status Status `json:"status"`

	Services []Service `json:"services",omitempty`
}
