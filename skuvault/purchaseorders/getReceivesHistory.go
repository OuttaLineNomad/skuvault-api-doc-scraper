package purchaseorders\n

type getReceivesHistory struct {
	Corrections []getReceivesHistory_sub2 ``
	Receives    []getReceivesHistory_sub3 ``
}

type getReceivesHistory_sub2 struct {
	Code                  string                    ``
	CorrectedDate         string                    ``
	NewQuantity           int64                     ``
	NewQuantity3pl        int64                     ``
	OldQuantity           int64                     ``
	OldQuantity3pl        int64                     ``
	PONumber              string                    ``
	PartNumber            string                    ``
	QuantityByReceiptDate []getReceivesHistory_sub1 ``
	ReceivedDate          string                    ``
	Sku                   string                    ``
	Username              string                    ``
}

type getReceivesHistory_sub3 struct {
	Code               string ``
	Location           string ``
	PONumber           string ``
	PartNumber         string ``
	Quantity           int64  ``
	Quantity3pl        int64  ``
	QuantityToLocation int64  ``
	ReceiptDate        string ``
	ReceivedDate       string ``
	Sku                string ``
	Username           string ``
	Warehouse          string ``
}

type getReceivesHistory_sub1 struct {
	NewQuantity int64  ``
	OldQuantity int64  ``
	ReceiptDate string ``
}
