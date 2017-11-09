package products

type updateProduct struct {
	Attributes               updateProduct_sub1   ``
	Brand                    string               ``
	Classification           string               ``
	Code                     string               ``
	Cost                     int64                ``
	Description              string               ``
	LongDescription          string               ``
	MinimumOrderQuantity     int64                ``
	MinimumOrderQuantityInfo string               ``
	Note                     string               ``
	PartNumber               string               ``
	Pictures                 []string             ``
	ReorderPoint             int64                ``
	RetailPrice              int64                ``
	SalePrice                int64                ``
	ShortDescription         string               ``
	Sku                      string               ``
	Statuses                 []string             ``
	SupplierInfo             []updateProduct_sub2 ``
	TenantToken              string               ``
	UserToken                string               ``
	VariationParentSku       string               ``
	Weight                   int64                ``
	WeightUnit               string               ``
}

type updateProduct_sub1 struct {
	AttributeName string ``
}

type updateProduct_sub2 struct {
	Cost               int64  ``
	LeadTime           int64  ``
	SupplierName       string ``
	SupplierPartNumber string ``
	IsActive           bool   ``
	IsPrimary          bool   ``
}


type updateProductResponse struct {
	Errors []interface{} ``
	Status string        ``
}
