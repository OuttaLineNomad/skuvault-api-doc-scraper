package inventory

type getInventoryByLocation struct {
	Errors []string                    ``
	Items  getInventoryByLocation_sub2 ``
}

type getInventoryByLocation_sub1 struct {
	LocationCode  string ``
	Quantity      int64  ``
	Reserve       bool   ``
	WarehouseCode string ``
}

type getInventoryByLocation_sub2 struct {
	SKU_as_Key []getInventoryByLocation_sub1 ``
}


type getItemQuantitiesResponse struct {
	ModifiedAfterDateTimeUtc  string      ``
	ModifiedBeforeDateTimeUtc string      ``
	PageNumber                int64       ``
	PageSize                  interface{} ``
	ProductCodes              []string    ``
	TenantToken               string      ``
	UserToken                 string      ``
}
