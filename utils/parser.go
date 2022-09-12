package utils

import (
	"coinbase_vwap/domain"
	"github.com/goccy/go-json"
)

type jsonAdaptors interface {
	domain.MatchesResponse | domain.SubscribeToMatchesMessage
}

// ReadJson Generic function to read json.
func ReadJson[T jsonAdaptors](data []byte, x *T) error {
	err := json.Unmarshal(data, x)
	if err != nil {
		return domain.ErrUnmarshalJson
	}

	return nil
}

// MarshalJson Generic function to marshal json.
func MarshalJson[T jsonAdaptors](data T) ([]byte, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, domain.ErrMarshalJson
	}

	return bytes, nil
}
