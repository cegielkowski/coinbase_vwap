package websocket

import (
	"coinbase_vwap/domain"
	"coinbase_vwap/utils"
)

var (
	matchesType    = "subscribe"
	matchesName    = "matches"
	typeNotAllowed = "subscriptions"
)

func getMatchesSubscribeMsg(productIds []string) ([]byte, error) {
	subscribeToMatchesMessage := domain.SubscribeToMatchesMessage{
		Type: matchesType,
		Channels: []domain.Channel{{
			Name:       matchesName,
			ProductIds: productIds,
		}},
	}
	result, err := utils.MarshalJson(subscribeToMatchesMessage)
	return result, err
}
