package purchaseorders
		// CreatePO creates http request for this SKU vault endpoint
		// Moderate Throttle
		// This call let&#39;s you create a PO using our API.
		func (lc *POLoginCredentials) CreatePO(pld *purchaseorders.CreatePO) *purchaseorders.CreatePOResponse {
			credPld := &postCreatePO {
				CreatePO:       pld,
				POLoginCredentials: lc,
			}

			response := &purchaseorders.CreatePOResponse{}
			purchaseordersCall(credPld, response, createPO)
			return response
		}
		// GetIncomingItems creates http request for this SKU vault endpoint
		// Get incoming items for incomplete purchase orders Throttle
		// NONE
		func (lc *POLoginCredentials) GetIncomingItems(pld *purchaseorders.GetIncomingItems) *purchaseorders.GetIncomingItemsResponse {
			credPld := &postGetIncomingItems {
				GetIncomingItems:       pld,
				POLoginCredentials: lc,
			}

			response := &purchaseorders.GetIncomingItemsResponse{}
			purchaseordersCall(credPld, response, getIncomingItems)
			return response
		}
		// GetPOs creates http request for this SKU vault endpoint
		// Heavy Throttle
		// Returns a list of purchase orders.
		func (lc *POLoginCredentials) GetPOs(pld *purchaseorders.GetPOs) *purchaseorders.GetPOsResponse {
			credPld := &postGetPOs {
				GetPOs:       pld,
				POLoginCredentials: lc,
			}

			response := &purchaseorders.GetPOsResponse{}
			purchaseordersCall(credPld, response, getPOs)
			return response
		}
		// GetReceivesHistory creates http request for this SKU vault endpoint
		// Heavy Throttle
		// Returns a list of purchase order receives and receipts.
		func (lc *POLoginCredentials) GetReceivesHistory(pld *purchaseorders.GetReceivesHistory) *purchaseorders.GetReceivesHistoryResponse {
			credPld := &postGetReceivesHistory {
				GetReceivesHistory:       pld,
				POLoginCredentials: lc,
			}

			response := &purchaseorders.GetReceivesHistoryResponse{}
			purchaseordersCall(credPld, response, getReceivesHistory)
			return response
		}
		// ReceivePOItems creates http request for this SKU vault endpoint
		// Moderate Throttle
		// NONE
		func (lc *POLoginCredentials) ReceivePOItems(pld *purchaseorders.ReceivePOItems) *purchaseorders.ReceivePOItemsResponse {
			credPld := &postReceivePOItems {
				ReceivePOItems:       pld,
				POLoginCredentials: lc,
			}

			response := &purchaseorders.ReceivePOItemsResponse{}
			purchaseordersCall(credPld, response, receivePOItems)
			return response
		}