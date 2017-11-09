package products

type createKit struct {
	AllowCreateAp bool             ``
	Code          string           ``
	KitLines      []createKit_sub1 ``
	Sku           string           ``
	TenantToken   string           ``
	Title         string           ``
	UserToken     string           ``
}

type createKit_sub1 struct {
	Combine  int64    ``
	Items    []string ``
	LineName string   ``
	Quantity int64    ``
}


type createKitResponse struct {
	Errors  []string ``
	Success string   ``
}
