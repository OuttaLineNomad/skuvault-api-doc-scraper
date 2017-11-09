package inventory

type getExternalWarehouses struct {
	Warehouses []getExternalWarehouses_sub1 ``
}

type getExternalWarehouses_sub1 struct {
	Code string ``
	ID   string ``
}


type getIncomingItemsResponse struct {
	PageNumber  int64  ``
	TenantToken string ``
	UserToken   string ``
}
