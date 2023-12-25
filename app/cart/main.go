package main

import (
	"context"
	"net"
	"os"

	"github.com/baiyutang/gomall/app/cart/biz/dal"
	"github.com/baiyutang/gomall/app/cart/conf"
	"github.com/baiyutang/gomall/app/cart/infra/mtl"
	"github.com/baiyutang/gomall/app/cart/infra/rpc"
	"github.com/baiyutang/gomall/app/cart/kitex_gen/cart/cartservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	_ = godotenv.Load()
	rpc.InitClient()
	mtl.InitMtl()
	dal.Init()
	opts := kitexInit()

	svr := cartservice.NewServer(new(CartServiceImpl), opts...)

	err = svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	dal.Init()
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	if os.Getenv("REGISTRY_ENABLE") == "true" {
		r, err := consul.NewConsulRegister(os.Getenv("REGISTRY_ADDR"))
		if err != nil {
			klog.Fatal(err)
		}
		opts = append(opts, server.WithRegistry(r))
	}
	p := provider.NewOpenTelemetryProvider(
		provider.WithSdkTracerProvider(mtl.TracerProvider),
		provider.WithEnableMetrics(false),
	)
	defer p.Shutdown(context.Background())
	opts = append(opts, server.WithSuite(tracing.NewServerSuite()))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	klog.SetOutput(&lumberjack.Logger{
		Filename:   conf.GetConf().Kitex.LogFileName,
		MaxSize:    conf.GetConf().Kitex.LogMaxSize,
		MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
		MaxAge:     conf.GetConf().Kitex.LogMaxAge,
	})
	return
}
