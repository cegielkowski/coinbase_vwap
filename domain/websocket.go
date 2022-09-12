package domain

type WSClient interface {
	SubscribeAndRead(productIds []string, errChannel chan<- error, matchesChan chan<- MatchesResponse)
	Close() error
}
