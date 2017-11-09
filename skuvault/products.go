package skuvault
	// CreateBrands creates http request for this SKU vault endpoint
	// Moderate Throttle
	// NONE
	func (lc *PLoginCredentials) CreateBrands(pld *products.CreateBrands) *products.CreateBrandsResponse {
		credPld := &postCreateBrands {
			CreateBrands:       pld,
			PLoginCredentials: lc,
		}

		response := &products.CreateBrandsResponse{}
		productsCall(credPld, response, createBrands)
		return response
	}
	// CreateKit creates http request for this SKU vault endpoint
	// Moderate Throttle
	// Using this call, users may create a kit inside of SkuVault.
	func (lc *PLoginCredentials) CreateKit(pld *products.CreateKit) *products.CreateKitResponse {
		credPld := &postCreateKit {
			CreateKit:       pld,
			PLoginCredentials: lc,
		}

		response := &products.CreateKitResponse{}
		productsCall(credPld, response, createKit)
		return response
	}
	// CreateProduct creates http request for this SKU vault endpoint
	// Moderate Throttle
	// Create products in SkuVault one at a time. Bulk version available: 
	func (lc *PLoginCredentials) CreateProduct(pld *products.CreateProduct) *products.CreateProductResponse {
		credPld := &postCreateProduct {
			CreateProduct:       pld,
			PLoginCredentials: lc,
		}

		response := &products.CreateProductResponse{}
		productsCall(credPld, response, createProduct)
		return response
	}
	// CreateProducts creates http request for this SKU vault endpoint
	// Heavy Throttle
	// Create 100 products per request.
	func (lc *PLoginCredentials) CreateProducts(pld *products.CreateProducts) *products.CreateProductsResponse {
		credPld := &postCreateProducts {
			CreateProducts:       pld,
			PLoginCredentials: lc,
		}

		response := &products.CreateProductsResponse{}
		productsCall(credPld, response, createProducts)
		return response
	}
	// CreateSuppliers creates http request for this SKU vault endpoint
	// Moderate Throttle
	// Returns the list of current Suppliers in a SkuVault account.
	func (lc *PLoginCredentials) CreateSuppliers(pld *products.CreateSuppliers) *products.CreateSuppliersResponse {
		credPld := &postCreateSuppliers {
			CreateSuppliers:       pld,
			PLoginCredentials: lc,
		}

		response := &products.CreateSuppliersResponse{}
		productsCall(credPld, response, createSuppliers)
		return response
	}
	// GetBrands creates http request for this SKU vault endpoint
	// Heavy Throttle
	// Returns a list of Brands created in SkuVault.
	func (lc *PLoginCredentials) GetBrands(pld *products.GetBrands) *products.GetBrandsResponse {
		credPld := &postGetBrands {
			GetBrands:       pld,
			PLoginCredentials: lc,
		}

		response := &products.GetBrandsResponse{}
		productsCall(credPld, response, getBrands)
		return response
	}
	// GetClassifications creates http request for this SKU vault endpoint
	// Heavy Throttle
	// Returns classifications and, if they exist, their named attributes.
	func (lc *PLoginCredentials) GetClassifications(pld *products.GetClassifications) *products.GetClassificationsResponse {
		credPld := &postGetClassifications {
			GetClassifications:       pld,
			PLoginCredentials: lc,
		}

		response := &products.GetClassificationsResponse{}
		productsCall(credPld, response, getClassifications)
		return response
	}
	// GetKits creates http request for this SKU vault endpoint
	// Heavy Throttle
	// Returns kit details.
	func (lc *PLoginCredentials) GetKits(pld *products.GetKits) *products.GetKitsResponse {
		credPld := &postGetKits {
			GetKits:       pld,
			PLoginCredentials: lc,
		}

		response := &products.GetKitsResponse{}
		productsCall(credPld, response, getKits)
		return response
	}
	// GetProducts creates http request for this SKU vault endpoint
	// Heavy Throttle
	// This call returns product(not kit) details. The date parameters include updating details as well as product creation. Details do not include quantity.
	func (lc *PLoginCredentials) GetProducts(pld *products.GetProducts) *products.GetProductsResponse {
		credPld := &postGetProducts {
			GetProducts:       pld,
			PLoginCredentials: lc,
		}

		response := &products.GetProductsResponse{}
		productsCall(credPld, response, getProducts)
		return response
	}
	// GetSuppliers creates http request for this SKU vault endpoint
	// Heavy Throttle
	// Returns the list of current Suppliers in a SkuVault account.
	func (lc *PLoginCredentials) GetSuppliers(pld *products.GetSuppliers) *products.GetSuppliersResponse {
		credPld := &postGetSuppliers {
			GetSuppliers:       pld,
			PLoginCredentials: lc,
		}

		response := &products.GetSuppliersResponse{}
		productsCall(credPld, response, getSuppliers)
		return response
	}
	// UpdateHandlingTime creates http request for this SKU vault endpoint
	// Severe Throttle
	// Update the handling time for each product per Amazon channel account, 500 at a time.
	func (lc *PLoginCredentials) UpdateHandlingTime(pld *products.UpdateHandlingTime) *products.UpdateHandlingTimeResponse {
		credPld := &postUpdateHandlingTime {
			UpdateHandlingTime:       pld,
			PLoginCredentials: lc,
		}

		response := &products.UpdateHandlingTimeResponse{}
		productsCall(credPld, response, updateHandlingTime)
		return response
	}
	// UpdateProduct creates http request for this SKU vault endpoint
	// Moderate Throttle
	// Update your product details. Bulk version available: 
	func (lc *PLoginCredentials) UpdateProduct(pld *products.UpdateProduct) *products.UpdateProductResponse {
		credPld := &postUpdateProduct {
			UpdateProduct:       pld,
			PLoginCredentials: lc,
		}

		response := &products.UpdateProductResponse{}
		productsCall(credPld, response, updateProduct)
		return response
	}
	// UpdateProducts creates http request for this SKU vault endpoint
	// Heavy Throttle
	// Update products in SkuVault, 100 at a time.
	func (lc *PLoginCredentials) UpdateProducts(pld *products.UpdateProducts) *products.UpdateProductsResponse {
		credPld := &postUpdateProducts {
			UpdateProducts:       pld,
			PLoginCredentials: lc,
		}

		response := &products.UpdateProductsResponse{}
		productsCall(credPld, response, updateProducts)
		return response
	}