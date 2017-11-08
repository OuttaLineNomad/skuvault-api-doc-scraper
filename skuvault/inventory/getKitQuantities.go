package inventory\n

type getKitQuantities struct {
	Kits []getKitQuantities_sub1 ``
}

type getKitQuantities_sub1 struct {
	AvailableQuantity       int64  ``
	LastModifiedDateTimeUtc string ``
	Sku                     string ``
}
