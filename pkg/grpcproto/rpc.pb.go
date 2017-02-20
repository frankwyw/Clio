// Code generated by protoc-gen-go.
// source: rpc.proto
// DO NOT EDIT!

/*
Package grpcproto is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	Userid
	RepTime
	RespTime
	K8SReq
	RegLoginReq
	RegResp
	RegGetImageTagRep
	RegImage
	RegImageWithId
	PromeReq
	Resp
	UseMrRep
	RespwithId
	UserInfo
*/
package grpcproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Userid struct {
	Userid int32 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
}

func (m *Userid) Reset()                    { *m = Userid{} }
func (m *Userid) String() string            { return proto.CompactTextString(m) }
func (*Userid) ProtoMessage()               {}
func (*Userid) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Userid) GetUserid() int32 {
	if m != nil {
		return m.Userid
	}
	return 0
}

type RepTime struct {
	//
	// resource update_time
	// 0:  Pod
	// 1:  PodTemplate
	// 2:  ReplicationController
	// 3:  Service
	// 4:  EndPoint
	// 5:  Node
	// 6:  Binding
	// 7:  Event
	// 8:  LimitRange
	// 9:  ResourceQuota
	// 10: NameSpace
	// 11: Secret
	// 12: ServiceAccount
	// 13: PersistentVolume
	// 14: PersistentVolumeClaim
	// 15: DeleteOptions
	// 16: ComponentStatus
	// 17: ConfigMap
	ResTypeId int32 `protobuf:"varint,1,opt,name=ResTypeId" json:"ResTypeId,omitempty"`
	Userid    int32 `protobuf:"varint,2,opt,name=userid" json:"userid,omitempty"`
}

func (m *RepTime) Reset()                    { *m = RepTime{} }
func (m *RepTime) String() string            { return proto.CompactTextString(m) }
func (*RepTime) ProtoMessage()               {}
func (*RepTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RepTime) GetResTypeId() int32 {
	if m != nil {
		return m.ResTypeId
	}
	return 0
}

func (m *RepTime) GetUserid() int32 {
	if m != nil {
		return m.Userid
	}
	return 0
}

type RespTime struct {
	Time string `protobuf:"bytes,1,opt,name=time" json:"time,omitempty"`
}

func (m *RespTime) Reset()                    { *m = RespTime{} }
func (m *RespTime) String() string            { return proto.CompactTextString(m) }
func (*RespTime) ProtoMessage()               {}
func (*RespTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RespTime) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

type K8SReq struct {
	// restime update
	Resint int32 `protobuf:"varint,1,opt,name=resint" json:"resint,omitempty"`
	// get, post, put, delete
	Reqtype string `protobuf:"bytes,2,opt,name=reqtype" json:"reqtype,omitempty"`
	// example api/v1/namespaces?watch=true&timeoutSeconds=3
	// example api/v1/namespaces, (basicpath + pathParam + quereparam)
	Url string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	// json in string
	Bodyparam string `protobuf:"bytes,4,opt,name=bodyparam" json:"bodyparam,omitempty"`
	Userid    int32  `protobuf:"varint,5,opt,name=userid" json:"userid,omitempty"`
}

func (m *K8SReq) Reset()                    { *m = K8SReq{} }
func (m *K8SReq) String() string            { return proto.CompactTextString(m) }
func (*K8SReq) ProtoMessage()               {}
func (*K8SReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *K8SReq) GetResint() int32 {
	if m != nil {
		return m.Resint
	}
	return 0
}

func (m *K8SReq) GetReqtype() string {
	if m != nil {
		return m.Reqtype
	}
	return ""
}

func (m *K8SReq) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *K8SReq) GetBodyparam() string {
	if m != nil {
		return m.Bodyparam
	}
	return ""
}

func (m *K8SReq) GetUserid() int32 {
	if m != nil {
		return m.Userid
	}
	return 0
}

type RegLoginReq struct {
	// not only localregistry but also remoteregistry default localregistry
	Url      string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Secure   bool   `protobuf:"varint,4,opt,name=secure" json:"secure,omitempty"`
	Userid   int32  `protobuf:"varint,5,opt,name=userid" json:"userid,omitempty"`
}

