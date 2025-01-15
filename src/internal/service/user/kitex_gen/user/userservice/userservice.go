// Code generated by Kitex v0.9.1. DO NOT EDIT.

package userservice

import (
	"context"
	"errors"
	user "github.com/FinnTew/FinnEcommerce/src/internal/service/user/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Register": kitex.NewMethodInfo(
		registerHandler,
		newRegisterArgs,
		newRegisterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Login": kitex.NewMethodInfo(
		loginHandler,
		newLoginArgs,
		newLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Logout": kitex.NewMethodInfo(
		logoutHandler,
		newLogoutArgs,
		newLogoutResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Delete": kitex.NewMethodInfo(
		deleteHandler,
		newDeleteArgs,
		newDeleteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Update": kitex.NewMethodInfo(
		updateHandler,
		newUpdateArgs,
		newUpdateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetUserInfo": kitex.NewMethodInfo(
		getUserInfoHandler,
		newGetUserInfoArgs,
		newGetUserInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	userServiceServiceInfo                = NewServiceInfo()
	userServiceServiceInfoForClient       = NewServiceInfoForClient()
	userServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForClient
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
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
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
		"PackageName": "user",
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

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.RegisterReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).Register(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *RegisterArgs:
		success, err := handler.(user.UserService).Register(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RegisterResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newRegisterArgs() interface{} {
	return &RegisterArgs{}
}

func newRegisterResult() interface{} {
	return &RegisterResult{}
}

type RegisterArgs struct {
	Req *user.RegisterReq
}

func (p *RegisterArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.RegisterReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RegisterArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RegisterArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RegisterArgs) Unmarshal(in []byte) error {
	msg := new(user.RegisterReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RegisterArgs_Req_DEFAULT *user.RegisterReq

func (p *RegisterArgs) GetReq() *user.RegisterReq {
	if !p.IsSetReq() {
		return RegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RegisterArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RegisterResult struct {
	Success *user.RegisterResp
}

var RegisterResult_Success_DEFAULT *user.RegisterResp

func (p *RegisterResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.RegisterResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RegisterResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RegisterResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RegisterResult) Unmarshal(in []byte) error {
	msg := new(user.RegisterResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RegisterResult) GetSuccess() *user.RegisterResp {
	if !p.IsSetSuccess() {
		return RegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.RegisterResp)
}

func (p *RegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RegisterResult) GetResult() interface{} {
	return p.Success
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.LoginReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).Login(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *LoginArgs:
		success, err := handler.(user.UserService).Login(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*LoginResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newLoginArgs() interface{} {
	return &LoginArgs{}
}

func newLoginResult() interface{} {
	return &LoginResult{}
}

type LoginArgs struct {
	Req *user.LoginReq
}

func (p *LoginArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.LoginReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *LoginArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *LoginArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *LoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *LoginArgs) Unmarshal(in []byte) error {
	msg := new(user.LoginReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var LoginArgs_Req_DEFAULT *user.LoginReq

func (p *LoginArgs) GetReq() *user.LoginReq {
	if !p.IsSetReq() {
		return LoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *LoginArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *LoginArgs) GetFirstArgument() interface{} {
	return p.Req
}

type LoginResult struct {
	Success *user.LoginResp
}

var LoginResult_Success_DEFAULT *user.LoginResp

func (p *LoginResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.LoginResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *LoginResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *LoginResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *LoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *LoginResult) Unmarshal(in []byte) error {
	msg := new(user.LoginResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *LoginResult) GetSuccess() *user.LoginResp {
	if !p.IsSetSuccess() {
		return LoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *LoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.LoginResp)
}

func (p *LoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *LoginResult) GetResult() interface{} {
	return p.Success
}

func logoutHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.LogoutReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).Logout(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *LogoutArgs:
		success, err := handler.(user.UserService).Logout(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*LogoutResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newLogoutArgs() interface{} {
	return &LogoutArgs{}
}

func newLogoutResult() interface{} {
	return &LogoutResult{}
}

type LogoutArgs struct {
	Req *user.LogoutReq
}

func (p *LogoutArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.LogoutReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *LogoutArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *LogoutArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *LogoutArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *LogoutArgs) Unmarshal(in []byte) error {
	msg := new(user.LogoutReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var LogoutArgs_Req_DEFAULT *user.LogoutReq

func (p *LogoutArgs) GetReq() *user.LogoutReq {
	if !p.IsSetReq() {
		return LogoutArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *LogoutArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *LogoutArgs) GetFirstArgument() interface{} {
	return p.Req
}

type LogoutResult struct {
	Success *user.LogoutResp
}

var LogoutResult_Success_DEFAULT *user.LogoutResp

func (p *LogoutResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.LogoutResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *LogoutResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *LogoutResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *LogoutResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *LogoutResult) Unmarshal(in []byte) error {
	msg := new(user.LogoutResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *LogoutResult) GetSuccess() *user.LogoutResp {
	if !p.IsSetSuccess() {
		return LogoutResult_Success_DEFAULT
	}
	return p.Success
}

func (p *LogoutResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.LogoutResp)
}

func (p *LogoutResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *LogoutResult) GetResult() interface{} {
	return p.Success
}

func deleteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.DeleteReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).Delete(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeleteArgs:
		success, err := handler.(user.UserService).Delete(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeleteArgs() interface{} {
	return &DeleteArgs{}
}

func newDeleteResult() interface{} {
	return &DeleteResult{}
}

type DeleteArgs struct {
	Req *user.DeleteReq
}

func (p *DeleteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.DeleteReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteArgs) Unmarshal(in []byte) error {
	msg := new(user.DeleteReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteArgs_Req_DEFAULT *user.DeleteReq

func (p *DeleteArgs) GetReq() *user.DeleteReq {
	if !p.IsSetReq() {
		return DeleteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteResult struct {
	Success *user.DeleteResp
}

var DeleteResult_Success_DEFAULT *user.DeleteResp

func (p *DeleteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.DeleteResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteResult) Unmarshal(in []byte) error {
	msg := new(user.DeleteResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteResult) GetSuccess() *user.DeleteResp {
	if !p.IsSetSuccess() {
		return DeleteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.DeleteResp)
}

func (p *DeleteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteResult) GetResult() interface{} {
	return p.Success
}

func updateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.UpdateReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).Update(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateArgs:
		success, err := handler.(user.UserService).Update(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateArgs() interface{} {
	return &UpdateArgs{}
}

func newUpdateResult() interface{} {
	return &UpdateResult{}
}

type UpdateArgs struct {
	Req *user.UpdateReq
}

func (p *UpdateArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.UpdateReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateArgs) Unmarshal(in []byte) error {
	msg := new(user.UpdateReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateArgs_Req_DEFAULT *user.UpdateReq

func (p *UpdateArgs) GetReq() *user.UpdateReq {
	if !p.IsSetReq() {
		return UpdateArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateResult struct {
	Success *user.UpdateResp
}

var UpdateResult_Success_DEFAULT *user.UpdateResp

func (p *UpdateResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.UpdateResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateResult) Unmarshal(in []byte) error {
	msg := new(user.UpdateResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateResult) GetSuccess() *user.UpdateResp {
	if !p.IsSetSuccess() {
		return UpdateResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.UpdateResp)
}

func (p *UpdateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateResult) GetResult() interface{} {
	return p.Success
}

func getUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.GetUserInfoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).GetUserInfo(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetUserInfoArgs:
		success, err := handler.(user.UserService).GetUserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserInfoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetUserInfoArgs() interface{} {
	return &GetUserInfoArgs{}
}

func newGetUserInfoResult() interface{} {
	return &GetUserInfoResult{}
}

type GetUserInfoArgs struct {
	Req *user.GetUserInfoReq
}

func (p *GetUserInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.GetUserInfoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserInfoArgs) Unmarshal(in []byte) error {
	msg := new(user.GetUserInfoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserInfoArgs_Req_DEFAULT *user.GetUserInfoReq

func (p *GetUserInfoArgs) GetReq() *user.GetUserInfoReq {
	if !p.IsSetReq() {
		return GetUserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetUserInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetUserInfoResult struct {
	Success *user.GetUserInfoResp
}

var GetUserInfoResult_Success_DEFAULT *user.GetUserInfoResp

func (p *GetUserInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.GetUserInfoResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserInfoResult) Unmarshal(in []byte) error {
	msg := new(user.GetUserInfoResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserInfoResult) GetSuccess() *user.GetUserInfoResp {
	if !p.IsSetSuccess() {
		return GetUserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.GetUserInfoResp)
}

func (p *GetUserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserInfoResult) GetResult() interface{} {
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

func (p *kClient) Register(ctx context.Context, Req *user.RegisterReq) (r *user.RegisterResp, err error) {
	var _args RegisterArgs
	_args.Req = Req
	var _result RegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, Req *user.LoginReq) (r *user.LoginResp, err error) {
	var _args LoginArgs
	_args.Req = Req
	var _result LoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Logout(ctx context.Context, Req *user.LogoutReq) (r *user.LogoutResp, err error) {
	var _args LogoutArgs
	_args.Req = Req
	var _result LogoutResult
	if err = p.c.Call(ctx, "Logout", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Delete(ctx context.Context, Req *user.DeleteReq) (r *user.DeleteResp, err error) {
	var _args DeleteArgs
	_args.Req = Req
	var _result DeleteResult
	if err = p.c.Call(ctx, "Delete", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Update(ctx context.Context, Req *user.UpdateReq) (r *user.UpdateResp, err error) {
	var _args UpdateArgs
	_args.Req = Req
	var _result UpdateResult
	if err = p.c.Call(ctx, "Update", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfo(ctx context.Context, Req *user.GetUserInfoReq) (r *user.GetUserInfoResp, err error) {
	var _args GetUserInfoArgs
	_args.Req = Req
	var _result GetUserInfoResult
	if err = p.c.Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
