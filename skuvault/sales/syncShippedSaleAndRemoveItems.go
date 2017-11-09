package sales

type syncShippedSaleAndRemoveItems struct {
	FulfilledItems []syncShippedSaleAndRemoveItems_sub1 ``
	ItemSkus       []syncShippedSaleAndRemoveItems_sub1 ``
	Notes          string                               ``
	OrderDateUtc   string                               ``
	OrderID        string                               ``
	OrderTotal     int64                                ``
	ShippingInfo   syncShippedSaleAndRemoveItems_sub2   ``
	TenantToken    string                               ``
	UserToken      string                               ``
	WarehouseID    int64                                ``
}

type syncShippedSaleAndRemoveItems_sub2 struct {
	City            string ``
	CompanyName     string ``
	Country         string ``
	Email           string ``
	FirstName       string ``
	LastName        string ``
	Line1           string ``
	Line2           string ``
	PhoneNumber     string ``
	Postal          string ``
	Region          string ``
	ShippingCarrier string ``
	ShippingClass   string ``
}

type syncShippedSaleAndRemoveItems_sub1 struct {
	Quantity  int64  ``
	Sku       string ``
	UnitPrice int64  ``
}


type syncShippedSaleAndRemoveItemsResponse struct {
	RemoveItemErrors []string ``
}
