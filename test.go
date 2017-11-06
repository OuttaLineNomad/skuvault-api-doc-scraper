package main

// AddItem creates http request for this SKU vault endpoint
// Moderate Throttle
// Add quantity to a warehouse location.
func (lc *PLoginCredentials) AddItem(pld *inventory.AddItem) *inventory.AddItemResponse {
	credPld := &postAddItem{
		AddItem:           pld,
		iLoginCredentials: lc,
	}

	response := &inventory.AddItemResponse{}
	inventoryCall(credPld, response, addItem)
	return response
}