func (m *RegLoginReq) Reset()                    { *m = RegLoginReq{} }
func (m *RegLoginReq) String() string            { return proto.CompactTextString(m) }
func (*RegLoginReq) ProtoMessage()               {}
func (*RegLoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RegLoginReq) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *RegLoginReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegLoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegLoginReq) GetSecure() bool {
	if m != nil {
		return m.Secure
	}
	return false
}

func (m *RegLoginReq) GetUserid() int32 {
	if m != nil {
		return m.Userid
	}
	return 0
}

type RegResp struct {
	// http code
	Respcode string `protobuf:"bytes,1,opt,name=respcode" json:"respcode,omitempty"`
	// loginresp, listreporesp getimagetagresp regpushresp DelImage
	Resp   string `protobuf:"bytes,2,opt,name=resp" json:"resp,omitempty"`
	Userid int32  `protobuf:"varint,3,opt,name=userid" json:"userid,omitempty"`
}

func (m *RegResp) Reset()                    { *m = RegResp{} }
func (m *RegResp) String() string            { return proto.CompactTextString(m) }
func (*RegResp) ProtoMessage()               {}
func (*RegResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RegResp) GetRespcode() string {
	if m != nil {
		return m.Respcode
	}
	return ""
}

func (m *RegResp) GetResp() string {
	if m != nil {
		return m.Resp
	}
	return ""
}

func (m *RegResp) GetUserid() int32 {
	if m != nil {
		return m.Userid
	}
	return 0
}

type RegGetImageTagRep struct {
	Images string `protobuf:"bytes,1,opt,name=images" json:"images,omitempty"`
	Userid int32  `protobuf:"varint,2,opt,name=userid" json:"userid,omitempty"`
}

func (m *RegGetImageTagRep) Reset()                    { *m = RegGetImageTagRep{} }
func (m *RegGetImageTagRep) String() string            { return proto.CompactTextString(m) }
func (*RegGetImageTagRep) ProtoMessage()               {}
func (*RegGetImageTagRep) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RegGetImageTagRep) GetImages() string {
	if m != nil {
		return m.Images
	}
	return ""
}

func (m *RegGetImageTagRep) GetUserid() int32 {
	if m != nil {
		return m.Userid
	}
	return 0
}

type RegImage struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// tag or digset
	Reference string `protobuf:"bytes,2,opt,name=reference" json:"reference,omitempty"`
}

func (m *RegImage) Reset()                    { *m = RegImage{} }
func (m *RegImage) String() string            { return proto.CompactTextString(m) }
func (*RegImage) ProtoMessage()               {}
func (*RegImage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RegImage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegImage) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

type RegImageWithId struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Reference string `protobuf:"bytes,2,opt,name=reference" json:"reference,omitempty"`
	Userid    int32  `protobuf:"varint,3,opt,name=userid" json:"userid,omitempty"`
}

func (m *RegImageWithId) Reset()                    { *m = RegImageWithId{} }
func (m *RegImageWithId) String() string            { return proto.CompactTextString(m) }
func (*RegImageWithId) ProtoMessage()               {}
func (*RegImageWithId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RegImageWithId) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegImageWithId) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *RegImageWithId) GetUserid() int32 {
	if m != nil {
		return m.Userid
	}
	return 0
}

type PromeReq struct {
	Is_Range bool   `protobuf:"varint,1,opt,name=is_Range,json=isRange" json:"is_Range,omitempty"`
	Url      string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Userid   int32  `protobuf:"varint,3,opt,name=userid" json:"userid,omitempty"`
}

func (m *PromeReq) Reset()                    { *m = PromeReq{} }
func (m *PromeReq) String() string            { return proto.CompactTextString(m) }
func (*PromeReq) ProtoMessage()               {}
func (*PromeReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *PromeReq) GetIs_Range() bool {
	if m != nil {
		return m.Is_Range
	}
	return false
}

func (m *PromeReq) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *PromeReq) GetUserid() int32 {
	if m != nil {
		return m.Userid
	}
	return 0
}

type Resp struct {
	Httpcode string `protobuf:"bytes,1,opt,name=httpcode" json:"httpcode,omitempty"`
	Resp     string `protobuf:"bytes,2,opt,name=resp" json:"resp,omitempty"`
}

