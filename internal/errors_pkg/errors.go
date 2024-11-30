package errors_pkg

import "errors"

var ErrLessZero = errors.New("less than zero")
var ErrInvalidRequest = errors.New("invalid request")
