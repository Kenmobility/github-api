package message

import "errors"

var (
	ErrNoRecordFound = errors.New("no record found")
	ErrNoDataFound   = errors.New("no data found")

	ErrRepositoryNotFound = errors.New("passed repository does not exist")
)
