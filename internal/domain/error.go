package domain

import "errors"

var (
	// ErrAllProvidesFailed is an error when all APIs failed
	ErrAllProvidesFailed = errors.New("all provides failed")
)
