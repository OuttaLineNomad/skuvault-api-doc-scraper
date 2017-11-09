package products

type updateHandlingTime struct {
	Errors []interface{} ``
	Status string        ``
}


type updateOnlineSaleStatusResponse struct {
	SaleID      string ``
	Status      string ``
	TenantToken string ``
	UserToken   string ``
}
