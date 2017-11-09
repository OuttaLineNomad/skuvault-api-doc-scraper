package sales

type getSoldItems struct {
	BreakDownKits bool   ``
	EndDateUtc    string ``
	StartDateUtc  string ``
	TenantToken   string ``
	UserToken     string ``
}


type getSoldItemsResponse struct {
	SoldItems []getSoldItemsResponse_sub1 ``
}

type getSoldItemsResponse_sub1 struct {
	Date       string ``
	Quantity   int64  ``
	Sku        string ``
	TotalPrice int64  ``
}
