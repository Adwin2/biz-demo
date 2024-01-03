// Code generated by Kitex v0.6.1. DO NOT EDIT.

package paymentservice

import (
	"context"
	"fmt"
	payment "github.com/cloudwego/biz-demo/gomall/app/order/kitex_gen/payment"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return paymentServiceServiceInfo
}

var paymentServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "PaymentService"
	handlerType := (*payment.PaymentService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Charge": kitex.NewMethodInfo(chargeHandler, newChargeArgs, newChargeResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "payment",
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

func chargeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(payment.ChargeRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(payment.PaymentService).Charge(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ChargeArgs:
		success, err := handler.(payment.PaymentService).Charge(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ChargeResult)
		realResult.Success = success
	}
	return nil
}
func newChargeArgs() interface{} {
	return &ChargeArgs{}
}

func newChargeResult() interface{} {
	return &ChargeResult{}
}

type ChargeArgs struct {
	Req *payment.ChargeRequest
}

func (p *ChargeArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(payment.ChargeRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ChargeArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ChargeArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ChargeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChargeArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ChargeArgs) Unmarshal(in []byte) error {
	msg := new(payment.ChargeRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ChargeArgs_Req_DEFAULT *payment.ChargeRequest

func (p *ChargeArgs) GetReq() *payment.ChargeRequest {
	if !p.IsSetReq() {
		return ChargeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChargeArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ChargeArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ChargeResult struct {
	Success *payment.ChargeResponse
}

var ChargeResult_Success_DEFAULT *payment.ChargeResponse

func (p *ChargeResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(payment.ChargeResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ChargeResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ChargeResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ChargeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChargeResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ChargeResult) Unmarshal(in []byte) error {
	msg := new(payment.ChargeResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChargeResult) GetSuccess() *payment.ChargeResponse {
	if !p.IsSetSuccess() {
		return ChargeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChargeResult) SetSuccess(x interface{}) {
	p.Success = x.(*payment.ChargeResponse)
}

func (p *ChargeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChargeResult) GetResult() interface{} {
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

func (p *kClient) Charge(ctx context.Context, Req *payment.ChargeRequest) (r *payment.ChargeResponse, err error) {
	var _args ChargeArgs
	_args.Req = Req
	var _result ChargeResult
	if err = p.c.Call(ctx, "Charge", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
