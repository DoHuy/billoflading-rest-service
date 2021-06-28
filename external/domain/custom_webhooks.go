package domain

import (
	"context"
	"time"
)

type DoSomethingWithLoggingFromAuth0UseCase interface {
	CacheActivityLogs(ctx context.Context, activityLogs []ActivityLog) error
}

type ActivityLog struct {
	LogID       string
	Type        string
	Description string
	ClientID    string
	ClientName  string
	IP          string
	UserAgent   string
	UserID      string
	Date        *time.Time
}

type ActivityLogRepo interface {
	SaveLogs(ctx context.Context, activityLogs []ActivityLog) error
}
