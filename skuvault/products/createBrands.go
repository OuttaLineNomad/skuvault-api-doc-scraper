package products

type createBrands struct {
	Brands      []createBrands_sub1 ``
	TenantToken string              ``
	UserToken   string              ``
}

type createBrands_sub1 struct {
	Name string ``
}


type createBrandsResponse struct {
	Errors []interface{} ``
	Status string        ``
}
