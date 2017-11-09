package inventory

type removeItem struct {
	Code         string ``
	LocationCode string ``
	Quantity     int64  ``
	Reason       string ``
	Sku          string ``
	TenantToken  string ``
	UserToken    string ``
	WarehouseID  int64  ``
}


type removeItemResponse struct {
	RemoveItemStatus string ``
}
