package tradier

type QuotesResponse struct {
	Quotes QuoteObject `json:"quotes"`
}

type QuoteObject struct {
	Quote []Quote `json:"quote"`
}

type Quote struct {
	Symbol string  `json:"symbol"`
	Last   float64 `json:"last"`
	Change float64 `json:"change"`
	Volume float64 `json:"volume"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
}

type EquityOrderResponse struct {
	Order EquityOrder `json:"order"`
}

type EquityOrder struct {
	ID int64 `json:"id"`
	Status string `json:"status"`
	PartnerID string `json:"partner_id"`
}
