package dao

import (
	"context"
	"github.com/elastic/go-elasticsearch/v6"
	"github.com/pkg/errors"
	"vtp-apis/external/dao/models"
	"vtp-apis/packages/errs"
)

const (
	chitietdon_index = "chitiet_don"
	vandonhanhtrinh_index= "vandon_hanhtrinh"
)

type ElasticsearchDAO struct {
	client *elasticsearch.Client
}

func NewElasticsearchDAO(db *elasticsearch.Client) *ElasticsearchDAO {
	return &ElasticsearchDAO{client: db}
}

func (s *ElasticsearchDAO) getDetailOrder(ctx context.Context, orderId string) error {
	errDb := s.db.WithContext(ctx).Create(&logs).Error
	if errDb != nil {
		return errors.Wrap(errDb, errs.WithStack(errs.ErrDatabase).Error())
	}
	return nil
}

func (s *ElasticsearchDAO) getThongtinVandonhanhtrinh(ctx context.Context, ) error {
	errDb := s.db.WithContext(ctx).Create(&logs).Error
	if errDb != nil {
		return errors.Wrap(errDb, errs.WithStack(errs.ErrDatabase).Error())
	}
	return nil
}