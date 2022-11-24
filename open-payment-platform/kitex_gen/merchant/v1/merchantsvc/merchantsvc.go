// Code generated by Kitex v0.4.2. DO NOT EDIT.

package merchantsvc

import (
	"context"
	v1 "github.com/cloudwego/biz-demo/open-payment-platform/kitex_gen/merchant/v1"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return merchantSvcServiceInfo
}

var merchantSvcServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "MerchantSvc"
	handlerType := (*v1.MerchantSvc)(nil)
	methods := map[string]kitex.MethodInfo{
		"Register":    kitex.NewMethodInfo(registerHandler, newMerchantSvcRegisterArgs, newMerchantSvcRegisterResult, false),
		"Report":      kitex.NewMethodInfo(reportHandler, newMerchantSvcReportArgs, newMerchantSvcReportResult, false),
		"QueryReport": kitex.NewMethodInfo(queryReportHandler, newMerchantSvcQueryReportArgs, newMerchantSvcQueryReportResult, false),
		"GetMerchant": kitex.NewMethodInfo(getMerchantHandler, newMerchantSvcGetMerchantArgs, newMerchantSvcGetMerchantResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "v1",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.2",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*v1.MerchantSvcRegisterArgs)
	realResult := result.(*v1.MerchantSvcRegisterResult)
	success, err := handler.(v1.MerchantSvc).Register(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMerchantSvcRegisterArgs() interface{} {
	return v1.NewMerchantSvcRegisterArgs()
}

func newMerchantSvcRegisterResult() interface{} {
	return v1.NewMerchantSvcRegisterResult()
}

func reportHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*v1.MerchantSvcReportArgs)
	realResult := result.(*v1.MerchantSvcReportResult)
	success, err := handler.(v1.MerchantSvc).Report(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMerchantSvcReportArgs() interface{} {
	return v1.NewMerchantSvcReportArgs()
}

func newMerchantSvcReportResult() interface{} {
	return v1.NewMerchantSvcReportResult()
}

func queryReportHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*v1.MerchantSvcQueryReportArgs)
	realResult := result.(*v1.MerchantSvcQueryReportResult)
	success, err := handler.(v1.MerchantSvc).QueryReport(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMerchantSvcQueryReportArgs() interface{} {
	return v1.NewMerchantSvcQueryReportArgs()
}

func newMerchantSvcQueryReportResult() interface{} {
	return v1.NewMerchantSvcQueryReportResult()
}

func getMerchantHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*v1.MerchantSvcGetMerchantArgs)
	realResult := result.(*v1.MerchantSvcGetMerchantResult)
	success, err := handler.(v1.MerchantSvc).GetMerchant(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMerchantSvcGetMerchantArgs() interface{} {
	return v1.NewMerchantSvcGetMerchantArgs()
}

func newMerchantSvcGetMerchantResult() interface{} {
	return v1.NewMerchantSvcGetMerchantResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, req *v1.RegisterReq) (r *v1.RegisterResp, err error) {
	var _args v1.MerchantSvcRegisterArgs
	_args.Req = req
	var _result v1.MerchantSvcRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Report(ctx context.Context, req *v1.ReportReq) (r *v1.ReportResp, err error) {
	var _args v1.MerchantSvcReportArgs
	_args.Req = req
	var _result v1.MerchantSvcReportResult
	if err = p.c.Call(ctx, "Report", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryReport(ctx context.Context, req *v1.QueryReportReq) (r *v1.QueryReportResp, err error) {
	var _args v1.MerchantSvcQueryReportArgs
	_args.Req = req
	var _result v1.MerchantSvcQueryReportResult
	if err = p.c.Call(ctx, "QueryReport", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetMerchant(ctx context.Context, req *v1.GetMerchantReq) (r *v1.GetMerchantResp, err error) {
	var _args v1.MerchantSvcGetMerchantArgs
	_args.Req = req
	var _result v1.MerchantSvcGetMerchantResult
	if err = p.c.Call(ctx, "GetMerchant", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
