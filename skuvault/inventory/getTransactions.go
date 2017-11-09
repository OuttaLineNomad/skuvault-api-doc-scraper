package inventory

type getTransactions struct {
	ExcludeTransactionReasons []string ``
	FromDate                  string   ``
	PageNumber                int64    ``
	PageSize                  int64    ``
	TenantToken               string   ``
	ToDate                    string   ``
	TransactionReasons        []string ``
	TransactionType           string   ``
	UserToken                 string   ``
	WarehouseID               string   ``
}


type getTransactionsResponse struct {
	Transactions []getTransactionsResponse_sub2 ``
}

type getTransactionsResponse_sub2 struct {
	Code              string                       ``
	Context           getTransactionsResponse_sub1 ``
	Location          string                       ``
	Quantity          int64                        ``
	QuantityAfter     int64                        ``
	QuantityBefore    int64                        ``
	ScannedCode       string                       ``
	Sku               string                       ``
	Title             string                       ``
	TransactionDate   string                       ``
	TransactionNote   string                       ``
	TransactionReason string                       ``
	TransactionType   string                       ``
	User              string                       ``
}

type getTransactionsResponse_sub1 struct {
	ID   string ``
	Type string ``
}
