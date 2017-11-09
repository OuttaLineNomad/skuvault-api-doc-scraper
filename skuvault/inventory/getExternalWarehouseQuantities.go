package inventory

type getExternalWarehouseQuantities struct {
	Errors     []string                              ``
	Quantities []getExternalWarehouseQuantities_sub1 ``
}

type getExternalWarehouseQuantities_sub1 struct {
	InStockQuantity  int64  ``
	InboundQuantity  int64  ``
	ReserveQuantity  int64  ``
	Sku              string ``
	TotalQuantity    int64  ``
	TransferQuantity int64  ``
}


type getExternalWarehousesResponse struct {
	TenantToken string ``
	UserToken   string ``
}
