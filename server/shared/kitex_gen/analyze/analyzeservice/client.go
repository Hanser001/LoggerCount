// Code generated by Kitex v0.5.2. DO NOT EDIT.

package analyzeservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	analyze "summer/server/shared/kitex_gen/analyze"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Analyze(ctx context.Context, req *analyze.AnalyzeRequest, callOptions ...callopt.Option) (r *analyze.AnalyzeResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kAnalyzeServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kAnalyzeServiceClient struct {
	*kClient
}

func (p *kAnalyzeServiceClient) Analyze(ctx context.Context, req *analyze.AnalyzeRequest, callOptions ...callopt.Option) (r *analyze.AnalyzeResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Analyze(ctx, req)
}