package vwap

import "math"

type prices struct {
	prices []float64
	max    float64
	min    float64
}

func (p *prices) checkIfIsMaxOrMin(price float64) (isMax, isMin bool) {
	isMax = false
	isMin = false
	if price == p.max {
		isMax = true
	}

	if price == p.min {
		isMin = true
	}

	return
}

func (p *prices) receive(price float64, full bool) float64 {
	isMax := false
	isMin := false
	toBeRemoved := float64(0)
	if full {
		toBeRemoved = p.prices[0]
		isMax, isMin = p.checkIfIsMaxOrMin(toBeRemoved)
		p.removeLast()
	}

	p.addNew(price)

	if isMax || p.max < price {
		p.findMax()
	}
	if isMin || p.min > price {
		p.findMin()
	}

	return toBeRemoved
}

func (p *prices) removeLast() {
	p.prices = p.prices[1:]
}

func (p *prices) addNew(price float64) {
	p.prices = append(p.prices, price)
}

func (p *prices) findMax() {
	p.max = float64(0)
	for _, value := range p.prices {
		if value > p.max {
			p.max = value
		}
	}
}

func (p *prices) findMin() {
	p.min = math.MaxFloat64
	for _, value := range p.prices {
		if value < p.min {
			p.min = value
		}
	}
}
