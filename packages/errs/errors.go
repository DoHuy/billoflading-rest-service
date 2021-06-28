package errs

import (
	"errors"
)

var (
	ErrBadRequest   = errors.New("bad request")
	UnErrAuthorized = errors.New("un authorized")
	ErrDatabase     = errors.New("database error")
)

// DAO errors
var (
	ErrInsertLogFailed = errors.New("insert database failed")
)

func RemoveDuplicateError(errors []error) (result []error) {
	encountered := map[error]bool{}
	for _, err := range errors {
		rootErr := Cause(err)
		if encountered[rootErr] {
			continue
		}
		encountered[rootErr] = true
		result = append(result, err)
	}
	return
}

func Any(errors ...error) error {
	for _, err := range errors {
		if err != nil {
			return err
		}
	}
	return nil
}
