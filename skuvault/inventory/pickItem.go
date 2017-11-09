package inventory

type pickItem struct {
	Code          string ``
	IsExpressPick bool   ``
	LocationCode  string ``
	Note          string ``
	Quantity      int64  ``
	Sku           string ``
	TenantToken   string ``
	UserToken     string ``
	WarehouseID   int64  ``
}


type pickItemResponse struct {
	PickItemStatus string ``
}
