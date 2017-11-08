package skuvault
	// AddItem creates http request for this SKU vault endpoint
	// Moderate Throttle
	// Add quantity to a warehouse location. 
	func (lc *ILoginCredentials) AddItem(pld *inventory.AddItem) *inventory.AddItemResponse {
		credPld := &postAddItem {
			AddItem:       pld,
			ILoginCredentials: lc,
		}

		response := &inventory.AddItemResponse{}
		inventoryCall(credPld, response, addItem)
		return response
	}
	// AddItemBulk creates http request for this SKU vault endpoint
	// Moderate Throttle
	// Add quantity to warehouse locations, 100 at a time.
	func (lc *ILoginCredentials) AddItemBulk(pld *inventory.AddItemBulk) *inventory.AddItemBulkResponse {
		credPld := &postAddItemBulk {
			AddItemBulk:       pld,
			ILoginCredentials: lc,
		}

		response := &inventory.AddItemBulkResponse{}
		inventoryCall(credPld, response, addItemBulk)
		return response
	}
	// GetAvailableQuantities creates http request for this SKU vault endpoint
	// Heavy Throttle
	// Retrieve a list of SKUs and their total available quantities across all warehouses. Available Quantity is the quantity that is actually available to sell across all your sales channels.
	func (lc *ILoginCredentials) GetAvailableQuantities(pld *inventory.GetAvailableQuantities) *inventory.GetAvailableQuantitiesResponse {
		credPld := &postGetAvailableQuantities {
			GetAvailableQuantities:       pld,
			ILoginCredentials: lc,
		}

		response := &inventory.GetAvailableQuantitiesResponse{}
		inventoryCall(credPld, response, getAvailableQuantities)
		return response
	}
	// GetExternalWarehouseQuantities creates http request for this SKU vault endpoint
	// Heavy Throttle
	// This call gets the quantities in a designated External Warehouse. Please note these are different than ordinary warehouses.
	func (lc *ILoginCredentials) GetExternalWarehouseQuantities(pld *inventory.GetExternalWarehouseQuantities) *inventory.GetExternalWarehouseQuantitiesResponse {
		credPld := &postGetExternalWarehouseQuantities {
			GetExternalWarehouseQuantities:       pld,
			ILoginCredentials: lc,
		}

		response := &inventory.GetExternalWarehouseQuantitiesResponse{}
		inventoryCall(credPld, response, getExternalWarehouseQuantities)
		return response
	}
	// GetExternalWarehouses creates http request for this SKU vault endpoint
	// Moderate Throttle
	// Returns your external warehouses. No page parameters.
	func (lc *ILoginCredentials) GetExternalWarehouses(pld *inventory.GetExternalWarehouses) *inventory.GetExternalWarehousesResponse {
		credPld := &postGetExternalWarehouses {
			GetExternalWarehouses:       pld,
			ILoginCredentials: lc,
		}

		response := &inventory.GetExternalWarehousesResponse{}
		inventoryCall(credPld, response, getExternalWarehouses)
		return response
	}
	// GetInventoryByLocation creates http request for this SKU vault endpoint
	// Heavy Throttle
	// Returns location and warehouse per product.
	func (lc *ILoginCredentials) GetInventoryByLocation(pld *inventory.GetInventoryByLocation) *inventory.GetInventoryByLocationResponse {
		credPld := &postGetInventoryByLocation {
			GetInventoryByLocation:       pld,
			ILoginCredentials: lc,
		}

		response := &inventory.GetInventoryByLocationResponse{}
		inventoryCall(credPld, response, getInventoryByLocation)
		return response
	}
	// GetItemQuantities creates http request for this SKU vault endpoint
	// Heavy Throttle
	// Returns product quantities.
	func (lc *ILoginCredentials) GetItemQuantities(pld *inventory.GetItemQuantities) *inventory.GetItemQuantitiesResponse {
		credPld := &postGetItemQuantities {
			GetItemQuantities:       pld,
			ILoginCredentials: lc,
		}

		response := &inventory.GetItemQuantitiesResponse{}
		inventoryCall(credPld, response, getItemQuantities)
		return response
	}
	// GetKitQuantities creates http request for this SKU vault endpoint
	// Heavy Throttle
	// Returns kit quantities.
	func (lc *ILoginCredentials) GetKitQuantities(pld *inventory.GetKitQuantities) *inventory.GetKitQuantitiesResponse {
		credPld := &postGetKitQuantities {
			GetKitQuantities:       pld,
			ILoginCredentials: lc,
		}

		response := &inventory.GetKitQuantitiesResponse{}
		inventoryCall(credPld, response, getKitQuantities)
		return response
	}