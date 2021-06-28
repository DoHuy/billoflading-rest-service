package dao

import (
	"context"
	"custom-webhook-store-logs/external/domain"
)

type WebHookLoggingDAO interface {
	InsertLogs(ctx context.Context, activityLogs []domain.ActivityLog) error
}
