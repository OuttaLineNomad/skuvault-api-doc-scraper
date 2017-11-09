package inventory

type setItemQuantity struct {
	SetItemQuantityStatusEnum string ``
}


type setShipmentFileResponse struct {
	Shipments   []setShipmentFileResponse_sub1 ``
	TenantToken string                         ``
	UserToken   string                         ``
}

type setShipmentFileResponse_sub1 struct {
	Carrier        string ``
	FilesData      string ``
	SaleID         string ``
	TrackingNumber string ``
}
