package inventory

type getItemQuantities struct {
	Items []getItemQuantities_sub1 ``
}

type getItemQuantities_sub1 struct {
	AvailableQuantity       int64  ``
	Code                    string ``
	HeldQuantity            int64  ``
	LastModifiedDateTimeUtc string ``
	PendingQuantity         int64  ``
	PickedQuantity          int64  ``
	Sku                     string ``
	TotalOnHand             int64  ``
}


type getKitQuantitiesResponse struct {
	ModifiedAfterDateTimeUtc  string ``
	ModifiedBeforeDateTimeUtc string ``
	PageNumber                int64  ``
	TenantToken               string ``
	UserToken                 string ``
}