func (m *Resp) Reset()                    { *m = Resp{} }
func (m *Resp) String() string            { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()               {}
func (*Resp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Resp) GetHttpcode() string {
	if m != nil {
		return m.Httpcode
	}
	return ""
}

func (m *Resp) GetResp() string {
	if m != nil {
		return m.Resp
	}
	return ""
}

type UseMrRep struct {
	Reqmethod string `protobuf:"bytes,1,opt,name=reqmethod" json:"reqmethod,omitempty"`
	// pathparam process in frontend
	Requrl    string `protobuf:"bytes,2,opt,name=requrl" json:"requrl,omitempty"`
	Bodyparam string `protobuf:"bytes,3,opt,name=bodyparam" json:"bodyparam,omitempty"`
	Userid    int32  `protobuf:"varint,4,opt,name=userid" json:"userid,omitempty"`
}

func (m *UseMrRep) Reset()                    { *m = UseMrRep{} }
func (m *UseMrRep) String() string            { return proto.CompactTextString(m) }
func (*UseMrRep) ProtoMessage()               {}
func (*UseMrRep) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *UseMrRep) GetReqmethod() string {
	if m != nil {
		return m.Reqmethod
	}
	return ""
}

func (m *UseMrRep) GetRequrl() string {
	if m != nil {
		return m.Requrl
	}
	return ""
}

func (m *UseMrRep) GetBodyparam() string {
	if m != nil {
		return m.Bodyparam
	}
	return ""
}

func (m *UseMrRep) GetUserid() int32 {
	if m != nil {
		return m.Userid
	}
	return 0
}

type RespwithId struct {
	Resp   *Resp `protobuf:"bytes,1,opt,name=resp" json:"resp,omitempty"`
	Userid int32 `protobuf:"varint,2,opt,name=userid" json:"userid,omitempty"`
}

func (m *RespwithId) Reset()                    { *m = RespwithId{} }
func (m *RespwithId) String() string            { return proto.CompactTextString(m) }
func (*RespwithId) ProtoMessage()               {}
func (*RespwithId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *RespwithId) GetResp() *Resp {
	if m != nil {
		return m.Resp
	}
	return nil
}

func (m *RespwithId) GetUserid() int32 {
	if m != nil {
		return m.Userid
	}
	return 0
}

