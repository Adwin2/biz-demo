// Code generated by Kitex v0.6.1. DO NOT EDIT.

package orderservice

import (
	"context"
	"fmt"
	order "github.com/cloudwego/biz-demo/gomall/app/checkout/kitex_gen/order"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return orderServiceServiceInfo
}

var orderServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "OrderService"
	handlerType := (*order.OrderService)(nil)
	methods := map[string]kitex.MethodInfo{
		"PlaceOrder": kitex.NewMethodInfo(placeOrderHandler, newPlaceOrderArgs, newPlaceOrderResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "order",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func placeOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.PlaceOrderRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).PlaceOrder(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PlaceOrderArgs:
		success, err := handler.(order.OrderService).PlaceOrder(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PlaceOrderResult)
		realResult.Success = success
	}
	return nil
}
func newPlaceOrderArgs() interface{} {
	return &PlaceOrderArgs{}
}

func newPlaceOrderResult() interface{} {
	return &PlaceOrderResult{}
}

type PlaceOrderArgs struct {
	Req *order.PlaceOrderRequest
}

func (p *PlaceOrderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.PlaceOrderRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PlaceOrderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PlaceOrderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PlaceOrderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PlaceOrderArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PlaceOrderArgs) Unmarshal(in []byte) error {
	msg := new(order.PlaceOrderRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PlaceOrderArgs_Req_DEFAULT *order.PlaceOrderRequest

func (p *PlaceOrderArgs) GetReq() *order.PlaceOrderRequest {
	if !p.IsSetReq() {
		return PlaceOrderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PlaceOrderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PlaceOrderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PlaceOrderResult struct {
	Success *order.PlaceOrderResponse
}

var PlaceOrderResult_Success_DEFAULT *order.PlaceOrderResponse

func (p *PlaceOrderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.PlaceOrderResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PlaceOrderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PlaceOrderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PlaceOrderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PlaceOrderResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PlaceOrderResult) Unmarshal(in []byte) error {
	msg := new(order.PlaceOrderResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PlaceOrderResult) GetSuccess() *order.PlaceOrderResponse {
	if !p.IsSetSuccess() {
		return PlaceOrderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PlaceOrderResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.PlaceOrderResponse)
}

func (p *PlaceOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PlaceOrderResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) PlaceOrder(ctx context.Context, Req *order.PlaceOrderRequest) (r *order.PlaceOrderResponse, err error) {
	var _args PlaceOrderArgs
	_args.Req = Req
	var _result PlaceOrderResult
	if err = p.c.Call(ctx, "PlaceOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
