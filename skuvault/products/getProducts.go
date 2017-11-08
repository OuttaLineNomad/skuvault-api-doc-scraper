package products\n

type getProducts struct {
	Errors   []interface{}      ``
	Products []getProducts_sub3 ``
}

type getProducts_sub3 struct {
	Attributes          []getProducts_sub1 ``
	Brand               string             ``
	Classification      string             ``
	Code                string             ``
	Cost                int64              ``
	CreatedDateUtc      string             ``
	Description         string             ``
	DisableQuantitySync bool               ``
	ID                  string             ``
	IncrementalQuantity int64              ``
	IsAlternateCode     bool               ``
	IsAlternateSKU      bool               ``
	LongDescription     string             ``
	Moq                 int64              ``
	MOQInfo             string             ``
	ModifiedDateUtc     string             ``
	PartNumber          string             ``
	Pictures            []string           ``
	QuantityAvailable   int64              ``
	QuantityInStock     int64              ``
	QuantityInbound     int64              ``
	QuantityIncoming    int64              ``
	QuantityOnHand      int64              ``
	QuantityOnHold      int64              ``
	QuantityPending     int64              ``
	QuantityPicked      int64              ``
	QuantityTotalFBA    int64              ``
	QuantityTransfer    int64              ``
	ReorderPoint        int64              ``
	RetailPrice         int64              ``
	SalePrice           int64              ``
	ShortDescription    string             ``
	Sku                 string             ``
	Statuses            []string           ``
	Supplier            string             ``
	SupplierInfo        []getProducts_sub2 ``
	VariationParentSku  string             ``
	WeightUnit          string             ``
	WeightValue         string             ``
}

type getProducts_sub2 struct {
	Cost               int64  ``
	IsActive           bool   ``
	IsPrimary          bool   ``
	LeadTime           int64  ``
	SupplierName       string ``
	SupplierPartNumber string ``
}

type getProducts_sub1 struct {
	Name  string ``
	Value string ``
}
