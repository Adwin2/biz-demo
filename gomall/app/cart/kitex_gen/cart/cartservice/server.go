// Code generated by Kitex v0.8.0. DO NOT EDIT.
package cartservice

import (
	cart "github.com/cloudwego/biz-demo/gomall/app/cart/kitex_gen/cart"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler cart.CartService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
