package products

type getSuppliers struct {
	PageNumber  int64  ``
	TenantToken string ``
	UserToken   string ``
}


type getSuppliersResponse struct {
	Supplier []getSuppliersResponse_sub1 ``
}

type getSuppliersResponse_sub1 struct {
	IsEnabled bool   ``
	Name      string ``
}
