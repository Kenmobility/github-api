package message

import "errors"

var (
	ErrNoRecordFound = errors.New("no record found")
	ErrNoDataFound   = errors.New("no data found")
	ErrInvalidInput  = errors.New("invalid input")

	ErrRepositoryNotFound      = errors.New("passed repository does not exist")
	ErrNoTrackingRepositorySet = errors.New("no repository set to track")
	ErrInvalidRepositoryName   = errors.New("invalid repository name, eg format is {owner/repositoryName}")
)
