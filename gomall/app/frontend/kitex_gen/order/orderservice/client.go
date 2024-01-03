// Code generated by Kitex v0.8.0. DO NOT EDIT.

package orderservice

import (
	"context"
	order "github.com/cloudwego/biz-demo/gomall/app/frontend/kitex_gen/order"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	PlaceOrder(ctx context.Context, Req *order.PlaceOrderRequest, callOptions ...callopt.Option) (r *order.PlaceOrderResponse, err error)
	ListOrder(ctx context.Context, Req *order.ListOrderRequest, callOptions ...callopt.Option) (r *order.ListOrderResponse, err error)
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
	return &kOrderServiceClient{
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

type kOrderServiceClient struct {
	*kClient
}

func (p *kOrderServiceClient) PlaceOrder(ctx context.Context, Req *order.PlaceOrderRequest, callOptions ...callopt.Option) (r *order.PlaceOrderResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PlaceOrder(ctx, Req)
}

func (p *kOrderServiceClient) ListOrder(ctx context.Context, Req *order.ListOrderRequest, callOptions ...callopt.Option) (r *order.ListOrderResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListOrder(ctx, Req)
}
