package inventory

type getWarehouseItemQuantities struct {
	PageNumber  int64  ``
	PageSize    int64  ``
	TenantToken string ``
	UserToken   string ``
	WarehouseID int64  ``
}


type getWarehouseItemQuantitiesResponse struct {
	ItemQuantities []getWarehouseItemQuantitiesResponse_sub1 ``
}

type getWarehouseItemQuantitiesResponse_sub1 struct {
	Quantity int64  ``
	Sku      string ``
}
