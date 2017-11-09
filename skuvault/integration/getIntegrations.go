package integration

type getIntegrations struct {
	Accounts []getIntegrations_sub1 ``
}

type getIntegrations_sub1 struct {
	ID     string ``
	LongID string ``
	Name   string ``
	Type   string ``
}


type getInventoryByLocationResponse struct {
	IsReturnByCodes bool          ``
	PageNumber      int64         ``
	PageSize        int64         ``
	ProductCodes    []interface{} ``
	ProductSKUs     []interface{} ``
	TenantToken     string        ``
	Usertoken       string        ``
}
