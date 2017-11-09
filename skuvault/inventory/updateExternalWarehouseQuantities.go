package inventory

type updateExternalWarehouseQuantities struct {
	Quantities  []updateExternalWarehouseQuantities_sub1 ``
	TenantToken string                                   ``
	UserToken   string                                   ``
	WarehouseID string                                   ``
}

type updateExternalWarehouseQuantities_sub1 struct {
	InStockQuantity  int64  ``
	InboundQuantity  int64  ``
	ReserveQuantity  int64  ``
	Sku              string ``
	TotalQuantity    int64  ``
	TransferQuantity int64  ``
}


type updateExternalWarehouseQuantitiesResponse struct {
	Errors []string ``
	Status string   ``
}
