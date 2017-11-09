package sales

type addShipments struct {
	Shipments   []addShipments_sub8 ``
	TenantToken string              ``
	UserToken   string              ``
}

type addShipments_sub7 struct {
	Address     addShipments_sub6 ``
	ThreePL     bool              ``
	WarehouseID string            ``
}

type addShipments_sub6 struct {
	Address1   string ``
	Address2   string ``
	City       string ``
	Company    string ``
	Country    string ``
	Email      string ``
	FirstName  string ``
	LastName   string ``
	MiddleName string ``
	PostalCode string ``
	Region     string ``
}

type addShipments_sub8 struct {
	AlternateID           string              ``
	Carrier               string              ``
	Class                 string              ``
	Costs                 []addShipments_sub1 ``
	EstimatedDeliveryDate string              ``
	EstimatedShipDate     string              ``
	LandedCosts           []addShipments_sub2 ``
	ManifestID            string              ``
	Note                  string              ``
	Parcels               []addShipments_sub5 ``
	SaleID                string              ``
	ShippedFrom           addShipments_sub7   ``
	Source                string              ``
	Status                string              ``
	TotalWeight           int64               ``
	TrackingNumber        string              ``
	TrackingURL           string              ``
	Type                  string              ``
	WeightUnit            string              ``
}

type addShipments_sub1 struct {
	Amount          int64  ``
	CostType        string ``
	CurrencyIsoCode string ``
}

type addShipments_sub2 struct {
	Amount          int64  ``
	CurrencyIsoCode string ``
	Sku             string ``
}

type addShipments_sub5 struct {
	Dimensions addShipments_sub3   ``
	Items      []addShipments_sub4 ``
	Note       string              ``
	Number     int64               ``
	Weight     int64               ``
	WeightUnit string              ``
}

type addShipments_sub3 struct {
	Height int64  ``
	Length int64  ``
	Unit   string ``
	Width  int64  ``
}

type addShipments_sub4 struct {
	Quantity int64  ``
	Sku      string ``
}


type addShipmentsResponse struct {
	Errors []interface{} ``
	Status string        ``
}
