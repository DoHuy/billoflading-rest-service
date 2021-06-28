package domain

import (
	"context"
	"time"
)

type DoSomethingWithVNSale interface {
	UpdateOrderState(ctx context.Context, state Trangthaivandon, id string) error
	FetchVandonhanhtrinh(ctx context.Context, mavandon string) (*Chitietvandon, error)
}

type Chitietvandon struct {
	ID            int         `json:"id"`
	ShipperID     string      `json:"shipper_id"`
	SellerID      interface{} `json:"seller_id"`
	PostID        int      `json:"post_id"`
	Post          string      `json:"post"`
	ProvinceID    string      `json:"province_id"`
	DistrictID    string      `json:"district_id"`
	WardID        string      `json:"ward_id"`
	State         int      `json:"state"`
	Time          int  `json:"time"`
	SourceAddress string      `json:"source_address"`
	DestAddress   string      `json:"dest_address"`
	SourcePhone   string      `json:"source_phone"`
	DestPhone     string      `json:"dest_phone"`
	Noted         string      `json:"noted"`
}

type Trangthaivandon struct {
	State       string     `json:"state"`
	Time        *time.Time `json:"time"`
	DestAddress string     `json:"dest_address"`
	DestPhone   string     `json:"dest_phone"`
}
