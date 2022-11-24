// Code generated by Kitex v0.4.2. DO NOT EDIT.

package merchantsvc

import (
	"context"
	v1 "github.com/cloudwego/biz-demo/open-payment-platform/kitex_gen/merchant/v1"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Register(ctx context.Context, req *v1.RegisterReq, callOptions ...callopt.Option) (r *v1.RegisterResp, err error)
	Report(ctx context.Context, req *v1.ReportReq, callOptions ...callopt.Option) (r *v1.ReportResp, err error)
	QueryReport(ctx context.Context, req *v1.QueryReportReq, callOptions ...callopt.Option) (r *v1.QueryReportResp, err error)
	GetMerchant(ctx context.Context, req *v1.GetMerchantReq, callOptions ...callopt.Option) (r *v1.GetMerchantResp, err error)
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
	return &kMerchantSvcClient{
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

type kMerchantSvcClient struct {
	*kClient
}

func (p *kMerchantSvcClient) Register(ctx context.Context, req *v1.RegisterReq, callOptions ...callopt.Option) (r *v1.RegisterResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, req)
}

func (p *kMerchantSvcClient) Report(ctx context.Context, req *v1.ReportReq, callOptions ...callopt.Option) (r *v1.ReportResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Report(ctx, req)
}

func (p *kMerchantSvcClient) QueryReport(ctx context.Context, req *v1.QueryReportReq, callOptions ...callopt.Option) (r *v1.QueryReportResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryReport(ctx, req)
}

func (p *kMerchantSvcClient) GetMerchant(ctx context.Context, req *v1.GetMerchantReq, callOptions ...callopt.Option) (r *v1.GetMerchantResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMerchant(ctx, req)
}
