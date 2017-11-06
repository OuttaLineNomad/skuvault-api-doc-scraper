
		// AddShipments creates http request for this SKU vault endpoint
		// Heavy Throttle
		// Using this call, users can add shipments to a sale.
		func (lc *PLoginCredentials) AddShipments(pld *sales.AddShipments) *sales.AddShipmentsResponse {
			credPld := &postAddShipments {
				AddShipments:       pld,
				sLoginCredentials: lc,
			}

			response := &sales.AddShipmentsResponse{}
			salesCall(credPld, response, addShipments)
			return response
		}