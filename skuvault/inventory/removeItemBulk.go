package inventory

type removeItemBulk struct {
	Items       []removeItemBulk_sub1 ``
	TenantToken string                ``
	UserToken   string                ``
}

type removeItemBulk_sub1 struct {
	Code         string ``
	LocationCode string ``
	Quantity     int64  ``
	Reason       string ``
	Sku          string ``
	WarehouseID  int64  ``
}


type removeItemBulkResponse struct {
	Errors  []string ``
	Results []string ``
	Status  string   ``
}
