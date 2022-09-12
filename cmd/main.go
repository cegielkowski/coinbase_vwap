package main

import (
	"coinbase_vwap/coinbase"
	"coinbase_vwap/coinbase/vwap"
	"coinbase_vwap/coinbase/websocket"
	"coinbase_vwap/config"
	"coinbase_vwap/domain"
	"log"
	"os"
)

func main() {
	// Load config.
	appConfig := config.LoadConfig()
	// Start socket connection.
	client, err := websocket.NewSocketClient(appConfig.WebServiceAddress)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	defer client.Close()

	// Create channels.
	errChan := make(chan error)
	matchesChan := make(chan domain.MatchesResponse)

	// Setup engines.
	calculationEngines := vwap.NewVwapCalculationEngines(appConfig.ValueLimitSize, len(appConfig.ProductIds))
	// Generate product structures.
	products := vwap.NewProducts(appConfig.ProductIds, calculationEngines)

	// Start to listen to messages.
	go client.SubscribeAndRead(appConfig.ProductIds, errChan, matchesChan)

	// Handle channels with responses or errors and calculate the vwap.
	coinbase.HandleChannels(matchesChan, products, errChan)
}
