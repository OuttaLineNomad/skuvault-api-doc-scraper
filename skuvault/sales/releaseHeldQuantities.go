package sales

type releaseHeldQuantities struct {
	SkusToRelease releaseHeldQuantities_sub1 ``
	TenantToken   string                     ``
	UserToken     string                     ``
}

type releaseHeldQuantities_sub1 struct {
	SKUAsKey int64 ``
}


type releaseHeldQuantitiesResponse struct {
	ReleasedQuantities releaseHeldQuantitiesResponse_sub1 ``
}

type releaseHeldQuantitiesResponse_sub1 struct {
	SKUAsKey int64 ``
}
