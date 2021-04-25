package domain

import "errors"

var (
	ErrServiceNameIsEmpty = errors.New("name can't be empty")
	ErrExternalUrlsEmpty  = errors.New("external_url can't be empty")
	ErrServiceNotFound    = errors.New("service not found")
)
