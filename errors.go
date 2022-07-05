package backendtechtest

import "errors"

var (
	ErrNotFound     = errors.New("E001")
	ErrNotAvailable = errors.New("E002")
)
