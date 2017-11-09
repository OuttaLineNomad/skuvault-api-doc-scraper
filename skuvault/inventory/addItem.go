package inventory

type addItem struct {
	Code         string ``
	LocationCode string ``
	Quantity     int64  ``
	Reason       string ``
	Sku          string ``
	TenantToken  string ``
	UserToken    string ``
	WarehouseID  int64  ``
}


type addItemResponse struct {
	AddItemStatus string ``
}
