package sales

type syncOnlineSales struct {
	Errors []syncOnlineSales_sub1 ``
	Status string                 ``
}

type syncOnlineSales_sub1 struct {
	ErrorMessages []string ``
	OrderID       string   ``
}


type syncOnlineSalesResponse struct {
	Errors []syncOnlineSalesResponse_sub1 ``
	Status string                         ``
}

type syncOnlineSalesResponse_sub1 struct {
	ErrorMessages []string ``
	OrderID       string   ``
}
