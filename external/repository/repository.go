package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"vtp-apis/external/dao"
	"vtp-apis/external/dao/models"
	"vtp-apis/external/domain"
)

type ElasticRepository struct {
	esDAO dao.DAO
}

func jsonToJson(obj1 interface{}, obj2 interface{}) error {
	raw, err := json.Marshal(obj1)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(raw, obj2); err != nil {
		return err
	}
	return nil
}

func NewESRepository(dao dao.DAO) *ElasticRepository {
	return &ElasticRepository{
		esDAO: dao,
	}
}

func (repo *ElasticRepository) extractModelToDomain(resources models.Chitietdon, vandon models.Vandonhanhtrinh) *domain.Chitietvandon {
	return &domain.Chitietvandon{
		ID:            vandon.OrderId,
		ShipperID:     fmt.Sprintf("%v", resources.Employee),
		SellerID:      resources.CusId,
		PostID:        vandon.OrderPostId,
		ProvinceID:    vandon.ProvinceCode,
		DistrictID:    vandon.DistrictCode,
		WardID:        vandon.WardsCode,
		State:         vandon.OrderStatus,
		Time:          resources.CreatedBy,
		SourceAddress: vandon.ReceiverAddress,
		SourcePhone:   vandon.ReceiverPhone,
		DestPhone:     vandon.SenderPhone,
		Noted:         vandon.OrderNote,
	}
}

func (repo *ElasticRepository) FetchByID(ctx context.Context, id int) (*domain.Chitietvandon, error) {
	chitietdon, vandonhanhtrinh, err := repo.esDAO.FetchByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return repo.extractModelToDomain(*chitietdon, *vandonhanhtrinh), err
}

type OracleRepository struct {
	oracleDAO dao.DAO
}

func NewOracleRepository(or *dao.OracleDAO) *OracleRepository {
	return &OracleRepository{
		oracleDAO: or,
	}
}
