package app

import "errors"

var (
	ErrInternal                        = errors.New("internal error")
	ErrAgentNotFound                   = errors.New("agent was not found in the repository")
	ErrCustomerNotFound                = errors.New("customer was not found in the repository")
	ErrCustomerAgentConnectionNotFound = errors.New("no connection found between the given customer and agent")
)
