package sales

type getShipments struct {
	SaleIds     []string ``
	TenantToken string   ``
	UserToken   string   ``
}


type getShipmentsResponse struct {
	Errors    []interface{}                ``
	Shipments []getShipmentsResponse_sub12 ``
	Status    string                       ``
}

type getShipmentsResponse_sub1 struct {
	A int64  ``
	S string ``
}

type getShipmentsResponse_sub11 struct {
	Address       getShipmentsResponse_sub10 ``
	ThreePL       bool                       ``
	WarehouseCode string                     ``
}

type getShipmentsResponse_sub10 struct {
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

type getShipmentsResponse_sub12 struct {
	AlternateID           string                      ``
	Carrier               string                      ``
	Class                 string                      ``
	Costs                 []getShipmentsResponse_sub2 ``
	CreatedDate           string                      ``
	EstimatedDeliveryDate string                      ``
	EstimatedShipDate     string                      ``
	FileIds               []getShipmentsResponse_sub3 ``
	LandedCosts           []getShipmentsResponse_sub4 ``
	ManifestID            string                      ``
	Note                  string                      ``
	Parcels               []getShipmentsResponse_sub9 ``
	SaleID                string                      ``
	ShippedFrom           getShipmentsResponse_sub11  ``
	Source                string                      ``
	Status                string                      ``
	TotalWeight           int64                       ``
	TrackingNumber        string                      ``
	TrackingURL           string                      ``
	Type                  string                      ``
	WeightUnit            string                      ``
}

type getShipmentsResponse_sub2 struct {
	Cost     getShipmentsResponse_sub1 ``
	CostType string                    ``
}

type getShipmentsResponse_sub4 struct {
	Cost getShipmentsResponse_sub1 ``
	Sku  string                    ``
}

type getShipmentsResponse_sub9 struct {
	Dimensions getShipmentsResponse_sub5   ``
	Items      []getShipmentsResponse_sub6 ``
	Kits       []getShipmentsResponse_sub8 ``
	Note       string                      ``
	Number     int64                       ``
	Weight     int64                       ``
	WeightUnit string                      ``
}

type getShipmentsResponse_sub5 struct {
	Height int64  ``
	Length int64  ``
	Unit   string ``
	Width  int64  ``
}

type getShipmentsResponse_sub8 struct {
	ID       string                      ``
	Items    []getShipmentsResponse_sub7 ``
	Quantity int64                       ``
	Sku      string                      ``
}

type getShipmentsResponse_sub6 struct {
	ID       string ``
	Quantity int64  ``
	Sku      string ``
}

type getShipmentsResponse_sub7 struct {
	ID       string ``
	Quantity int64  ``
}

type getShipmentsResponse_sub3 struct{}
