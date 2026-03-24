package core

import "errors"

var (
	ErrBadAlgorithm   = errors.New("unsupported algorithm")
	ErrKeyNotFound    = errors.New("key not found")
	ErrKeyDisabled    = errors.New("key disabled")
	ErrInvalidVersion = errors.New("invalid version")
)
