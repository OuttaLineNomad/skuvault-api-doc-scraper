package skuvault
	// GetIntegrations creates http request for this SKU vault endpoint
	// Severe Throttle
	// Returns a list of your enabled channel accounts. No page parameters.
	func (lc *LoginCredentials) GetIntegrations(pld *integration.GetIntegrations) *integration.GetIntegrationsResponse {
		credPld := &postGetIntegrations {
			GetIntegrations:       pld,
			LoginCredentials: lc,
		}

		response := &integration.GetIntegrationsResponse{}
		integrationCall(credPld, response, getIntegrations)
		return response
	}
	// GetTokens creates http request for this SKU vault endpoint
	// Very Light Throttle
	// Use this call to retrieve your API tokens from SkuVault using your login email and password.
	func (lc *LoginCredentials) GetTokens(pld *.GetTokens) *.GetTokensResponse {
		credPld := &postGetTokens {
			GetTokens:       pld,
			LoginCredentials: lc,
		}

		response := &.GetTokensResponse{}
		Call(credPld, response, getTokens)
		return response
	}