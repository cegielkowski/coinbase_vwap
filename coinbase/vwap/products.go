package vwap

import "coinbase_vwap/domain"

// NewProducts will create new a domain.Products object.
func NewProducts(productIds []string, vwapCalculationEngineGenerator []domain.VwapCalculationEngine) domain.Products {
	products := make(map[string]domain.VwapCalculationEngine)

	for key, value := range productIds {
		products[value] = vwapCalculationEngineGenerator[key]
	}

	return products
}
