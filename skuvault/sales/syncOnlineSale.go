package sales

type syncOnlineSale struct {
	OrderID string ``
	Status  string ``
}


type syncOnlineSalesResponse struct {
	Sales       []syncOnlineSalesResponse_sub3 ``
	TenantToken string                         ``
	Usertoken   string                         ``
}

type syncOnlineSalesResponse_sub3 struct {
	CheckoutStatus string                         ``
	FulfilledItems []syncOnlineSalesResponse_sub1 ``
	ItemSkus       []syncOnlineSalesResponse_sub1 ``
	MarketplaceID  string                         ``
	Notes          string                         ``
	OrderDateUtc   string                         ``
	OrderID        string                         ``
	OrderTotal     int64                          ``
	PaymentStatus  string                         ``
	SaleState      string                         ``
	ShippingInfo   syncOnlineSalesResponse_sub2   ``
}

type syncOnlineSalesResponse_sub2 struct {
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
	ShippingStatus  string ``
}

type syncOnlineSalesResponse_sub1 struct {
	Quantity  int64  ``
	Sku       string ``
	UnitPrice int64  ``
}
