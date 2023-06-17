// Code generated by Kitex v0.5.2. DO NOT EDIT.

package fileservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	file "summer/server/shared/kitex_gen/file"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UploadFile(ctx context.Context, req *file.UploadRequest, callOptions ...callopt.Option) (r *file.UploadResponse, err error)
	DownloadFile(ctx context.Context, req *file.DownloadRequest, callOptions ...callopt.Option) (r *file.DownloadResponse, err error)
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
	return &kFileServiceClient{
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

type kFileServiceClient struct {
	*kClient
}

func (p *kFileServiceClient) UploadFile(ctx context.Context, req *file.UploadRequest, callOptions ...callopt.Option) (r *file.UploadResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadFile(ctx, req)
}

func (p *kFileServiceClient) DownloadFile(ctx context.Context, req *file.DownloadRequest, callOptions ...callopt.Option) (r *file.DownloadResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DownloadFile(ctx, req)
}