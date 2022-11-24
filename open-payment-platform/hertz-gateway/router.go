// Code generated by hertz generator.

package main

import (
	"os"
	"strings"

	handler "github.com/cloudwego/biz-demo/open-payment-platform/hertz-gateway/biz/handler"
	"github.com/cloudwego/biz-demo/open-payment-platform/hertz-gateway/biz/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/registry-nacos/resolver"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	registerGateway(r)
}

func registerGateway(r *server.Hertz) {
	group := r.Group("/gateway").Use(middleware.GatewayAuth()...)

	if handler.SvcMap == nil {
		handler.SvcMap = make(map[string]genericclient.Client)
	}
	idlPath := "./idl/"
	c, err := os.ReadDir(idlPath)
	if err != nil {
		hlog.Fatalf("new thrift file provider failed: %v", err)
	}
	nacosResolver, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		hlog.Fatalf("err:%v", err)
	}

	for _, entry := range c {
		if entry.IsDir() || entry.Name() == "base.thrift" {
			continue
		}
		svcName := strings.ReplaceAll(entry.Name(), ".thrift", "")

		provider, err := generic.NewThriftFileProvider(entry.Name(), idlPath)
		if err != nil {
			hlog.Fatalf("new thrift file provider failed: %v", err)
			break
		}
		g, err := generic.HTTPThriftGeneric(provider)
		if err != nil {
			hlog.Fatal(err)
		}
		cli, err := genericclient.NewClient(
			svcName,
			g,
			client.WithResolver(nacosResolver),
			client.WithTransportProtocol(transport.TTHeader),
			client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		)
		if err != nil {
			hlog.Fatal(err)
		}

		handler.SvcMap[svcName] = cli
		group.POST("/:svc", handler.Gateway)
	}
}
