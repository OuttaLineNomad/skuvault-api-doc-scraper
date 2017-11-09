package products

type getBrands struct {
	Brands []getBrands_sub1 ``
}

type getBrands_sub1 struct {
	IsEnabled bool   ``
	Name      string ``
}


type getClassificationsResponse struct {
	PageNumber  int64  ``
	TenantToken string ``
	UserToken   string ``
}
