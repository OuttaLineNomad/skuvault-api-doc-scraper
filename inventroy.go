package inventory

// AddItem creates http request for this SKU vault endpoint
// Moderate Throttle
// Add quantity to a warehouse location.
func (lc *ILoginCredentials) AddItem(pld *inventory.AddItem) *inventory.AddItemResponse {
	credPld := &postAddItem{
		AddItem:           pld,
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
	credPld := &postAddItemBulk{
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
	credPld := &postGetAvailableQuantities{
		GetAvailableQuantities: pld,
		ILoginCredentials:      lc,
	}

	response := &inventory.GetAvailableQuantitiesResponse{}
	inventoryCall(credPld, response, getAvailableQuantities)
	return response
}

// GetExternalWarehouseQuantities creates http request for this SKU vault endpoint
// Heavy Throttle
// This call gets the quantities in a designated External Warehouse. Please note these are different than ordinary warehouses.
func (lc *ILoginCredentials) GetExternalWarehouseQuantities(pld *inventory.GetExternalWarehouseQuantities) *inventory.GetExternalWarehouseQuantitiesResponse {
	credPld := &postGetExternalWarehouseQuantities{
		GetExternalWarehouseQuantities: pld,
		ILoginCredentials:              lc,
	}

	response := &inventory.GetExternalWarehouseQuantitiesResponse{}
	inventoryCall(credPld, response, getExternalWarehouseQuantities)
	return response
}

// GetExternalWarehouses creates http request for this SKU vault endpoint
// Moderate Throttle
// Returns your external warehouses. No page parameters.
func (lc *ILoginCredentials) GetExternalWarehouses(pld *inventory.GetExternalWarehouses) *inventory.GetExternalWarehousesResponse {
	credPld := &postGetExternalWarehouses{
		GetExternalWarehouses: pld,
		ILoginCredentials:     lc,
	}

	response := &inventory.GetExternalWarehousesResponse{}
	inventoryCall(credPld, response, getExternalWarehouses)
	return response
}

// GetInventoryByLocation creates http request for this SKU vault endpoint
// Heavy Throttle
// Returns location and warehouse per product.
func (lc *ILoginCredentials) GetInventoryByLocation(pld *inventory.GetInventoryByLocation) *inventory.GetInventoryByLocationResponse {
	credPld := &postGetInventoryByLocation{
		GetInventoryByLocation: pld,
		ILoginCredentials:      lc,
	}

	response := &inventory.GetInventoryByLocationResponse{}
	inventoryCall(credPld, response, getInventoryByLocation)
	return response
}

// GetItemQuantities creates http request for this SKU vault endpoint
// Heavy Throttle
// Returns product quantities.
func (lc *ILoginCredentials) GetItemQuantities(pld *inventory.GetItemQuantities) *inventory.GetItemQuantitiesResponse {
	credPld := &postGetItemQuantities{
		GetItemQuantities: pld,
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
	credPld := &postGetKitQuantities{
		GetKitQuantities:  pld,
		ILoginCredentials: lc,
	}

	response := &inventory.GetKitQuantitiesResponse{}
	inventoryCall(credPld, response, getKitQuantities)
	return response
}

// GetTransactions creates http request for this SKU vault endpoint
// Heavy Throttle
// Look at your transaction history.
func (lc *ILoginCredentials) GetTransactions(pld *inventory.GetTransactions) *inventory.GetTransactionsResponse {
	credPld := &postGetTransactions{
		GetTransactions:   pld,
		ILoginCredentials: lc,
	}

	response := &inventory.GetTransactionsResponse{}
	inventoryCall(credPld, response, getTransactions)
	return response
}

// GetWarehouseItemQuantities creates http request for this SKU vault endpoint
// Heavy Throttle
// This call returns SKUs and quantities from a specified warehouse. 10,000 SKUs returned per page.
func (lc *ILoginCredentials) GetWarehouseItemQuantities(pld *inventory.GetWarehouseItemQuantities) *inventory.GetWarehouseItemQuantitiesResponse {
	credPld := &postGetWarehouseItemQuantities{
		GetWarehouseItemQuantities: pld,
		ILoginCredentials:          lc,
	}

	response := &inventory.GetWarehouseItemQuantitiesResponse{}
	inventoryCall(credPld, response, getWarehouseItemQuantities)
	return response
}

// GetWarehouseItemQuantity creates http request for this SKU vault endpoint
// Heavy Throttle
// Returns the quantity for a specified SKU.
func (lc *ILoginCredentials) GetWarehouseItemQuantity(pld *inventory.GetWarehouseItemQuantity) *inventory.GetWarehouseItemQuantityResponse {
	credPld := &postGetWarehouseItemQuantity{
		GetWarehouseItemQuantity: pld,
		ILoginCredentials:        lc,
	}

	response := &inventory.GetWarehouseItemQuantityResponse{}
	inventoryCall(credPld, response, getWarehouseItemQuantity)
	return response
}

// GetWarehouses creates http request for this SKU vault endpoint
// Severe Throttle
// Returns all your regular warehouses.
func (lc *ILoginCredentials) GetWarehouses(pld *inventory.GetWarehouses) *inventory.GetWarehousesResponse {
	credPld := &postGetWarehouses{
		GetWarehouses:     pld,
		ILoginCredentials: lc,
	}

	response := &inventory.GetWarehousesResponse{}
	inventoryCall(credPld, response, getWarehouses)
	return response
}

// PickItem creates http request for this SKU vault endpoint
// Moderate Throttle
// Perform a pick transaction through the API.
func (lc *ILoginCredentials) PickItem(pld *inventory.PickItem) *inventory.PickItemResponse {
	credPld := &postPickItem{
		PickItem:          pld,
		ILoginCredentials: lc,
	}

	response := &inventory.PickItemResponse{}
	inventoryCall(credPld, response, pickItem)
	return response
}

// RemoveItem creates http request for this SKU vault endpoint
// Moderate Throttle
// Remove quantity from a warehouse location.
func (lc *ILoginCredentials) RemoveItem(pld *inventory.RemoveItem) *inventory.RemoveItemResponse {
	credPld := &postRemoveItem{
		RemoveItem:        pld,
		ILoginCredentials: lc,
	}

	response := &inventory.RemoveItemResponse{}
	inventoryCall(credPld, response, removeItem)
	return response
}

// RemoveItemBulk creates http request for this SKU vault endpoint
// Moderate Throttle
// Remove quantity from warehouse locations, 100 at a time.
func (lc *ILoginCredentials) RemoveItemBulk(pld *inventory.RemoveItemBulk) *inventory.RemoveItemBulkResponse {
	credPld := &postRemoveItemBulk{
		RemoveItemBulk:    pld,
		ILoginCredentials: lc,
	}

	response := &inventory.RemoveItemBulkResponse{}
	inventoryCall(credPld, response, removeItemBulk)
	return response
}

// SetItemQuantities creates http request for this SKU vault endpoint
// Moderate Throttle
// Sets quantity for multiple products in one request.
func (lc *ILoginCredentials) SetItemQuantities(pld *inventory.SetItemQuantities) *inventory.SetItemQuantitiesResponse {
	credPld := &postSetItemQuantities{
		SetItemQuantities: pld,
		ILoginCredentials: lc,
	}

	response := &inventory.SetItemQuantitiesResponse{}
	inventoryCall(credPld, response, setItemQuantities)
	return response
}

// SetItemQuantity creates http request for this SKU vault endpoint
// Moderate Throttle
// This lets you explicitly set quantity for an item in a warehouse&#39;s location.
func (lc *ILoginCredentials) SetItemQuantity(pld *inventory.SetItemQuantity) *inventory.SetItemQuantityResponse {
	credPld := &postSetItemQuantity{
		SetItemQuantity:   pld,
		ILoginCredentials: lc,
	}

	response := &inventory.SetItemQuantityResponse{}
	inventoryCall(credPld, response, setItemQuantity)
	return response
}

// UpdateExternalWarehouseQuantities creates http request for this SKU vault endpoint
// Severe Throttle
// Set the quantity of SKUs in a specified external warehouse. The limit is 100,000 SKUs per call.
func (lc *ILoginCredentials) UpdateExternalWarehouseQuantities(pld *inventory.UpdateExternalWarehouseQuantities) *inventory.UpdateExternalWarehouseQuantitiesResponse {
	credPld := &postUpdateExternalWarehouseQuantities{
		UpdateExternalWarehouseQuantities: pld,
		ILoginCredentials:                 lc,
	}

	response := &inventory.UpdateExternalWarehouseQuantitiesResponse{}
	inventoryCall(credPld, response, updateExternalWarehouseQuantities)
	return response
}
