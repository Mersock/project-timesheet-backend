package v1

import "errors"

const (
	_defaultInternalServerErr = "Internal server error"
	_defaultNotFound          = "Not found"
	_defaultBadReq            = "Bad request"
	_defaultConflict          = "Conflict"
	_defaultUnauthorized      = "Unauthorized"
)

var ErrDuplicateRow = errors.New("duplicate")
