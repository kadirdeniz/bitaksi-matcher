package pkg

import "errors"

var ErrDriverNotFound = errors.New("driver not found")
var ErrInvalidRequest = errors.New("invalid request")
var ErrInternalServer = errors.New("internal server error")
