package usecase

import (
	"context"
	"custom-webhook-store-logs/external/domain"
)

type CachingActivityLogFromAuth0UseCase struct {
	activityLogRepo domain.ActivityLogRepo
}

func NewCachingActivityLogUseCase(activityLogRepo domain.ActivityLogRepo) CachingActivityLogFromAuth0UseCase {
	return CachingActivityLogFromAuth0UseCase{
		activityLogRepo: activityLogRepo,
	}
}

func (u *CachingActivityLogFromAuth0UseCase) CacheActivityLogs(ctx context.Context, activityLogs []domain.ActivityLog) error {
	return u.activityLogRepo.SaveLogs(ctx, activityLogs)
}
