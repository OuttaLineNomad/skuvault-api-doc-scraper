package inventory\n

type addItemBulk struct {
	Errors  []addItemBulk_sub1 ``
	Results []string           ``
	Status  string             ``
}

type addItemBulk_sub1 struct {
	ErrorMessages []string ``
	Sku           string   ``
}
