package inventory

type getWarehouses struct {
	PageNumber  int64  ``
	TenantToken string ``
	UserToken   string ``
}


type getWarehousesResponse struct {
	Warehouses []getWarehousesResponse_sub1 ``
}

type getWarehousesResponse_sub1 struct {
	Code string ``
	ID   string ``
}
