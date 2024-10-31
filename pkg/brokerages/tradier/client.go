package tradier

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Sackbuoy/trader/internal/types"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Configuration Configuration
	resty         *resty.Client
}

func CreateTradierClient(
	_ context.Context,
	config Configuration,
) (*Client, error) {
	restyClient := resty.New().
		SetBaseURL(config.URL).
		SetHeader("Accept", "application/json").
		SetAuthToken(config.Auth.AccessToken)

	return &Client{
		resty:         restyClient,
		Configuration: config,
	}, nil
}

func (c *Client) GetQuotes(ctx context.Context, symbols []string) ([]any, error) {
	symbolsQuery := strings.Join(symbols, ",")
	var result QuotesResponse

	req := c.resty.NewRequest().
		SetQueryParam("symbols", symbolsQuery).
		SetResult(&result)

	resp, err := req.Get("/v1/markets/quotes")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		msg := fmt.Sprintf("Failed to get quotes: Status Code %d, message %s", resp.StatusCode(), string(resp.Body()))
		return nil, errors.New(msg)
	}

	return nil, nil
}

func (c *Client) PlaceEquityOrder(ctx context.Context, equity types.Equity) error {
	var result EquityOrderResponse
	endpoint := fmt.Sprintf("/v1/accounts/%s/orders", c.Configuration.AccountID)

	req := c.resty.NewRequest().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetResult(&result).
		SetFormData(map[string]string{
			"class": "equity",
			"symbol": equity.Ticker,
			"side": "buy",
			"quantity": "1", // TODO: how do I decide this?
			"type": "market", // TODO: how do I decide this?
			"duration": "day", // TODO: how do I decide this?
		})

	resp, err := req.Post(endpoint)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		msg := fmt.Sprintf(
			"Buy failed with status: %d\nmsg: %s",
			resp.StatusCode(),
			string(resp.Body()),
		)
		return errors.New(msg)
	}

	return nil
}
