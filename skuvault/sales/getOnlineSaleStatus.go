package sales

type getOnlineSaleStatus struct {
	Sales []getOnlineSaleStatus_sub2 ``
}

type getOnlineSaleStatus_sub2 struct {
	ID              string                     ``
	Items           []getOnlineSaleStatus_sub1 ``
	LastPrintedDate string                     ``
	Notes           string                     ``
	PrintedStatus   bool                       ``
	Status          string                     ``
}

type getOnlineSaleStatus_sub1 struct {
	OnlineSaleItemStatus string ``
	Quantity             int64  ``
	Sku                  string ``
}


type getPOsResponse struct {
	IncludeProducts           bool     ``
	ModifiedAfterDateTimeUtc  string   ``
	ModifiedBeforeDateTimeUtc string   ``
	PONumbers                 []string ``
	PageNumber                int64    ``
	Status                    string   ``
	TenantToken               string   ``
	UserToken                 string   ``
}