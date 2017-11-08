package sales\n

type getSalesByDate struct {
	Errors []interface{}         ``
	Sales  []getSalesByDate_sub5 ``
	Status string                ``
}

type getSalesByDate_sub2 struct {
	A int64  ``
	S string ``
}

type getSalesByDate_sub4 struct {
	Address1   string ``
	Address2   string ``
	City       string ``
	Country    string ``
	PostalCode string ``
	Region     string ``
}

type getSalesByDate_sub5 struct {
	Channel         string                ``
	Client          string                ``
	ContactInfo     getSalesByDate_sub1   ``
	FulfilledItems  []interface{}         ``
	FulfilledKits   []interface{}         ``
	ID              string                ``
	LastPrintedDate string                ``
	Marketplace     string                ``
	MerchantItems   []getSalesByDate_sub3 ``
	MerchantKits    []interface{}         ``
	Notes           string                ``
	PrintedStatus   bool                  ``
	SaleDate        string                ``
	ShippingCarrier string                ``
	ShippingCharge  getSalesByDate_sub2   ``
	ShippingClass   string                ``
	ShippingCost    getSalesByDate_sub2   ``
	ShippingInfo    getSalesByDate_sub4   ``
	Status          string                ``
}

type getSalesByDate_sub1 struct {
	Company   string ``
	Email     string ``
	FirstName string ``
	LastName  string ``
	Phone     string ``
}

type getSalesByDate_sub3 struct {
	PartNumber string              ``
	Quantity   int64               ``
	Sku        string              ``
	UnitPrice  getSalesByDate_sub2 ``
}
