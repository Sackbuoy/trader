package avgcompare

import (
	"context"
	"os"
	"time"

	"github.com/Sackbuoy/trader/internal/pipeline"
	"github.com/Sackbuoy/trader/internal/types"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

var _ pipeline.PipelineStep = (*AvgCompare)(nil)

type AvgCompare struct {
	Configuration Configuration
	Client        *polygon.Client
}

func New(_ context.Context, config Configuration) (*AvgCompare, error) {
	polygonClient := polygon.New(os.Getenv("POLYGON_API_KEY"))

	return &AvgCompare{
		Configuration: config,
		Client:        polygonClient,
	}, nil
}

func (a AvgCompare) Description() string {
	return "Compares the averages over a short vs a long period of time, and marks positions to buy if short > long"
}

func (a AvgCompare) Process(_ context.Context, input []types.Equity) ([]types.Equity, error) {
	returnVal := make([]types.Equity, 0)

	for _, v := range input {
		// get long first, then use that data to get the short
		// that way theres only 1 API request
		aggs, err := a.getAggsForTimeRange(v, a.Configuration.Long)
		if err != nil {
			return nil, err
		}

		shortVal := int(a.Configuration.Short.Hours())
		short, err := a.averageAggs(aggs, shortVal)
		if err != nil {
			return nil, err
		}

		long, err := a.averageAggs(aggs, len(aggs))
		if err != nil {
			return nil, err
		}

		if short > long {
			v.Action = "BUY"
			returnVal = append(returnVal, v)
		}
	}

	return returnVal, nil
}

func (a AvgCompare) averageAggs(aggs []models.Agg, num int) (float64, error) {
	var tmp float64
	if num >= len(aggs) {
		num = len(aggs)
	}
	for i := 0; i < num-1; i++ {
		tmp += aggs[i].VWAP
	}

	return tmp / float64(num), nil
}

func (a AvgCompare) getAggsForTimeRange(equity types.Equity, timeRange time.Duration) ([]models.Agg, error) {
	now := time.Now()

	// number of aggs returned here is: a.Configuration.Short/Timespan
	params := models.ListAggsParams{
		Ticker:     equity.Ticker,
		Multiplier: 1,
		// make configurable? this means each average is over 1 day
		Timespan: models.Timespan(models.Hour),
		To:       models.Millis(now),
		From:     models.Millis(now.Add(-timeRange)),
	}
	aggsIterator := a.Client.ListAggs(context.Background(), &params)

	cur := aggsIterator.Item()
	result := []models.Agg{
		cur,
	}
	// numItems := a.Configuration.Short.Hours()/24
	//
	// var tmp = cur.VWAP
	for aggsIterator.Next() {
		// Volume Weighted Aggregate Price: sum of all recorded prices divided by
		// number of trades in timespan
		// tmp += cur.VWAP
		cur = aggsIterator.Item()
		result = append(result, cur)
	}

	return result, nil
}
