package purchaseorders

type createPO struct {
	ArrivalDueDate       string          ``
	LineItems            []createPO_sub1 ``
	OrderCancelDate      string          ``
	OrderDate            string          ``
	PaymentStatus        string          ``
	Payments             []createPO_sub2 ``
	PoNumber             string          ``
	PrivateNotes         string          ``
	PublicNotes          string          ``
	RequestedShipDate    string          ``
	SentStatus           string          ``
	ShipToAddress        string          ``
	ShipToWarehouse      string          ``
	ShippingCarrierClass createPO_sub3   ``
	SupplierName         string          ``
	TenantToken          string          ``
	TermsName            string          ``
	TrackingInfo         string          ``
	UserToken            string          ``
}

type createPO_sub2 struct {
	Amount      int64  ``
	Note        string ``
	PaymentName string ``
}

type createPO_sub3 struct {
	CarrierName string ``
	ClassName   string ``
}

type createPO_sub1 struct {
	Cost          int64  ``
	Identifier    string ``
	PrivateNotes  string ``
	PublicNotes   string ``
	Quantity      int64  ``
	QuantityTo3PL int64  ``
	Sku           string ``
	Variant       string ``
}


type createPOResponse struct {
	CreatePOStatus string ``
}
