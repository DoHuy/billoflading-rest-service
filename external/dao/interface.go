package dao

import (
	"context"
	"vtp-apis/external/dao/models"
	"vtp-apis/external/domain"
)

type DAO interface {
	FetchByID(ctx context.Context, id int) (*models.Chitietdon, *models.Vandonhanhtrinh, error)
	UpdateOrderStatus(ctx context.Context, state domain.Trangthaivandon) error
}
