package coinbase

import (
	"coinbase_vwap/domain"
	"log"
)

// HandleChannels Handle the channel responses and print the vwap.
func HandleChannels(matchesChan <-chan domain.MatchesResponse, products domain.Products, errChan <-chan error) {
	for {
		select {
		case match := <-matchesChan:
			vwapCalculated := products[match.ProductId].Calculate(match.Price, match.Size)
			log.Printf("TradeID: %s, VWAP: %f", match.ProductId, vwapCalculated)
		case err := <-errChan:
			log.Println(domain.ErrInternalServerError.Error())
			log.Println(err)
			return
		}
	}
}
