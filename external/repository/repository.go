package repository

import (
	"context"
	"custom-webhook-store-logs/external/dao"
	"custom-webhook-store-logs/external/dao/models"
	"custom-webhook-store-logs/external/domain"
)

type ActivityLogRepository struct {
	webHookDAO dao.WebHookDAOGorm
}

func NewActivityLogRepository(webHookDAO dao.WebHookDAOGorm) *ActivityLogRepository {
	return &ActivityLogRepository{
		webHookDAO: webHookDAO,
	}
}

func (repo *ActivityLogRepository) extractResourceToModel(resources []domain.ActivityLog) []models.ActivityLog {
	var activityLogs []models.ActivityLog
	for _, resource := range resources {
		activityLogs = append(activityLogs, models.ActivityLog{
			LogID:       resource.LogID,
			Type:        resource.Type,
			Description: resource.Description,
			ClientID:    resource.ClientID,
			ClientName:  resource.ClientName,
			IP:          resource.IP,
			UserID:      resource.UserID,
			UserAgent:   resource.UserAgent,
			CreatedAt:   resource.Date,
		})
	}
	return activityLogs
}

func (repo *ActivityLogRepository) SaveLogs(ctx context.Context, logs []domain.ActivityLog) error {
	return repo.webHookDAO.InsertLogs(ctx, repo.extractResourceToModel(logs))
}
