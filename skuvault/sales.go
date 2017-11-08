package skuvault
	// AddShipments creates http request for this SKU vault endpoint
	// Heavy Throttle
	// Using this call, users can add shipments to a sale.
	func (lc *SLoginCredentials) AddShipments(pld *sales.AddShipments) *sales.AddShipmentsResponse {
		credPld := &postAddShipments {
			AddShipments:       pld,
			SLoginCredentials: lc,
		}

		response := &sales.AddShipmentsResponse{}
		salesCall(credPld, response, addShipments)
		return response
	}
	// CreateHolds creates http request for this SKU vault endpoint
	// Moderate Throttle
	// NONE
	func (lc *SLoginCredentials) CreateHolds(pld *sales.CreateHolds) *sales.CreateHoldsResponse {
		credPld := &postCreateHolds {
			CreateHolds:       pld,
			SLoginCredentials: lc,
		}

		response := &sales.CreateHoldsResponse{}
		salesCall(credPld, response, createHolds)
		return response
	}
	// GetOnlineSaleStatus creates http request for this SKU vault endpoint
	// Heavy Throttle
	// Returns a list of sales and their statuses.
	func (lc *SLoginCredentials) GetOnlineSaleStatus(pld *sales.GetOnlineSaleStatus) *sales.GetOnlineSaleStatusResponse {
		credPld := &postGetOnlineSaleStatus {
			GetOnlineSaleStatus:       pld,
			SLoginCredentials: lc,
		}

		response := &sales.GetOnlineSaleStatusResponse{}
		salesCall(credPld, response, getOnlineSaleStatus)
		return response
	}
	// GetSales creates http request for this SKU vault endpoint
	// Heavy Throttle
	// Use this call to retrieve a list of sales from SkuVault. 10,000 sales are returned per page.
	func (lc *SLoginCredentials) GetSales(pld *sales.GetSales) *sales.GetSalesResponse {
		credPld := &postGetSales {
			GetSales:       pld,
			SLoginCredentials: lc,
		}

		response := &sales.GetSalesResponse{}
		salesCall(credPld, response, getSales)
		return response
	}
	// GetSalesByDate creates http request for this SKU vault endpoint
	// Heavy Throttle
	// Returns sales based on a date range. 10,000 sales are returned per page.
	func (lc *SLoginCredentials) GetSalesByDate(pld *sales.GetSalesByDate) *sales.GetSalesByDateResponse {
		credPld := &postGetSalesByDate {
			GetSalesByDate:       pld,
			SLoginCredentials: lc,
		}

		response := &sales.GetSalesByDateResponse{}
		salesCall(credPld, response, getSalesByDate)
		return response
	}