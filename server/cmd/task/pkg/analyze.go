package pkg

import (
	"context"
	"summer/server/shared/kitex_gen/analyze/analyzeservice"
)

type AnalyzeManger struct {
	ac analyzeservice.Client
}

func NewAnalyzeManger(client analyzeservice.Client) *AnalyzeManger {
	return &AnalyzeManger{ac: client}
}

func (a AnalyzeManger) Analyze(ctx context.Context) {

	panic("implement me")
}
