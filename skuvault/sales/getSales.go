package sales\n

type getSales struct {
	Sales []getSales_sub7 ``
}

type getSales_sub2 struct {
	A int64  ``
	S string ``
}

type getSales_sub6 struct {
	Address1   string ``
	Address2   string ``
	City       string ``
	Country    string ``
	PostalCode string ``
	Region     string ``
}

type getSales_sub7 struct {
	ChannelID       string          ``
	ContactInfo     getSales_sub1   ``
	FulfilledItems  []interface{}   ``
	FulfilledKits   []interface{}   ``
	ID              string          ``
	LastPrintedDate string          ``
	Marketplace     string          ``
	MarketplaceID   string          ``
	Notes           string          ``
	PrintedStatus   bool            ``
	SaleDate        string          ``
	SaleItems       []getSales_sub3 ``
	SaleKits        []getSales_sub5 ``
	ShippingCarrier string          ``
	ShippingCharge  getSales_sub2   ``
	ShippingClass   string          ``
	ShippingCost    getSales_sub2   ``
	ShippingInfo    getSales_sub6   ``
	Status          string          ``
}

type getSales_sub1 struct {
	Company   string ``
	Email     string ``
	FirstName string ``
	LastName  string ``
	Phone     string ``
}

type getSales_sub5 struct {
	KitItems  getSales_sub4 ``
	Quantity  int64         ``
	Sku       string        ``
	UnitPrice getSales_sub2 ``
}

type getSales_sub3 struct {
	Quantity  int64         ``
	Sku       string        ``
	UnitPrice getSales_sub2 ``
}

type getSales_sub4 struct{}
