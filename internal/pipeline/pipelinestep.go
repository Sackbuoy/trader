package pipeline

import (
	"github.com/Sackbuoy/trader/internal/types"
	"golang.org/x/net/context"
)

type PipelineStep interface {
	Process(context.Context, []types.Equity) ([]types.Equity, error)
	Description() string
}
