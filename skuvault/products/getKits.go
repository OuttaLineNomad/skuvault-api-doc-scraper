package products\n

type getKits struct {
	Errors []string       ``
	Kits   []getKits_sub3 ``
}

type getKits_sub3 struct {
	AvailableQuantity                        int64          ``
	AvailableQuantityLastModifiedDateTimeUtc string         ``
	Code                                     string         ``
	Cost                                     int64          ``
	Description                              string         ``
	KitLines                                 []getKits_sub2 ``
	LastModifiedDateTimeUtc                  string         ``
	Sku                                      string         ``
}

type getKits_sub1 struct {
	Code        string ``
	Description string ``
	Sku         string ``
}

type getKits_sub2 struct {
	Combine  int64          ``
	Items    []getKits_sub1 ``
	LineName string         ``
	Quantity int64          ``
}
