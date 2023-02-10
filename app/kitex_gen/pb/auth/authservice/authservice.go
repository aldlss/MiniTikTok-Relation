// Code generated by Kitex v0.4.4. DO NOT EDIT.

package authservice

import (
	"context"
	"fmt"
	auth "github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/auth"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return authServiceServiceInfo
}

var authServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "AuthService"
	handlerType := (*auth.AuthService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Auth":         kitex.NewMethodInfo(authHandler, newAuthArgs, newAuthResult, false),
		"RetriveToken": kitex.NewMethodInfo(retriveTokenHandler, newRetriveTokenArgs, newRetriveTokenResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "mini_tiktok.proto.auth",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func authHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.AuthRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).Auth(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *AuthArgs:
		success, err := handler.(auth.AuthService).Auth(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AuthResult)
		realResult.Success = success
	}
	return nil
}
func newAuthArgs() interface{} {
	return &AuthArgs{}
}

func newAuthResult() interface{} {
	return &AuthResult{}
}

type AuthArgs struct {
	Req *auth.AuthRequest
}

func (p *AuthArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.AuthRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AuthArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AuthArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AuthArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in AuthArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *AuthArgs) Unmarshal(in []byte) error {
	msg := new(auth.AuthRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AuthArgs_Req_DEFAULT *auth.AuthRequest

func (p *AuthArgs) GetReq() *auth.AuthRequest {
	if !p.IsSetReq() {
		return AuthArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AuthArgs) IsSetReq() bool {
	return p.Req != nil
}

type AuthResult struct {
	Success *auth.AuthResponse
}

var AuthResult_Success_DEFAULT *auth.AuthResponse

func (p *AuthResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.AuthResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AuthResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AuthResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AuthResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in AuthResult")
	}
	return proto.Marshal(p.Success)
}

func (p *AuthResult) Unmarshal(in []byte) error {
	msg := new(auth.AuthResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AuthResult) GetSuccess() *auth.AuthResponse {
	if !p.IsSetSuccess() {
		return AuthResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AuthResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.AuthResponse)
}

func (p *AuthResult) IsSetSuccess() bool {
	return p.Success != nil
}

func retriveTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.TokenRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).RetriveToken(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *RetriveTokenArgs:
		success, err := handler.(auth.AuthService).RetriveToken(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RetriveTokenResult)
		realResult.Success = success
	}
	return nil
}
func newRetriveTokenArgs() interface{} {
	return &RetriveTokenArgs{}
}

func newRetriveTokenResult() interface{} {
	return &RetriveTokenResult{}
}

type RetriveTokenArgs struct {
	Req *auth.TokenRequest
}

func (p *RetriveTokenArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.TokenRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RetriveTokenArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RetriveTokenArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RetriveTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in RetriveTokenArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *RetriveTokenArgs) Unmarshal(in []byte) error {
	msg := new(auth.TokenRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RetriveTokenArgs_Req_DEFAULT *auth.TokenRequest

func (p *RetriveTokenArgs) GetReq() *auth.TokenRequest {
	if !p.IsSetReq() {
		return RetriveTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RetriveTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

type RetriveTokenResult struct {
	Success *auth.TokenResponse
}

var RetriveTokenResult_Success_DEFAULT *auth.TokenResponse

func (p *RetriveTokenResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.TokenResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RetriveTokenResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RetriveTokenResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RetriveTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in RetriveTokenResult")
	}
	return proto.Marshal(p.Success)
}

func (p *RetriveTokenResult) Unmarshal(in []byte) error {
	msg := new(auth.TokenResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RetriveTokenResult) GetSuccess() *auth.TokenResponse {
	if !p.IsSetSuccess() {
		return RetriveTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RetriveTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.TokenResponse)
}

func (p *RetriveTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Auth(ctx context.Context, Req *auth.AuthRequest) (r *auth.AuthResponse, err error) {
	var _args AuthArgs
	_args.Req = Req
	var _result AuthResult
	if err = p.c.Call(ctx, "Auth", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RetriveToken(ctx context.Context, Req *auth.TokenRequest) (r *auth.TokenResponse, err error) {
	var _args RetriveTokenArgs
	_args.Req = Req
	var _result RetriveTokenResult
	if err = p.c.Call(ctx, "RetriveToken", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
