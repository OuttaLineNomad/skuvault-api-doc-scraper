package purchaseorders\n

type getIncomingItems struct {
	Errors    []interface{}           ``
	SoldItems []getIncomingItems_sub1 ``
	Status    string                  ``
}

type getIncomingItems_sub1 struct {
	Date       string ``
	Quantity   int64  ``
	Sku        string ``
	TotalPrice int64  ``
}
