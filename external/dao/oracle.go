package dao

import (
	"context"
	"github.com/godror/godror"
	"vtp-apis/external/dao/models"
	"vtp-apis/external/domain"
)

const (
	TableName = "VTP_VANDONHANHTRINH"
)

type OracleDAO struct {
	client *godror.Conn
}

func (o OracleDAO) FetchByID(ctx context.Context, id int) (*models.Chitietdon, *models.Vandonhanhtrinh, error) {
	panic("implement me")
}

func (o OracleDAO) UpdateOrderStatus(ctx context.Context, state domain.Trangthaivandon) error {
	panic("implement me")
}

func NewOracleDAO(db *godror.Conn) DAO {
	return &OracleDAO{client: db}
}
