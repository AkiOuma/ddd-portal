package common

import "errors"

var ErrInvalidRoleCode = errors.New("error: invalid role code")
var ErrInvalidResourceType = errors.New("error: invalid resource type")
var ErrInvalidResourceValue = errors.New("error: invalid resource value")
var ErrPermissionNotFound = errors.New("error: permission collection not found")
