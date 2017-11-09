package purchaseorders

type getPOs struct {
	PurchaseOrders []getPOs_sub3 ``
}

type getPOs_sub3 struct {
	ActualShippedDate    string        ``
	ArrivalDueDate       string        ``
	Costs                []interface{} ``
	CreatedDate          string        ``
	LineItems            []getPOs_sub1 ``
	OrderCancelDate      string        ``
	OrderDate            string        ``
	PaymentStatus        string        ``
	PoNumber             string        ``
	PrivateNotes         string        ``
	PublicNotes          string        ``
	RequestedShipDate    string        ``
	ShipToAddress        string        ``
	ShipToWarehouse      string        ``
	ShippingCarrierClass getPOs_sub2   ``
	Status               string        ``
	SupplierName         string        ``
	TermsName            string        ``
	TrackingInfo         string        ``
}

type getPOs_sub2 struct {
	CarrierName string ``
	ClassName   string ``
}

type getPOs_sub1 struct {
	Cost                  int64  ``
	Identifier            string ``
	PrivateNotes          string ``
	ProductID             string ``
	PublicNotes           string ``
	Quantity              int64  ``
	QuantityTo3PL         int64  ``
	ReceivedQuantity      int64  ``
	ReceivedQuantityTo3PL int64  ``
	RetailCost            int64  ``
	Sku                   string ``
	Variant               string ``
}


type getProductsResponse struct {
	ModifiedAfterDateTimeUtc  string   ``
	ModifiedBeforeDateTimeUtc string   ``
	PageNumber                int64    ``
	PageSize                  int64    ``
	ProductCodes              []string ``
	ProductSKUs               []string ``
	TenantToken               string   ``
	UserToken                 string   ``
}
