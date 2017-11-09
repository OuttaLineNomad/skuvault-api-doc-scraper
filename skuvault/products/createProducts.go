package products

type createProducts struct {
	Errors []interface{} ``
	Status string        ``
}


type createSuppliersResponse struct {
	Suppliers   []createSuppliersResponse_sub1 ``
	TenantToken string                         ``
	UserToken   string                         ``
}

type createSuppliersResponse_sub1 struct {
	EmailTemplateMessage string   ``
	EmailTemplateSubject string   ``
	Emails               []string ``
	Name                 string   ``
}
