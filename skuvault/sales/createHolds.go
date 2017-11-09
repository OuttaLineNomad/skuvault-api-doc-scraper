package sales

type createHolds struct {
	Holds       []createHolds_sub1 ``
	TenantToken string             ``
	UserToken   string             ``
}

type createHolds_sub1 struct {
	Description       string ``
	ExpirationDateUtc string ``
	Quantity          int64  ``
	Sku               string ``
}


type createHoldsResponse []struct {
	Quantity int64  ``
	Sku      string ``
}
