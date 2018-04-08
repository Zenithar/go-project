package repositories

import "github.com/pkg/errors"

var (
	// ErrNoResult is raised when database return no result
	ErrNoResult = errors.New("repository: no result for given query")
)
