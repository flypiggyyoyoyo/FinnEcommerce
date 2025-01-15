// Code generated by Kitex v0.9.1. DO NOT EDIT.

package emailservice

import (
	"context"
	"errors"
	email "github.com/FinnTew/FinnEcommerce/src/internal/service/email/kitex_gen/email"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"SendEmail": kitex.NewMethodInfo(
		sendEmailHandler,
		newSendEmailArgs,
		newSendEmailResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	emailServiceServiceInfo                = NewServiceInfo()
	emailServiceServiceInfoForClient       = NewServiceInfoForClient()
	emailServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return emailServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return emailServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return emailServiceServiceInfoForClient
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
	serviceName := "EmailService"
	handlerType := (*email.EmailService)(nil)
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
		"PackageName": "email",
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

func sendEmailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(email.SendEmailRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(email.EmailService).SendEmail(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *SendEmailArgs:
		success, err := handler.(email.EmailService).SendEmail(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SendEmailResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newSendEmailArgs() interface{} {
	return &SendEmailArgs{}
}

func newSendEmailResult() interface{} {
	return &SendEmailResult{}
}

type SendEmailArgs struct {
	Req *email.SendEmailRequest
}

func (p *SendEmailArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(email.SendEmailRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SendEmailArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SendEmailArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SendEmailArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SendEmailArgs) Unmarshal(in []byte) error {
	msg := new(email.SendEmailRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SendEmailArgs_Req_DEFAULT *email.SendEmailRequest

func (p *SendEmailArgs) GetReq() *email.SendEmailRequest {
	if !p.IsSetReq() {
		return SendEmailArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SendEmailArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SendEmailArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SendEmailResult struct {
	Success *email.SendEmailResponse
}

var SendEmailResult_Success_DEFAULT *email.SendEmailResponse

func (p *SendEmailResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(email.SendEmailResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SendEmailResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SendEmailResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SendEmailResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SendEmailResult) Unmarshal(in []byte) error {
	msg := new(email.SendEmailResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendEmailResult) GetSuccess() *email.SendEmailResponse {
	if !p.IsSetSuccess() {
		return SendEmailResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SendEmailResult) SetSuccess(x interface{}) {
	p.Success = x.(*email.SendEmailResponse)
}

func (p *SendEmailResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SendEmailResult) GetResult() interface{} {
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

func (p *kClient) SendEmail(ctx context.Context, Req *email.SendEmailRequest) (r *email.SendEmailResponse, err error) {
	var _args SendEmailArgs
	_args.Req = Req
	var _result SendEmailResult
	if err = p.c.Call(ctx, "SendEmail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
