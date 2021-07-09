package usecase

import (
	"context"
	"strconv"
	"vtp-apis/external/domain"
	"vtp-apis/external/repository"
)

type VNSaleUserCase struct {
	esRepo     *repository.ElasticRepository
	oracleRepo *repository.OracleRepository
}

func (V VNSaleUserCase) UpdateOrderState(ctx context.Context, state domain.Trangthaivandon, id string) error {
	panic("implement me")
}

func (V VNSaleUserCase) FetchVandonhanhtrinh(ctx context.Context, mavandon string) (*domain.Chitietvandon, error) {
	code, _ := strconv.Atoi(mavandon)
	chitietvandon, err := V.esRepo.FetchByID(ctx, code)
	if err != nil {
		return nil, err
	}
	// extract to chi tiet order resp

	return &domain.Chitietvandon{
		ID:            chitietvandon.ID,
		ShipperID:     chitietvandon.ShipperID,
		SellerID:      chitietvandon.SellerID,
		PostID:        chitietvandon.PostID,
		Post:          chitietvandon.Post,
		ProvinceID:    chitietvandon.ProvinceID,
		DistrictID:    chitietvandon.DistrictID,
		WardID:        chitietvandon.WardID,
		State:         chitietvandon.State,
		Time:          chitietvandon.Time,
		SourceAddress: chitietvandon.SourceAddress,
		DestAddress:   chitietvandon.DestAddress,
		SourcePhone:   chitietvandon.SourcePhone,
	}, err
}

func NewVNSaleUserCase(es *repository.ElasticRepository, oracle *repository.OracleRepository) *VNSaleUserCase {
	return &VNSaleUserCase{
		esRepo:     es,
		oracleRepo: oracle,
	}
}
