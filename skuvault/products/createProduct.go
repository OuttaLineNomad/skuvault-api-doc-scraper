package products

type createProduct struct {
	AllowCreateAp            bool                 ``
	Attributes               createProduct_sub1   ``
	Brand                    string               ``
	Classification           string               ``
	Code                     interface{}          ``
	Cost                     int64                ``
	Description              string               ``
	LongDescription          string               ``
	MinimumOrderQuantity     int64                ``
	MinimumOrderQuantityInfo string               ``
	Note                     string               ``
	PartNumber               string               ``
	Pictures                 []string             ``
	ReorderPoint             int64                ``
	RetailPrice              string               ``
	SalePrice                int64                ``
	ShortDescription         string               ``
	Sku                      string               ``
	Statuses                 []string             ``
	SupplierInfo             []createProduct_sub2 ``
	TenantToken              string               ``
	UserToken                string               ``
	VariationParentSku       string               ``
	Weight                   int64                ``
	WeightUnit               string               ``
}

type createProduct_sub1 struct {
	AttributeName string ``
}

type createProduct_sub2 struct {
	Cost               float64 ``
	LeadTime           int64   ``
	SupplierName       string  ``
	SupplierPartNumber string  ``
	IsActive           bool    ``
	IsPrimary          bool    ``
}


type createProductResponse struct {
	Errors []string ``
	Status string   ``
}
