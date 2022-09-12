package vwap

import (
	"math"
	"testing"
)

func Test_vwapCalculationEngine_Calculate(t *testing.T) {
	type fields struct {
		size                     int16
		prices                   prices
		volumes                  []float64
		cumulatedVolume          float64
		totalVolumeWeightedPrice float64
		valueLimitSize           int16
	}
	type args struct {
		price  float64
		volume float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "basic test",
			fields: fields{
				size: 0,
				prices: prices{
					prices: make([]float64, 0),
					max:    0,
					min:    math.MaxFloat64,
				},
				volumes:                  make([]float64, 0),
				cumulatedVolume:          0,
				totalVolumeWeightedPrice: 0,
				valueLimitSize:           200,
			},
			args: args{
				price:  1,
				volume: 1,
			},
			want: 1,
		},
		{
			name: "basic test",
			fields: fields{
				size: 3,
				prices: prices{
					prices: []float64{2.1, 1.3, 1.823},
					max:    2.1,
					min:    1.3,
				},
				volumes:                  []float64{3.11, 3.2134, 2.341223},
				cumulatedVolume:          8.664623,
				totalVolumeWeightedPrice: 14.976469529,
				valueLimitSize:           200,
			},
			args: args{
				price:  1.32123,
				volume: 1.12443,
			},
			want: 1.6816846509973944,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &vwapCalculationEngine{
				size:                     tt.fields.size,
				prices:                   tt.fields.prices,
				volumes:                  tt.fields.volumes,
				cumulatedVolume:          tt.fields.cumulatedVolume,
				totalVolumeWeightedPrice: tt.fields.totalVolumeWeightedPrice,
				valueLimitSize:           tt.fields.valueLimitSize,
			}
			got := v.Calculate(tt.args.price, tt.args.volume)
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkReadCalculate(b *testing.B) {
	calcEngine := newVwapCalculationEngine(200)
	for i := 0; i < b.N; i++ {
		calcEngine.Calculate(0.07919, 0.00000473)
	}
}
