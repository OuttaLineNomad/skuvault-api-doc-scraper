package products

type createSuppliers struct {
	Errors []string ``
}


type getAvailableQuantitiesResponse struct {
	ExpandAlternateSkus       bool        ``
	ModifiedAfterDateTimeUtc  string      ``
	ModifiedBeforeDateTimeUtc string      ``
	PageNumber                int64       ``
	PageSize                  interface{} ``
	TenantToken               string      ``
	UserToken                 string      ``
}
