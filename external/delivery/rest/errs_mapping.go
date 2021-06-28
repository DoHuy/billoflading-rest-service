package rest

import (
	"net/http"

	"custom-webhook-store-logs/packages/errs"
	"custom-webhook-store-logs/packages/ginh"
)

var (
	ErrsMapDefault = ginh.ErrsMap{
		errs.ErrBadRequest:      ginh.ErrorMeta{Code: 400000, Message: "Bad request", HTTPStatus: http.StatusBadRequest},
		errs.UnErrAuthorized:    ginh.ErrorMeta{Code: 400001, Message: "un authorized", HTTPStatus: http.StatusUnauthorized},
		errs.ErrDatabase:        ginh.ErrorMeta{Code: 500001, Message: "query database have errors", HTTPStatus: http.StatusInternalServerError},
		errs.ErrInsertLogFailed: ginh.ErrorMeta{Code: 500002, Message: "insert record failed", HTTPStatus: http.StatusInternalServerError},
	}
)
