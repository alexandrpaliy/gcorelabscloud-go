package gcore

import (
	"fmt"

	gcorecloud "github.com/alexandrpaliy/gcorelabscloud-go"
)

// ErrEndpointNotFound is the error when no suitable endpoint can be found
// in the user's catalog
type ErrEndpointNotFound struct{ gcorecloud.BaseError }

func (e ErrEndpointNotFound) Error() string {
	return "No suitable endpoint could be found in the service catalog."
}

// ErrInvalidAvailabilityProvided is the error when an invalid endpoint
// availability is provided
type ErrInvalidAvailabilityProvided struct{ gcorecloud.ErrInvalidInput }

func (e ErrInvalidAvailabilityProvided) Error() string {
	return fmt.Sprintf("Unexpected availability in endpoint query: %s", e.Value)
}

// ErrNoAuthURL is the error when the OS_AUTH_URL environment variable is not
// found
type ErrNoAuthURL struct{ gcorecloud.ErrInvalidInput }

func (e ErrNoAuthURL) Error() string {
	return "Environment variable OS_AUTH_URL needs to be set."
}

// ErrNoUsername is the error when the OS_USERNAME environment variable is not
// found
type ErrNoUsername struct{ gcorecloud.ErrInvalidInput }

func (e ErrNoUsername) Error() string {
	return "Environment variable OS_USERNAME needs to be set."
}

// ErrNoPassword is the error when the OS_PASSWORD environment variable is not
// found
type ErrNoPassword struct{ gcorecloud.ErrInvalidInput }

func (e ErrNoPassword) Error() string {
	return "Environment variable OS_PASSWORD needs to be set."
}
