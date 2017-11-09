package inventory

type getKitQuantities struct {
	Kits []getKitQuantities_sub1 ``
}

type getKitQuantities_sub1 struct {
	AvailableQuantity       int64  ``
	LastModifiedDateTimeUtc string ``
	Sku                     string ``
}


type getKitsResponse struct {
	AvailableQuantityModifiedAfterDateTimeUtc  string   ``
	AvailableQuantityModifiedBeforeDateTimeUtc string   ``
	GetAvailableQuantity                       bool     ``
	IncludeKitCost                             bool     ``
	KitSKUs                                    []string ``
	ModifiedAfterDateTimeUtc                   string   ``
	ModifiedBeforeDateTimeUtc                  string   ``
	PageNumber                                 int64    ``
	TenantToken                                string   ``
	UserToken                                  string   ``
}
