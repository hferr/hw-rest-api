package app

import "errors"

var (
	ErrInternal      = errors.New("internal error")
	ErrAgentNotFound = errors.New("agent was not found in the repository")
)
