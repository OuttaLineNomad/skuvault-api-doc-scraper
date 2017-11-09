package inventory

type getWarehouseItemQuantity struct {
	Sku         string ``
	TenantToken string ``
	UserToken   string ``
	WarehouseID int64  ``
}


type getWarehouseItemQuantityResponse struct {
	Errors              []string ``
	TotalQuantityOnHand int64    ``
}
