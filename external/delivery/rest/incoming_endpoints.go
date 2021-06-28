package rest

import (
	"custom-webhook-store-logs/external/domain"
	"custom-webhook-store-logs/packages/ginh"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CustomWebHooksExecutionSerializer struct {
	HookExecution domain.DoSomethingWithLoggingFromAuth0UseCase
}

func NewCustomWebHooksSerializer(uCase domain.DoSomethingWithLoggingFromAuth0UseCase) *CustomWebHooksExecutionSerializer {
	return &CustomWebHooksExecutionSerializer{
		HookExecution: uCase,
	}
}

// @Summary Store log from auth0
// @Description Hook store log from Auth0
// @Tags [store logs]
// @Accept json
// @Produce json
// @Param Body body rest.logsDTOReq true "Body message"
// @Success 201 {object} ginh.response{meta=ginh.ResponseMeta{code=int}}
// @failure 500 {object} ginh.response{meta=ginh.ResponseMeta{code=int}}
// @Router /web-hooks/logs/incoming [post]
func (s *CustomWebHooksExecutionSerializer) IncomingActivityLogsFromAuth0(ctx *gin.Context) {

	logs, err := s.buildActivityLogDomainFromLogDTO(ctx)
	if err != nil {
		ginh.LogError(ctx, "parse log from request failed ", err)
		ginh.BuildErrorResponse(ctx, err, nil)
		return
	}
	/*todo: save log to mysql */
	if err := s.HookExecution.CacheActivityLogs(ctx, logs); err != nil {
		ginh.LogError(ctx, "cache activity logs failed ", err)
		ginh.BuildErrorResponse(ctx, err, nil)
		return
	}
	ginh.LogInfo(ctx, "cache activity logs success ")
	ginh.BuildSuccessResponse(ctx, http.StatusCreated, nil)
}

func (s *CustomWebHooksExecutionSerializer) buildActivityLogDomainFromLogDTO(ctx *gin.Context) ([]domain.ActivityLog, error) {
	var logsDto logsDTOReq
	var activityLogs []domain.ActivityLog
	if err := ctx.ShouldBindJSON(&logsDto); err != nil {
		return nil, err
	}
	for _, log := range logsDto.Logs {
		date, err := time.Parse(time.RFC3339, log.Data.Date)
		if err != nil {
			return nil, err
		}
		activityLogs = append(activityLogs, domain.ActivityLog{
			LogID:       log.LogID,
			Type:        log.Data.Type,
			Description: log.Data.Description,
			ClientID:    log.Data.ClientID,
			ClientName:  log.Data.ClientName,
			IP:          log.Data.IP,
			UserAgent:   log.Data.UserAgent,
			UserID:      log.Data.UserID,
			Date:        &date,
		})
	}
	return activityLogs, nil

}
