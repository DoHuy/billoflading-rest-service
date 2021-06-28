package ginh

import (
	"errors"
	"net/http"

	"vtp-apis/packages/errs"

	"github.com/gin-gonic/gin"
)

var (
	internalErrMeta = ErrorMeta{HTTPStatus: http.StatusInternalServerError, Code: 500, Message: "Internal server error"}
	defaultErrsMaps = make(ErrsMap)
)

type ErrorMeta struct {
	HTTPStatus int
	Code       interface{}
	Message    string
}

type ErrsMap map[error]ErrorMeta

func (em ErrsMap) AssertErrorCodesUnique() error {
	codeMap := make(map[interface{}]struct{})
	for _, meta := range em {
		if _, ok := codeMap[meta.Code]; ok {
			return errors.New("duplicated error code")
		}
		codeMap[meta.Code] = struct{}{}
	}
	return nil
}

func RegisterDefaultErrorsMap(errsMap ErrsMap) {
	defaultErrsMaps = errsMap
}

func BuildSuccessResponse(ctx *gin.Context, status int, body interface{}) {
	BuildStandardResponse(ctx, status, body, ResponseMeta{Code: int64(status)})
}

func BuildStandardResponse(ctx *gin.Context, status int, body interface{}, meta interface{}) {
	ctx.JSON(status, response{Data: body, Meta: meta})
}

func BuildErrorResponse(ctx *gin.Context, err error, body interface{}) {
	BuildErrorResponseWithErrorsMaps(ctx, err, body)
}

func BuildErrorResponseWithErrorsMaps(ctx *gin.Context, err error, body interface{}, errsMaps ...ErrsMap) {
	errsMaps = append(errsMaps, defaultErrsMaps)
	rootErr := errs.Cause(err)
	// convert error using errors maps
	errMeta := mapToErrMeta(rootErr, errsMaps...)
	ctx.JSON(errMeta.HTTPStatus, response{
		Data: body,
		Meta: ResponseMeta{
			Code:    errMeta.Code,
			Message: errMeta.Message,
		},
	})
}

func mapToErrMeta(err error, errsMap ...ErrsMap) ErrorMeta {
	for _, errsMap := range errsMap {
		if e, ok := errsMap[err]; ok {
			return e
		}
	}
	return internalErrMeta
}

type response struct {
	Data interface{} `json:"data,omitempty"`
	Meta interface{} `json:"meta,omitempty"`
}

type ResponseMeta struct {
	Code    interface{} `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
}

type DetailErrorsMeta struct {
	ResponseMeta
	DetailErrors []detailErrorResponse `json:"errors,omitempty"`
}

type detailErrorResponse struct {
	Code    interface{} `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
}
