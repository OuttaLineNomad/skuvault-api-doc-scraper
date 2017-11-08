package inventory\n

type getAvailableQuantities struct {
	Items []getAvailableQuantities_sub1 ``
}

type getAvailableQuantities_sub1 struct {
	AvailableQuantity       int64  ``
	IsAlternateSku          bool   ``
	LastModifiedDateTimeUtc string ``
	Sku                     string ``
}
