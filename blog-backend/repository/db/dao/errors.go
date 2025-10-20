package dao

import "errors"

var (
	ErrRecordNotFound   = errors.New("record not found")
	ErrPermissionDenied = errors.New("permission denied")
)