type UserInfo struct {
	User   string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Passwd string `protobuf:"bytes,2,opt,name=passwd" json:"passwd,omitempty"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *UserInfo) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UserInfo) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

func init() {
	proto.RegisterType((*Userid)(nil), "grpcproto.Userid")
	proto.RegisterType((*RepTime)(nil), "grpcproto.RepTime")
	proto.RegisterType((*RespTime)(nil), "grpcproto.RespTime")
	proto.RegisterType((*K8SReq)(nil), "grpcproto.K8sReq")
	proto.RegisterType((*RegLoginReq)(nil), "grpcproto.regLoginReq")
	proto.RegisterType((*RegResp)(nil), "grpcproto.regResp")
	proto.RegisterType((*RegGetImageTagRep)(nil), "grpcproto.RegGetImageTagRep")
	proto.RegisterType((*RegImage)(nil), "grpcproto.RegImage")
	proto.RegisterType((*RegImageWithId)(nil), "grpcproto.RegImageWithId")
	proto.RegisterType((*PromeReq)(nil), "grpcproto.promeReq")
	proto.RegisterType((*Resp)(nil), "grpcproto.Resp")
	proto.RegisterType((*UseMrRep)(nil), "grpcproto.UseMrRep")
	proto.RegisterType((*RespwithId)(nil), "grpcproto.RespwithId")
	proto.RegisterType((*UserInfo)(nil), "grpcproto.UserInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for G service

type GClient interface {
	Login(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*Resp, error)
	GetResTime(ctx context.Context, in *RepTime, opts ...grpc.CallOption) (*RespTime, error)
	K8SRestApi(ctx context.Context, in *K8SReq, opts ...grpc.CallOption) (*Resp, error)
	K8SStreamApi(ctx context.Context, in *K8SReq, opts ...grpc.CallOption) (G_K8SStreamApiClient, error)
	// operate after login
	RegLogin(ctx context.Context, in *RegLoginReq, opts ...grpc.CallOption) (*Resp, error)
	// pull, push only the server where backend is
	RegListRepo(ctx context.Context, in *Userid, opts ...grpc.CallOption) (*Resp, error)
	RegGetImageTag(ctx context.Context, in *RegGetImageTagRep, opts ...grpc.CallOption) (*Resp, error)
	RegPush(ctx context.Context, in *RegImageWithId, opts ...grpc.CallOption) (*Resp, error)
	RegDelImage(ctx context.Context, in *RegImageWithId, opts ...grpc.CallOption) (*Resp, error)
	RegPull(ctx context.Context, in *RegImageWithId, opts ...grpc.CallOption) (*Resp, error)
	// queryrange or query
	PromQuery(ctx context.Context, in *PromeReq, opts ...grpc.CallOption) (*Resp, error)
	UserMApi(ctx context.Context, in *UseMrRep, opts ...grpc.CallOption) (*Resp, error)
}

type gClient struct {
	cc *grpc.ClientConn
}

func NewGClient(cc *grpc.ClientConn) GClient {
	return &gClient{cc}
}

func (c *gClient) Login(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/grpcproto.G/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gClient) GetResTime(ctx context.Context, in *RepTime, opts ...grpc.CallOption) (*RespTime, error) {
	out := new(RespTime)
	err := grpc.Invoke(ctx, "/grpcproto.G/GetResTime", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gClient) K8SRestApi(ctx context.Context, in *K8SReq, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/grpcproto.G/K8sRestApi", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gClient) K8SStreamApi(ctx context.Context, in *K8SReq, opts ...grpc.CallOption) (G_K8SStreamApiClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_G_serviceDesc.Streams[0], c.cc, "/grpcproto.G/K8sStreamApi", opts...)
	if err != nil {
		return nil, err
	}
	x := &gK8SStreamApiClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type G_K8SStreamApiClient interface {
	Recv() (*Resp, error)
	grpc.ClientStream
}

type gK8SStreamApiClient struct {
	grpc.ClientStream
}

func (x *gK8SStreamApiClient) Recv() (*Resp, error) {
	m := new(Resp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gClient) RegLogin(ctx context.Context, in *RegLoginReq, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/grpcproto.G/RegLogin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gClient) RegListRepo(ctx context.Context, in *Userid, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/grpcproto.G/RegListRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gClient) RegGetImageTag(ctx context.Context, in *RegGetImageTagRep, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/grpcproto.G/RegGetImageTag", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gClient) RegPush(ctx context.Context, in *RegImageWithId, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/grpcproto.G/RegPush", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gClient) RegDelImage(ctx context.Context, in *RegImageWithId, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/grpcproto.G/RegDelImage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gClient) RegPull(ctx context.Context, in *RegImageWithId, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/grpcproto.G/RegPull", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gClient) PromQuery(ctx context.Context, in *PromeReq, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/grpcproto.G/PromQuery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gClient) UserMApi(ctx context.Context, in *UseMrRep, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/grpcproto.G/UserMApi", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for G service

type GServer interface {
	Login(context.Context, *UserInfo) (*Resp, error)
	GetResTime(context.Context, *RepTime) (*RespTime, error)
	K8SRestApi(context.Context, *K8SReq) (*Resp, error)
	K8SStreamApi(*K8SReq, G_K8SStreamApiServer) error
	// operate after login
	RegLogin(context.Context, *RegLoginReq) (*Resp, error)
	// pull, push only the server where backend is
	RegListRepo(context.Context, *Userid) (*Resp, error)
	RegGetImageTag(context.Context, *RegGetImageTagRep) (*Resp, error)
	RegPush(context.Context, *RegImageWithId) (*Resp, error)
	RegDelImage(context.Context, *RegImageWithId) (*Resp, error)
	RegPull(context.Context, *RegImageWithId) (*Resp, error)
	// queryrange or query
	PromQuery(context.Context, *PromeReq) (*Resp, error)
	UserMApi(context.Context, *UseMrRep) (*Resp, error)
}

func RegisterGServer(s *grpc.Server, srv GServer) {
	s.RegisterService(&_G_serviceDesc, srv)
}

func _G_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.G/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GServer).Login(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _G_GetResTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepTime)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GServer).GetResTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.G/GetResTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GServer).GetResTime(ctx, req.(*RepTime))
	}
	return interceptor(ctx, in, info, handler)
}

func _G_K8SRestApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(K8SReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GServer).K8SRestApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.G/K8SRestApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GServer).K8SRestApi(ctx, req.(*K8SReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _G_K8SStreamApi_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(K8SReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GServer).K8SStreamApi(m, &gK8SStreamApiServer{stream})
}

type G_K8SStreamApiServer interface {
	Send(*Resp) error
	grpc.ServerStream
}

type gK8SStreamApiServer struct {
	grpc.ServerStream
}

func (x *gK8SStreamApiServer) Send(m *Resp) error {
	return x.ServerStream.SendMsg(m)
}

func _G_RegLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GServer).RegLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.G/RegLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GServer).RegLogin(ctx, req.(*RegLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _G_RegListRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Userid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GServer).RegListRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.G/RegListRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GServer).RegListRepo(ctx, req.(*Userid))
	}
	return interceptor(ctx, in, info, handler)
}

func _G_RegGetImageTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegGetImageTagRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GServer).RegGetImageTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.G/RegGetImageTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GServer).RegGetImageTag(ctx, req.(*RegGetImageTagRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _G_RegPush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegImageWithId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GServer).RegPush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.G/RegPush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GServer).RegPush(ctx, req.(*RegImageWithId))
	}
	return interceptor(ctx, in, info, handler)
}

func _G_RegDelImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegImageWithId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GServer).RegDelImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.G/RegDelImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GServer).RegDelImage(ctx, req.(*RegImageWithId))
	}
	return interceptor(ctx, in, info, handler)
}

func _G_RegPull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegImageWithId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GServer).RegPull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.G/RegPull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GServer).RegPull(ctx, req.(*RegImageWithId))
	}
	return interceptor(ctx, in, info, handler)
}

func _G_PromQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GServer).PromQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.G/PromQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GServer).PromQuery(ctx, req.(*PromeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _G_UserMApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UseMrRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GServer).UserMApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.G/UserMApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GServer).UserMApi(ctx, req.(*UseMrRep))
	}
	return interceptor(ctx, in, info, handler)
}

var _G_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcproto.G",
	HandlerType: (*GServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _G_Login_Handler,
		},
		{
			MethodName: "GetResTime",
			Handler:    _G_GetResTime_Handler,
		},
		{
			MethodName: "K8sRestApi",
			Handler:    _G_K8SRestApi_Handler,
		},
		{
			MethodName: "RegLogin",
			Handler:    _G_RegLogin_Handler,
		},
		{
			MethodName: "RegListRepo",
			Handler:    _G_RegListRepo_Handler,
		},
		{
			MethodName: "RegGetImageTag",
			Handler:    _G_RegGetImageTag_Handler,
		},
		{
			MethodName: "RegPush",
			Handler:    _G_RegPush_Handler,
		},
		{
			MethodName: "RegDelImage",
			Handler:    _G_RegDelImage_Handler,
		},
		{
			MethodName: "RegPull",
			Handler:    _G_RegPull_Handler,
		},
		{
			MethodName: "PromQuery",
			Handler:    _G_PromQuery_Handler,
		},
		{
			MethodName: "UserMApi",
			Handler:    _G_UserMApi_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "K8sStreamApi",
			Handler:       _G_K8SStreamApi_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 660 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6b, 0xdb, 0x4c,
	0x10, 0x8e, 0xe2, 0x2f, 0x69, 0xf2, 0xf2, 0xb6, 0xd9, 0x40, 0x70, 0x42, 0x28, 0x41, 0xbd, 0xf4,
	0x94, 0x16, 0x87, 0xa6, 0x81, 0x16, 0x4a, 0x68, 0x21, 0x98, 0x34, 0x34, 0xd9, 0x26, 0x14, 0x7a,
	0x29, 0x8a, 0x35, 0x91, 0x05, 0x96, 0x25, 0xef, 0xae, 0x1b, 0x7c, 0xed, 0xa5, 0x7f, 0xa3, 0x3f,
	0xb5, 0xcc, 0x48, 0x6b, 0x49, 0xae, 0x0c, 0xcd, 0x6d, 0x67, 0x76, 0x9e, 0x67, 0x9e, 0x9d, 0x8f,
	0x05, 0x4f, 0x65, 0xa3, 0xa3, 0x4c, 0xa5, 0x26, 0x15, 0x5e, 0xa4, 0xb2, 0x11, 0x1f, 0xfd, 0x43,
	0xe8, 0xde, 0x6a, 0x54, 0x71, 0x28, 0x76, 0xa1, 0x3b, 0xe7, 0x53, 0xdf, 0x39, 0x74, 0x5e, 0x74,
	0x64, 0x61, 0xf9, 0xef, 0xa1, 0x27, 0x31, 0xbb, 0x89, 0x13, 0x14, 0x07, 0xe0, 0x49, 0xd4, 0x37,
	0x8b, 0x0c, 0x87, 0x36, 0xaa, 0x74, 0x54, 0x08, 0x36, 0x6b, 0x04, 0xcf, 0xc0, 0x95, 0xa8, 0x73,
	0x06, 0x01, 0x6d, 0x13, 0x27, 0xc8, 0x60, 0x4f, 0xf2, 0xd9, 0xff, 0xe9, 0x40, 0xf7, 0xe2, 0x54,
	0x4b, 0x9c, 0x11, 0x85, 0x42, 0x1d, 0x4f, 0x8d, 0xd5, 0x90, 0x5b, 0xa2, 0x0f, 0x3d, 0x85, 0x33,
	0xb3, 0xc8, 0x90, 0xb9, 0x3d, 0x69, 0x4d, 0xf1, 0x14, 0x5a, 0x73, 0x35, 0xe9, 0xb7, 0xd8, 0x4b,
	0x47, 0x12, 0x79, 0x97, 0x86, 0x8b, 0x2c, 0x50, 0x41, 0xd2, 0x6f, 0xb3, 0xbf, 0x74, 0x54, 0x44,
	0x76, 0x6a, 0x22, 0x7f, 0x39, 0xb0, 0xa5, 0x30, 0xfa, 0x94, 0x46, 0xf1, 0x94, 0x94, 0x14, 0xbc,
	0x4e, 0xc9, 0xbb, 0x0f, 0x2e, 0xc5, 0x4e, 0x83, 0xc4, 0x8a, 0x58, 0xda, 0x74, 0x97, 0x05, 0x5a,
	0x3f, 0xa4, 0x2a, 0x2c, 0xa4, 0x2c, 0x6d, 0xca, 0xa8, 0x71, 0x34, 0x57, 0xc8, 0x62, 0x5c, 0x59,
	0x58, 0x6b, 0x95, 0x5c, 0xd3, 0x5b, 0x23, 0xaa, 0x18, 0xd1, 0x2a, 0xd4, 0xd9, 0x28, 0x0d, 0x6d,
	0xc5, 0x96, 0x36, 0x55, 0x92, 0xce, 0x85, 0x14, 0x3e, 0x57, 0x28, 0x5b, 0x35, 0xca, 0x0f, 0xb0,
	0x2d, 0x31, 0x3a, 0x47, 0x33, 0x4c, 0x82, 0x08, 0x6f, 0x82, 0x48, 0x22, 0x07, 0xc7, 0x64, 0xea,
	0x82, 0xba, 0xb0, 0xd6, 0xb6, 0xf1, 0x1d, 0xb5, 0x31, 0x62, 0x06, 0x4a, 0xce, 0x75, 0x28, 0xda,
	0xc8, 0x35, 0x38, 0x00, 0x4f, 0xe1, 0x3d, 0x2a, 0x9c, 0x8e, 0x6c, 0x81, 0x4a, 0x87, 0xff, 0x0d,
	0xfe, 0xb7, 0xe8, 0xaf, 0xb1, 0x19, 0x0f, 0xc3, 0xc7, 0x73, 0xac, 0x7d, 0xde, 0x67, 0x70, 0x33,
	0x95, 0x26, 0x48, 0x7d, 0xdb, 0x03, 0x37, 0xd6, 0xdf, 0x65, 0x30, 0x8d, 0x72, 0x66, 0x57, 0xf6,
	0x62, 0xcd, 0xa6, 0x6d, 0xe9, 0x66, 0xd9, 0xd2, 0x75, 0x84, 0x27, 0xd0, 0xb6, 0xf5, 0x1f, 0x1b,
	0x53, 0xab, 0xbf, 0xb5, 0x9b, 0xea, 0xef, 0xff, 0x00, 0xf7, 0x56, 0xe3, 0xa5, 0xa2, 0xf2, 0xf2,
	0x53, 0x66, 0x09, 0x9a, 0x71, 0x1a, 0x16, 0xe0, 0xd2, 0x91, 0x0f, 0xfa, 0xac, 0x94, 0x53, 0x58,
	0xf5, 0xe1, 0x6d, 0xad, 0x1f, 0xde, 0x76, 0x4d, 0xef, 0x10, 0x80, 0xf4, 0x3e, 0xe4, 0x85, 0x7d,
	0x5e, 0x28, 0xa3, 0xa4, 0x5b, 0x83, 0x27, 0x47, 0xcb, 0x65, 0x3f, 0xa2, 0xa0, 0xbf, 0x46, 0x65,
	0x73, 0xe5, 0xe9, 0xf4, 0x04, 0x35, 0x9c, 0xde, 0xa7, 0xf4, 0x44, 0xf2, 0xda, 0x0e, 0xd1, 0x99,
	0x70, 0x3c, 0xd9, 0xa1, 0x15, 0x9e, 0x5b, 0x83, 0xdf, 0x1d, 0x70, 0xce, 0xc5, 0x4b, 0xe8, 0xf0,
	0x06, 0x89, 0x9d, 0x4a, 0x56, 0xcb, 0xb7, 0xbf, 0x2a, 0xc5, 0xdf, 0x10, 0x6f, 0x00, 0xce, 0xd1,
	0xd0, 0x1f, 0xc2, 0xbf, 0x43, 0x2d, 0x80, 0x7f, 0x8c, 0xfd, 0x9d, 0x15, 0x10, 0x39, 0xfd, 0x0d,
	0x31, 0x00, 0xe0, 0x3f, 0x43, 0x9b, 0xb3, 0x2c, 0x16, 0xdb, 0x95, 0xa0, 0xfc, 0x2b, 0x69, 0x4a,
	0x76, 0x02, 0xff, 0x5d, 0x9c, 0xea, 0x2f, 0x46, 0x61, 0x90, 0xfc, 0x33, 0xea, 0x95, 0x23, 0x5e,
	0xf3, 0xe4, 0xe7, 0x0f, 0xdb, 0xad, 0x04, 0x54, 0xfe, 0x8b, 0xa6, 0x74, 0xc7, 0xb0, 0x45, 0xb0,
	0x58, 0x1b, 0x89, 0x59, 0x5a, 0xcb, 0x96, 0x7f, 0xb9, 0x4d, 0xa0, 0x33, 0xde, 0x93, 0xca, 0xaa,
	0x8a, 0x83, 0x5a, 0xd0, 0xca, 0x16, 0x37, 0xd7, 0xb4, 0x27, 0x31, 0xba, 0x9a, 0xeb, 0xb1, 0xd8,
	0xab, 0x63, 0x2b, 0xeb, 0xd7, 0x04, 0x7c, 0xcb, 0x82, 0x3f, 0xe2, 0x24, 0x5f, 0xf2, 0xc7, 0x81,
	0x6d, 0xd6, 0xc9, 0xe4, 0x91, 0xc0, 0x63, 0xf0, 0xae, 0x54, 0x9a, 0x5c, 0xcf, 0x51, 0x2d, 0x6a,
	0x73, 0x63, 0x77, 0xba, 0x09, 0x34, 0xc8, 0xc7, 0xf4, 0x92, 0xda, 0xb8, 0x32, 0x6b, 0xbc, 0x7e,
	0x0d, 0x98, 0xbb, 0x2e, 0x5b, 0xc7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x76, 0x29, 0xd2, 0x31,
	0x09, 0x07, 0x00, 0x00,
}
