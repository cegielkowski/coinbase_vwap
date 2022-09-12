package vwap

import (
	"coinbase_vwap/domain"
	"math"
)

type vwapCalculationEngine struct {
	size                     int16
	prices                   prices
	volumes                  []float64
	cumulatedVolume          float64
	totalVolumeWeightedPrice float64
	valueLimitSize           int16
}

// Calculate and return the vwap in real time.
func (v *vwapCalculationEngine) Calculate(price, volume float64) float64 {
	isFull := v.size == v.valueLimitSize
	removedValue := v.prices.receive(price, isFull)
	v.incrementSize(isFull)

	if isFull {
		removedVolume := v.recalculateVolume()
		v.subTotalVolumeWeightedPrice(removedVolume, removedValue)
	}

	v.addVolume(volume)
	v.addTotalVolumeWeightedPrice(volume, price)

	return v.calculateVwap(v.totalVolumeWeightedPrice, v.cumulatedVolume)
}

func (v *vwapCalculationEngine) calculateVwap(totalVolumeWeightedPrice float64, cumulatedVolume float64) float64 {
	if cumulatedVolume == 0 {
		return 0
	}

	return totalVolumeWeightedPrice / cumulatedVolume
}

func (v *vwapCalculationEngine) addTotalVolumeWeightedPrice(volume float64, price float64) {
	v.totalVolumeWeightedPrice += volume * price
}

func (v *vwapCalculationEngine) subTotalVolumeWeightedPrice(volume float64, price float64) {
	v.totalVolumeWeightedPrice -= volume * price
}

func (v *vwapCalculationEngine) recalculateVolume() float64 {
	toBeRemoved := v.volumes[0]
	v.cumulatedVolume -= toBeRemoved
	v.volumes = v.volumes[1:]
	return toBeRemoved
}

func (v *vwapCalculationEngine) addVolume(volume float64) {
	v.cumulatedVolume += volume
	v.volumes = append(v.volumes, volume)
}

func (v *vwapCalculationEngine) incrementSize(isFull bool) {
	if !isFull {
		v.size = v.size + 1
	}
}

func newVwapCalculationEngine(valueLimitSize int16) domain.VwapCalculationEngine {
	return &vwapCalculationEngine{
		prices:         prices{prices: make([]float64, 0, valueLimitSize), min: math.MaxFloat64},
		volumes:        make([]float64, 0, valueLimitSize),
		valueLimitSize: valueLimitSize,
	}
}

// NewVwapCalculationEngines will create new an newVwapCalculationEngine object representation of []domain.VwapCalculationEngine interface.
func NewVwapCalculationEngines(valueLimitSize int16, size int) []domain.VwapCalculationEngine {
	var engines []domain.VwapCalculationEngine
	for i := 0; i < size; i++ {
		engines = append(engines, newVwapCalculationEngine(valueLimitSize))
	}

	return engines
}
