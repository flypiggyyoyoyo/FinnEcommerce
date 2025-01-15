// Code generated by Kitex v0.9.1. DO NOT EDIT.

package orderservice

import (
	"context"
	"errors"
	order "github.com/FinnTew/FinnEcommerce/src/internal/service/order/kitex_gen/order"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"PlaceOrder": kitex.NewMethodInfo(
		placeOrderHandler,
		newPlaceOrderArgs,
		newPlaceOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ListOrder": kitex.NewMethodInfo(
		listOrderHandler,
		newListOrderArgs,
		newListOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"MarkOrderPaid": kitex.NewMethodInfo(
		markOrderPaidHandler,
		newMarkOrderPaidArgs,
		newMarkOrderPaidResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"CancelOrder": kitex.NewMethodInfo(
		cancelOrderHandler,
		newCancelOrderArgs,
		newCancelOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateOrder": kitex.NewMethodInfo(
		updateOrderHandler,
		newUpdateOrderArgs,
		newUpdateOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	orderServiceServiceInfo                = NewServiceInfo()
	orderServiceServiceInfoForClient       = NewServiceInfoForClient()
	orderServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return orderServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return orderServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return orderServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "OrderService"
	handlerType := (*order.OrderService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "order",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func placeOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.PlaceOrderReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).PlaceOrder(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *PlaceOrderArgs:
		success, err := handler.(order.OrderService).PlaceOrder(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PlaceOrderResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newPlaceOrderArgs() interface{} {
	return &PlaceOrderArgs{}
}

func newPlaceOrderResult() interface{} {
	return &PlaceOrderResult{}
}

type PlaceOrderArgs struct {
	Req *order.PlaceOrderReq
}

func (p *PlaceOrderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.PlaceOrderReq)
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
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *PlaceOrderArgs) Unmarshal(in []byte) error {
	msg := new(order.PlaceOrderReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PlaceOrderArgs_Req_DEFAULT *order.PlaceOrderReq

func (p *PlaceOrderArgs) GetReq() *order.PlaceOrderReq {
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
	Success *order.PlaceOrderResp
}

var PlaceOrderResult_Success_DEFAULT *order.PlaceOrderResp

func (p *PlaceOrderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.PlaceOrderResp)
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
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *PlaceOrderResult) Unmarshal(in []byte) error {
	msg := new(order.PlaceOrderResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PlaceOrderResult) GetSuccess() *order.PlaceOrderResp {
	if !p.IsSetSuccess() {
		return PlaceOrderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PlaceOrderResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.PlaceOrderResp)
}

func (p *PlaceOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PlaceOrderResult) GetResult() interface{} {
	return p.Success
}

func listOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.ListOrderReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).ListOrder(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListOrderArgs:
		success, err := handler.(order.OrderService).ListOrder(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListOrderResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListOrderArgs() interface{} {
	return &ListOrderArgs{}
}

func newListOrderResult() interface{} {
	return &ListOrderResult{}
}

type ListOrderArgs struct {
	Req *order.ListOrderReq
}

func (p *ListOrderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.ListOrderReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListOrderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListOrderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListOrderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListOrderArgs) Unmarshal(in []byte) error {
	msg := new(order.ListOrderReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListOrderArgs_Req_DEFAULT *order.ListOrderReq

func (p *ListOrderArgs) GetReq() *order.ListOrderReq {
	if !p.IsSetReq() {
		return ListOrderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListOrderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListOrderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListOrderResult struct {
	Success *order.ListOrderResp
}

var ListOrderResult_Success_DEFAULT *order.ListOrderResp

func (p *ListOrderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.ListOrderResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListOrderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListOrderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListOrderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListOrderResult) Unmarshal(in []byte) error {
	msg := new(order.ListOrderResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListOrderResult) GetSuccess() *order.ListOrderResp {
	if !p.IsSetSuccess() {
		return ListOrderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListOrderResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.ListOrderResp)
}

func (p *ListOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListOrderResult) GetResult() interface{} {
	return p.Success
}

func markOrderPaidHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.MarkOrderPaidReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).MarkOrderPaid(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *MarkOrderPaidArgs:
		success, err := handler.(order.OrderService).MarkOrderPaid(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MarkOrderPaidResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newMarkOrderPaidArgs() interface{} {
	return &MarkOrderPaidArgs{}
}

func newMarkOrderPaidResult() interface{} {
	return &MarkOrderPaidResult{}
}

type MarkOrderPaidArgs struct {
	Req *order.MarkOrderPaidReq
}

func (p *MarkOrderPaidArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.MarkOrderPaidReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *MarkOrderPaidArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *MarkOrderPaidArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *MarkOrderPaidArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *MarkOrderPaidArgs) Unmarshal(in []byte) error {
	msg := new(order.MarkOrderPaidReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MarkOrderPaidArgs_Req_DEFAULT *order.MarkOrderPaidReq

func (p *MarkOrderPaidArgs) GetReq() *order.MarkOrderPaidReq {
	if !p.IsSetReq() {
		return MarkOrderPaidArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MarkOrderPaidArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *MarkOrderPaidArgs) GetFirstArgument() interface{} {
	return p.Req
}

type MarkOrderPaidResult struct {
	Success *order.MarkOrderPaidResp
}

var MarkOrderPaidResult_Success_DEFAULT *order.MarkOrderPaidResp

func (p *MarkOrderPaidResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.MarkOrderPaidResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *MarkOrderPaidResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *MarkOrderPaidResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *MarkOrderPaidResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *MarkOrderPaidResult) Unmarshal(in []byte) error {
	msg := new(order.MarkOrderPaidResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MarkOrderPaidResult) GetSuccess() *order.MarkOrderPaidResp {
	if !p.IsSetSuccess() {
		return MarkOrderPaidResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MarkOrderPaidResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.MarkOrderPaidResp)
}

func (p *MarkOrderPaidResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MarkOrderPaidResult) GetResult() interface{} {
	return p.Success
}

func cancelOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.CancelOrderReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).CancelOrder(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CancelOrderArgs:
		success, err := handler.(order.OrderService).CancelOrder(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CancelOrderResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCancelOrderArgs() interface{} {
	return &CancelOrderArgs{}
}

func newCancelOrderResult() interface{} {
	return &CancelOrderResult{}
}

type CancelOrderArgs struct {
	Req *order.CancelOrderReq
}

func (p *CancelOrderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.CancelOrderReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CancelOrderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CancelOrderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CancelOrderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CancelOrderArgs) Unmarshal(in []byte) error {
	msg := new(order.CancelOrderReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CancelOrderArgs_Req_DEFAULT *order.CancelOrderReq

func (p *CancelOrderArgs) GetReq() *order.CancelOrderReq {
	if !p.IsSetReq() {
		return CancelOrderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CancelOrderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CancelOrderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CancelOrderResult struct {
	Success *order.CancelOrderResp
}

var CancelOrderResult_Success_DEFAULT *order.CancelOrderResp

func (p *CancelOrderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.CancelOrderResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CancelOrderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CancelOrderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CancelOrderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CancelOrderResult) Unmarshal(in []byte) error {
	msg := new(order.CancelOrderResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CancelOrderResult) GetSuccess() *order.CancelOrderResp {
	if !p.IsSetSuccess() {
		return CancelOrderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CancelOrderResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.CancelOrderResp)
}

func (p *CancelOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CancelOrderResult) GetResult() interface{} {
	return p.Success
}

func updateOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.UpdateOrderReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).UpdateOrder(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateOrderArgs:
		success, err := handler.(order.OrderService).UpdateOrder(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateOrderResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateOrderArgs() interface{} {
	return &UpdateOrderArgs{}
}

func newUpdateOrderResult() interface{} {
	return &UpdateOrderResult{}
}

type UpdateOrderArgs struct {
	Req *order.UpdateOrderReq
}

func (p *UpdateOrderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.UpdateOrderReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateOrderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateOrderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateOrderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateOrderArgs) Unmarshal(in []byte) error {
	msg := new(order.UpdateOrderReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateOrderArgs_Req_DEFAULT *order.UpdateOrderReq

func (p *UpdateOrderArgs) GetReq() *order.UpdateOrderReq {
	if !p.IsSetReq() {
		return UpdateOrderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateOrderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateOrderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateOrderResult struct {
	Success *order.UpdateOrderResp
}

var UpdateOrderResult_Success_DEFAULT *order.UpdateOrderResp

func (p *UpdateOrderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.UpdateOrderResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateOrderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateOrderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateOrderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateOrderResult) Unmarshal(in []byte) error {
	msg := new(order.UpdateOrderResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateOrderResult) GetSuccess() *order.UpdateOrderResp {
	if !p.IsSetSuccess() {
		return UpdateOrderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateOrderResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.UpdateOrderResp)
}

func (p *UpdateOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateOrderResult) GetResult() interface{} {
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

func (p *kClient) PlaceOrder(ctx context.Context, Req *order.PlaceOrderReq) (r *order.PlaceOrderResp, err error) {
	var _args PlaceOrderArgs
	_args.Req = Req
	var _result PlaceOrderResult
	if err = p.c.Call(ctx, "PlaceOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListOrder(ctx context.Context, Req *order.ListOrderReq) (r *order.ListOrderResp, err error) {
	var _args ListOrderArgs
	_args.Req = Req
	var _result ListOrderResult
	if err = p.c.Call(ctx, "ListOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MarkOrderPaid(ctx context.Context, Req *order.MarkOrderPaidReq) (r *order.MarkOrderPaidResp, err error) {
	var _args MarkOrderPaidArgs
	_args.Req = Req
	var _result MarkOrderPaidResult
	if err = p.c.Call(ctx, "MarkOrderPaid", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CancelOrder(ctx context.Context, Req *order.CancelOrderReq) (r *order.CancelOrderResp, err error) {
	var _args CancelOrderArgs
	_args.Req = Req
	var _result CancelOrderResult
	if err = p.c.Call(ctx, "CancelOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateOrder(ctx context.Context, Req *order.UpdateOrderReq) (r *order.UpdateOrderResp, err error) {
	var _args UpdateOrderArgs
	_args.Req = Req
	var _result UpdateOrderResult
	if err = p.c.Call(ctx, "UpdateOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
