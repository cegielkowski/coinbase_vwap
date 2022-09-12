package utils

import (
	"coinbase_vwap/domain"
	"testing"
)

func BenchmarkReadJson(b *testing.B) {
	var toBench = []byte(`{"type":"last_match","trade_id":28734370,"maker_order_id":"43b63a3f-24e9-4e9e-bf4b-dd05fe8e7f2b","taker_order_id":"04e5e396-cee4-4a4f-9bed-bd213b2f69eb","side":"sell","size":"0.00116169","price":"0.08064","product_id":"ETH-BTC","sequence":5793897281,"time":"2022-09-12T00:46:21.388844Z"}`)
	for i := 0; i < b.N; i++ {
		var toUnmarshal domain.MatchesResponse
		_ = ReadJson(toBench, &toUnmarshal)
	}
}

func BenchmarkMarshalJson(b *testing.B) {
	teste := domain.SubscribeToMatchesMessage{
		Type: "subscribe",
		Channels: []domain.Channel{{
			Name:       "matches",
			ProductIds: nil,
		}},
	}
	for i := 0; i < b.N; i++ {
		_, _ = MarshalJson(teste)
	}
}
