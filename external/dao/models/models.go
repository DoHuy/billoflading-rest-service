package models

const (
	IndexChitietdon      = "chitiet_don"
	IndexVandonhanhtrinh = "vandon_hanhtrinh"
)

type Vandonhanhtrinh struct {
	OrderId            int         `json:"order_id"`
	OrderNumber        string      `json:"order_number"`
	OrderReference     string      `json:"order_reference"`
	GroupaddressId     int         `json:"groupaddress_id"`
	CusId              int         `json:"cus_id"`
	Partner            int         `json:"partner"`
	DeliveryDate       int64       `json:"delivery_date"`
	DeliveryEmployer   interface{} `json:"delivery_employer"`
	SenderFullname     string      `json:"sender_fullname"`
	SenderAddress      string      `json:"sender_address"`
	SenderPhone        string      `json:"sender_phone"`
	SenderEmail        interface{} `json:"sender_email"`
	SenderWard         int         `json:"sender_ward"`
	SenderDistrict     int         `json:"sender_district"`
	SenderProvince     int         `json:"sender_province"`
	SenderPostId       interface{} `json:"sender_post_id"`
	SenderEmployer     interface{} `json:"sender_employer"`
	SenderLatitude     int         `json:"sender_latitude"`
	SenderLongitude    int         `json:"sender_longitude"`
	ReceiverFullname   string      `json:"receiver_fullname"`
	ReceiverAddress    string      `json:"receiver_address"`
	ReceiverPhone      string      `json:"receiver_phone"`
	ReceiverEmail      interface{} `json:"receiver_email"`
	ReceiverWard       int         `json:"receiver_ward"`
	ReceiverDistrict   int         `json:"receiver_district"`
	ReceiverProvince   int         `json:"receiver_province"`
	ReceiverPostId     interface{} `json:"receiver_post_id"`
	ReceiverEmployer   interface{} `json:"receiver_employer"`
	ReceiverLatitude   int         `json:"receiver_latitude"`
	ReceiverLongitude  int         `json:"receiver_longitude"`
	ProductName        string      `json:"product_name"`
	ProductDescription interface{} `json:"product_description"`
	ProductQuantity    int         `json:"product_quantity"`
	ProductPrice       int         `json:"product_price"`
	ProductWeight      int         `json:"product_weight"`
	ProductType        string      `json:"product_type"`
	OrderPayment       int         `json:"order_payment"`
	OrderService       string      `json:"order_service"`
	OrderServiceAdd    string      `json:"order_service_add"`
	OrderVoucher       interface{} `json:"order_voucher"`
	OrderStatus        int         `json:"order_status"`
	OrderPostId        int         `json:"order_post_id"`
	OrderSystemdate    int64       `json:"order_systemdate"`
	OrderAcceptdate    interface{} `json:"order_acceptdate"`
	OrderSuccessdate   interface{} `json:"order_successdate"`
	OrderEmployer      int         `json:"order_employer"`
	OrderNote          string      `json:"order_note"`
	MoneyCollection    int         `json:"money_collection"`
	MoneyTotalfee      int         `json:"money_totalfee"`
	MoneyFeecod        int         `json:"money_feecod"`
	MoneyFeevas        int         `json:"money_feevas"`
	MoneyFeeinsurrance int         `json:"money_feeinsurrance"`
	MoneyFee           int         `json:"money_fee"`
	MoneyFeeother      int         `json:"money_feeother"`
	MoneyTotalvat      int         `json:"money_totalvat"`
	MoneyTotal         int         `json:"money_total"`
	OrderType          int         `json:"order_type"`
	PostCode           string      `json:"post_code"`
	SenderPostCode     string      `json:"sender_post_code"`
	ServiceName        string      `json:"service_name"`
	ProvinceCode       string      `json:"province_code"`
	DistrictName       string      `json:"district_name"`
	DistrictCode       string      `json:"district_code"`
	WardsCode          string      `json:"wards_code"`
	IsPending          int         `json:"is_pending"`
	OrderAction505     int         `json:"order_action_505"`
	FeeCollected       int         `json:"fee_collected"`
	CollectedName      interface{} `json:"collected_name"`
	CollectedAddress   interface{} `json:"collected_address"`
}

type Chitietdon struct {
	CusId         interface{} `json:"cus_id"`
	Partner       interface{} `json:"partner"`
	PartnerCode   string      `json:"partner_code"`
	OrderNumber   string      `json:"order_number"`
	OrderStatus   int         `json:"order_status"`
	TrackingTime  int64       `json:"tracking_time"`
	PostId        interface{} `json:"post_id"`
	PostCode      string      `json:"post_code"`
	TrackingNote  string      `json:"tracking_note"`
	CreatedBy     int         `json:"created_by"`
	Employee      int         `json:"employee"`
	Pushed        bool        `json:"pushed"`
	EmployeeName  interface{} `json:"employee_name"`
	EmployeePhone interface{} `json:"employee_phone"`
}
