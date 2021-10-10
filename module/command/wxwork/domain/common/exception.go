package common

import "errors"

var ErrSyncUser = errors.New("error: An error occurred while synchronizing users")

var ErrUnknownInternal = errors.New("error: unknown internal error")
