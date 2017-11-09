package products

type getClassifications struct {
	Classifications []getClassifications_sub2 ``
}

type getClassifications_sub2 struct {
	Attributes []getClassifications_sub1 ``
	IsEnabled  bool                      ``
	Name       string                    ``
}

type getClassifications_sub1 struct {
	IsEnabled  bool     ``
	IsRequired bool     ``
	Name       string   ``
	Values     []string ``
}


type getExternalWarehouseQuantitiesResponse struct {
	GetTotalFbaQuantities bool        ``
	PageNumber            int64       ``
	PageSize              interface{} ``
	TenantToken           string      ``
	Usertoken             string      ``
	WarehouseID           string      ``
}
