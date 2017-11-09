package purchaseorders

type receivePOItems struct {
	LineItems    []receivePOItems_sub1 ``
	PoNumber     string                ``
	ReceiptDate  string                ``
	SupplierName string                ``
	TenantToken  string                ``
	UserToken    string                ``
	WarehouseID  int64                 ``
}

type receivePOItems_sub1 struct {
	Location      string ``
	Quantity      int64  ``
	QuantityTo3PL int64  ``
	Sku           string ``
}


type receivePOItemsResponse struct {
	Errors []string ``
	Status string   ``
}
