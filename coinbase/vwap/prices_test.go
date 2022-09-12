package vwap

import (
	"math"
	"testing"
)

func Test_prices_addNew(t *testing.T) {
	type fields struct {
		prices []float64
		max    float64
		min    float64
	}
	type args struct {
		price float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "add first time",
			fields: fields{
				prices: []float64{1, 2, 3},
				max:    0,
				min:    0,
			},
			args: args{
				price: 4,
			},
		},
		{
			name: "add second time",
			fields: fields{
				prices: []float64{},
				max:    0,
				min:    0,
			},
			args: args{
				price: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &prices{
				prices: tt.fields.prices,
				max:    tt.fields.max,
				min:    tt.fields.min,
			}
			p.addNew(tt.args.price)
			lastPrice := p.prices[len(p.prices)-1]
			if lastPrice != tt.args.price {
				t.Errorf("addNew() lastPrice = %v, want %v", lastPrice, tt.args.price)
			}
		})
	}
}

func Test_prices_findMax(t *testing.T) {
	type fields struct {
		prices []float64
		max    float64
		min    float64
	}
	tests := []struct {
		name     string
		fields   fields
		expected float64
	}{
		{
			name: "first test",
			fields: fields{
				prices: []float64{1, 3, 6, 1, 2, 3},
				max:    0,
				min:    0,
			},
			expected: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &prices{
				prices: tt.fields.prices,
				max:    tt.fields.max,
				min:    tt.fields.min,
			}
			p.findMax()

			if p.max != tt.expected {
				t.Errorf("findMax() p.max = %v, expected %v", p.max, tt.expected)
			}
		})
	}
}

func Test_prices_findMin(t *testing.T) {
	type fields struct {
		prices []float64
		max    float64
		min    float64
	}
	tests := []struct {
		name     string
		fields   fields
		expected float64
	}{
		{
			name: "first test",
			fields: fields{
				prices: []float64{1, 3, 6, 1, 2, 3},
				max:    0,
				min:    math.MaxFloat64,
			},
			expected: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &prices{
				prices: tt.fields.prices,
				max:    tt.fields.max,
				min:    tt.fields.min,
			}
			p.findMin()
			if p.min != tt.expected {
				t.Errorf("findMin() p.min = %v, expected %v", p.min, tt.expected)
			}
		})
	}
}
