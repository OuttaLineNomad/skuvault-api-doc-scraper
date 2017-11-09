package inventory

type setItemQuantities struct {
	Errors []setItemQuantities_sub1 ``
	Status string                   ``
}

type setItemQuantities_sub1 struct {
	Code  string ``
	Error string ``
	Sku   string ``
}


type setItemQuantityResponse struct {
	Code         string ``
	LocationCode string ``
	Quantity     int64  ``
	Sku          string ``
	TenantToken  string ``
	UserToken    string ``
	WarehouseID  int64  ``
}
