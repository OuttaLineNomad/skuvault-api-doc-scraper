package sales
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
		// GetSoldItems creates http request for this SKU vault endpoint
		// Heavy Throttle
		// Returns a list of sold items filtered by date. 10,000 sales are returned per page.
		func (lc *SLoginCredentials) GetSoldItems(pld *sales.GetSoldItems) *sales.GetSoldItemsResponse {
			credPld := &postGetSoldItems {
				GetSoldItems:       pld,
				SLoginCredentials: lc,
			}

			response := &sales.GetSoldItemsResponse{}
			salesCall(credPld, response, getSoldItems)
			return response
		}
		// ReleaseHeldQuantities creates http request for this SKU vault endpoint
		// Moderate Throttle
		// Release holds before their expiration date expires.
		func (lc *SLoginCredentials) ReleaseHeldQuantities(pld *sales.ReleaseHeldQuantities) *sales.ReleaseHeldQuantitiesResponse {
			credPld := &postReleaseHeldQuantities {
				ReleaseHeldQuantities:       pld,
				SLoginCredentials: lc,
			}

			response := &sales.ReleaseHeldQuantitiesResponse{}
			salesCall(credPld, response, releaseHeldQuantities)
			return response
		}
		// SyncOnlineSale creates http request for this SKU vault endpoint
		// Moderate Throttle
		// Sync an online sale to SkuVault. If the sale does not exists, it&#39;s created. If it does exist, it&#39;s updated.  ShippingStatus is required to create sale, but not update. ItemSkus is always required. Bulk version available: 
		func (lc *SLoginCredentials) SyncOnlineSale(pld *sales.SyncOnlineSale) *sales.SyncOnlineSaleResponse {
			credPld := &postSyncOnlineSale {
				SyncOnlineSale:       pld,
				SLoginCredentials: lc,
			}

			response := &sales.SyncOnlineSaleResponse{}
			salesCall(credPld, response, syncOnlineSale)
			return response
		}
		// SyncOnlineSales creates http request for this SKU vault endpoint
		// Severe Throttle
		//  Can make this call 2x per minute, 100 sales max
		func (lc *SLoginCredentials) SyncOnlineSales(pld *sales.SyncOnlineSales) *sales.SyncOnlineSalesResponse {
			credPld := &postSyncOnlineSales {
				SyncOnlineSales:       pld,
				SLoginCredentials: lc,
			}

			response := &sales.SyncOnlineSalesResponse{}
			salesCall(credPld, response, syncOnlineSales)
			return response
		}
		// SyncShippedSaleAndRemoveItems creates http request for this SKU vault endpoint
		// Moderate Throttle
		// This method syncs a shipped sale and auto-removes quantity. It will auto-remove quantities based on the first match alphabetically in your warehouses &amp; locations. If a warehouse is provided, then it will use that warehouse.
		func (lc *SLoginCredentials) SyncShippedSaleAndRemoveItems(pld *sales.SyncShippedSaleAndRemoveItems) *sales.SyncShippedSaleAndRemoveItemsResponse {
			credPld := &postSyncShippedSaleAndRemoveItems {
				SyncShippedSaleAndRemoveItems:       pld,
				SLoginCredentials: lc,
			}

			response := &sales.SyncShippedSaleAndRemoveItemsResponse{}
			salesCall(credPld, response, syncShippedSaleAndRemoveItems)
			return response
		}
		// UpdateOnlineSaleStatus creates http request for this SKU vault endpoint
		// Light Throttle
		// Update the status of a sale.
		func (lc *SLoginCredentials) UpdateOnlineSaleStatus(pld *sales.UpdateOnlineSaleStatus) *sales.UpdateOnlineSaleStatusResponse {
			credPld := &postUpdateOnlineSaleStatus {
				UpdateOnlineSaleStatus:       pld,
				SLoginCredentials: lc,
			}

			response := &sales.UpdateOnlineSaleStatusResponse{}
			salesCall(credPld, response, updateOnlineSaleStatus)
			return response
		}
		// UpdateShipments creates http request for this SKU vault endpoint
		// Severe Throttle
		// Using this call, users can update shipments to a sale.
		func (lc *SLoginCredentials) UpdateShipments(pld *sales.UpdateShipments) *sales.UpdateShipmentsResponse {
			credPld := &postUpdateShipments {
				UpdateShipments:       pld,
				SLoginCredentials: lc,
			}

			response := &sales.UpdateShipmentsResponse{}
			salesCall(credPld, response, updateShipments)
			return response
		}