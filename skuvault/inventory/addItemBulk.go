package inventory

type addItemBulk struct {
	Items       []addItemBulk_sub1 ``
	TenantToken string             ``
	UserToken   string             ``
}

type addItemBulk_sub1 struct {
	Code         string ``
	LocationCode string ``
	Quantity     int64  ``
	Reason       string ``
	Sku          string ``
	WarehouseID  int64  ``
}


type addItemBulkResponse struct {
	Errors  []addItemBulkResponse_sub1 ``
	Results []string                   ``
	Status  string                     ``
}

type addItemBulkResponse_sub1 struct {
	ErrorMessages []string ``
	Sku           string   ``
}
