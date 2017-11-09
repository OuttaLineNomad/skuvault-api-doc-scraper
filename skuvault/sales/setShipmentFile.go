package sales

type setShipmentFile struct {
	Errors  []interface{}          ``
	FileIds []setShipmentFile_sub1 ``
	Status  string                 ``
}

type setShipmentFile_sub1 struct{}


type syncOnlineSaleResponse struct {
	CheckoutStatus string                        ``
	FulfilledItems []syncOnlineSaleResponse_sub1 ``
	ItemSkus       []syncOnlineSaleResponse_sub1 ``
	MarketplaceID  string                        ``
	Notes          string                        ``
	OrderDateUtc   string                        ``
	OrderID        string                        ``
	OrderTotal     int64                         ``
	PaymentStatus  string                        ``
	SaleState      string                        ``
	ShippingInfo   syncOnlineSaleResponse_sub2   ``
	TenantToken    string                        ``
	UserToken      string                        ``
}

type syncOnlineSaleResponse_sub2 struct {
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

type syncOnlineSaleResponse_sub1 struct {
	Quantity  int64  ``
	Sku       string ``
	UnitPrice int64  ``
}
