package websocket

import (
	"coinbase_vwap/domain"
	"coinbase_vwap/utils"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type socketClient struct {
	conn    *websocket.Conn
	address string
}

// NewSocketClient will create new an socketClient object representation of domain.WSClient interface.
func NewSocketClient(address string) (domain.WSClient, error) {
	client := &socketClient{
		address: address,
	}
	err := client.dial()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (s *socketClient) dial() error {
	conn, _, err := websocket.DefaultDialer.Dial(s.address, nil)
	s.conn = conn
	if err != nil {
		return domain.ErrWebSocketConnection
	}

	return nil
}

// SubscribeAndRead subscribe to matches channel and listen to the messages.
func (s *socketClient) SubscribeAndRead(productIds []string, errChannel chan<- error, matchesChan chan<- domain.MatchesResponse) {
	err := s.subscribeMatches(productIds)
	if err != nil {
		errChannel <- err
	}

	s.messageListener(matchesChan, errChannel)
}

func (s *socketClient) subscribeMatches(productIds []string) error {
	subscriptionMessage, err := getMatchesSubscribeMsg(productIds)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = s.conn.WriteMessage(websocket.TextMessage, subscriptionMessage)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *socketClient) messageListener(matchesChan chan<- domain.MatchesResponse, errChan chan<- error) {
	var matchesResponse domain.MatchesResponse
	failedToReadJsonCount := 0
	failedToConnect := 0

	for {
		_, msg, err := s.conn.ReadMessage()
		if err != nil {
			if failedToConnect >= 3 {
				errChan <- domain.ErrWebSocketConnectionLost
				return
			}
			time.Sleep(2 * time.Second)
			err = s.dial()
			if err != nil {
				failedToConnect += 1
			}
		}
		err = utils.ReadJson(msg, &matchesResponse)
		if err != nil {
			failedToReadJsonCount += 1
			if failedToReadJsonCount >= 3 {
				errChan <- err
				return
			}
		}

		if matchesResponse.Type == typeNotAllowed {
			continue
		}

		matchesChan <- matchesResponse
	}

}

// Close Just close the connection.
func (s *socketClient) Close() error {
	return s.conn.Close()
}
