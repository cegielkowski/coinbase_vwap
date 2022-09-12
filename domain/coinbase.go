package domain

type MatchesResponse struct {
	Type      string  `json:"type"`
	Size      float64 `json:"size,string"`
	Price     float64 `json:"price,string"`
	ProductId string  `json:"product_id"`
}

type Channel struct {
	Name       string   `json:"name"`
	ProductIds []string `json:"product_ids"`
}

type SubscribeToMatchesMessage struct {
	Type     string    `json:"type"`
	Channels []Channel `json:"channels"`
}

type Products map[string]VwapCalculationEngine
