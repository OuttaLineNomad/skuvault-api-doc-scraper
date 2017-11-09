package sales

type getSalesByDate []struct {
	Channel         string                ``
	ChannelID       string                ``
	Client          string                ``
	ContactInfo     getSalesByDate_sub1   ``
	FulfilledItems  []interface{}         ``
	FulfilledKits   []interface{}         ``
	ID              string                ``
	LastPrintedDate string                ``
	Marketplace     string                ``
	MarketplaceID   string                ``
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


type getSalesByDateResponse struct {
	Errors []interface{}                 ``
	Sales  []getSalesByDateResponse_sub5 ``
	Status string                        ``
}

type getSalesByDateResponse_sub2 struct {
	A int64  ``
	S string ``
}

type getSalesByDateResponse_sub4 struct {
	Address1   string ``
	Address2   string ``
	City       string ``
	Country    string ``
	PostalCode string ``
	Region     string ``
}

type getSalesByDateResponse_sub5 struct {
	Channel         string                        ``
	Client          string                        ``
	ContactInfo     getSalesByDateResponse_sub1   ``
	FulfilledItems  []interface{}                 ``
	FulfilledKits   []interface{}                 ``
	ID              string                        ``
	LastPrintedDate string                        ``
	Marketplace     string                        ``
	MerchantItems   []getSalesByDateResponse_sub3 ``
	MerchantKits    []interface{}                 ``
	Notes           string                        ``
	PrintedStatus   bool                          ``
	SaleDate        string                        ``
	ShippingCarrier string                        ``
	ShippingCharge  getSalesByDateResponse_sub2   ``
	ShippingClass   string                        ``
	ShippingCost    getSalesByDateResponse_sub2   ``
	ShippingInfo    getSalesByDateResponse_sub4   ``
	Status          string                        ``
}

type getSalesByDateResponse_sub1 struct {
	Company   string ``
	Email     string ``
	FirstName string ``
	LastName  string ``
	Phone     string ``
}

type getSalesByDateResponse_sub3 struct {
	PartNumber string                      ``
	Quantity   int64                       ``
	Sku        string                      ``
	UnitPrice  getSalesByDateResponse_sub2 ``
}
