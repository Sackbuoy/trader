package types

import "fmt"

type Equity struct {
	Ticker    string
	Action    string
	MarketCap float64
}

func (e Equity) Print() {
	fmt.Printf("Ticker: %s\n", e.Ticker)
	fmt.Printf("Action: %s\n", e.Action)
	fmt.Printf("Market Cap: %f\n", e.MarketCap)
}
