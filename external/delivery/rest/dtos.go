package rest

import "time"

type ChitietvandonResp struct {
	ID            int         `json:"id"`
	ShipperID     string      `json:"shipper_id"`
	SellerID      interface{} `json:"seller_id"`
	PostID        string      `json:"post_id"`
	Post          string      `json:"post"`
	ProvinceID    string      `json:"province_id"`
	DistrictID    string      `json:"district_id"`
	WardID        string      `json:"ward_id"`
	State         string      `json:"state"`
	Time          *time.Time  `json:"time"`
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
