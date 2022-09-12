package domain

type VwapCalculationEngine interface {
	Calculate(price, volume float64) float64
}
