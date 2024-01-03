// Code generated by Kitex v0.4.3. DO NOT EDIT.

package userservice

import (
	"context"
	"github.com/cloudwego/biz-demo/easy_note/kitex_gen/demouser"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateUser(ctx context.Context, req *demouser.CreateUserRequest, callOptions ...callopt.Option) (r *demouser.CreateUserResponse, err error)
	MGetUser(ctx context.Context, req *demouser.MGetUserRequest, callOptions ...callopt.Option) (r *demouser.MGetUserResponse, err error)
	CheckUser(ctx context.Context, req *demouser.CheckUserRequest, callOptions ...callopt.Option) (r *demouser.CheckUserResponse, err error)
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
	return &kUserServiceClient{
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

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) CreateUser(ctx context.Context, req *demouser.CreateUserRequest, callOptions ...callopt.Option) (r *demouser.CreateUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUser(ctx, req)
}

func (p *kUserServiceClient) MGetUser(ctx context.Context, req *demouser.MGetUserRequest, callOptions ...callopt.Option) (r *demouser.MGetUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MGetUser(ctx, req)
}

func (p *kUserServiceClient) CheckUser(ctx context.Context, req *demouser.CheckUserRequest, callOptions ...callopt.Option) (r *demouser.CheckUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckUser(ctx, req)
}
