// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/person/person.proto

package person

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	response "github.com/wulungtry/techtalk/api/proto/response"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AddPersonModel struct {
	PersonalId           string               `protobuf:"bytes,1,opt,name=PersonalId,proto3" json:"PersonalId,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	CityOfBirth          string               `protobuf:"bytes,3,opt,name=CityOfBirth,proto3" json:"CityOfBirth,omitempty"`
	DateOfBirth          *timestamp.Timestamp `protobuf:"bytes,4,opt,name=DateOfBirth,proto3" json:"DateOfBirth,omitempty"`
	Height               int32                `protobuf:"varint,5,opt,name=Height,proto3" json:"Height,omitempty"`
	Weight               int32                `protobuf:"varint,6,opt,name=Weight,proto3" json:"Weight,omitempty"`
	IsMarried            bool                 `protobuf:"varint,7,opt,name=IsMarried,proto3" json:"IsMarried,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AddPersonModel) Reset()         { *m = AddPersonModel{} }
func (m *AddPersonModel) String() string { return proto.CompactTextString(m) }
func (*AddPersonModel) ProtoMessage()    {}
func (*AddPersonModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8d05f5e6fca1807, []int{0}
}

func (m *AddPersonModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPersonModel.Unmarshal(m, b)
}
func (m *AddPersonModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPersonModel.Marshal(b, m, deterministic)
}
func (m *AddPersonModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPersonModel.Merge(m, src)
}
func (m *AddPersonModel) XXX_Size() int {
	return xxx_messageInfo_AddPersonModel.Size(m)
}
func (m *AddPersonModel) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPersonModel.DiscardUnknown(m)
}

var xxx_messageInfo_AddPersonModel proto.InternalMessageInfo

func (m *AddPersonModel) GetPersonalId() string {
	if m != nil {
		return m.PersonalId
	}
	return ""
}

func (m *AddPersonModel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddPersonModel) GetCityOfBirth() string {
	if m != nil {
		return m.CityOfBirth
	}
	return ""
}

func (m *AddPersonModel) GetDateOfBirth() *timestamp.Timestamp {
	if m != nil {
		return m.DateOfBirth
	}
	return nil
}

func (m *AddPersonModel) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *AddPersonModel) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *AddPersonModel) GetIsMarried() bool {
	if m != nil {
		return m.IsMarried
	}
	return false
}

func init() {
	proto.RegisterType((*AddPersonModel)(nil), "personal.AddPersonModel")
}

func init() { proto.RegisterFile("proto/person/person.proto", fileDescriptor_f8d05f5e6fca1807) }

var fileDescriptor_f8d05f5e6fca1807 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x51, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0x26, 0xbd, 0x6d, 0x6e, 0x3b, 0x85, 0x7b, 0x65, 0xc0, 0x12, 0x63, 0xd1, 0xd0, 0x55, 0x70,
	0x31, 0x81, 0xba, 0x13, 0x41, 0xac, 0x2e, 0xec, 0xa2, 0x5a, 0xa2, 0x20, 0x74, 0x37, 0x6d, 0x4e,
	0x93, 0xc1, 0x24, 0x13, 0x66, 0x26, 0x4a, 0xb7, 0xbe, 0x82, 0x8f, 0xe6, 0x2b, 0xf8, 0x04, 0x3e,
	0x81, 0x38, 0x33, 0xa9, 0x75, 0x95, 0xef, 0x2f, 0x87, 0xef, 0x9c, 0x41, 0x07, 0x95, 0xe0, 0x8a,
	0x47, 0x15, 0x08, 0xc9, 0x4b, 0xfb, 0x21, 0x5a, 0xc3, 0x5d, 0xc3, 0x68, 0xee, 0x5f, 0xa4, 0x4c,
	0x65, 0xf5, 0x92, 0xac, 0x78, 0x11, 0xbd, 0xd4, 0x79, 0x5d, 0xa6, 0x4a, 0x6c, 0x22, 0x05, 0xab,
	0x4c, 0xd1, 0xfc, 0x29, 0xa2, 0x15, 0x8b, 0xcc, 0x18, 0x01, 0xb2, 0xe2, 0xa5, 0x84, 0x2d, 0x30,
	0xa3, 0xfc, 0x61, 0xca, 0x79, 0x9a, 0x83, 0x4e, 0xd2, 0xb2, 0xe4, 0x8a, 0x2a, 0xc6, 0x4b, 0x69,
	0xdd, 0x63, 0xeb, 0x6a, 0xb6, 0xac, 0xd7, 0x91, 0x62, 0x05, 0x48, 0x45, 0x8b, 0xca, 0x04, 0x46,
	0x9f, 0x0e, 0xfa, 0x77, 0x99, 0x24, 0x73, 0xdd, 0x67, 0xc6, 0x13, 0xc8, 0xf1, 0x11, 0x42, 0x73,
	0x5b, 0x6f, 0x9a, 0x78, 0x4e, 0xe0, 0x84, 0xbd, 0x78, 0x47, 0xc1, 0x18, 0xb5, 0x6f, 0x69, 0x01,
	0x5e, 0x4b, 0x3b, 0x1a, 0xe3, 0x00, 0xf5, 0xaf, 0x98, 0xda, 0xdc, 0xad, 0x27, 0x4c, 0xa8, 0xcc,
	0xfb, 0xa3, 0xad, 0x5d, 0x09, 0x9f, 0xa3, 0xfe, 0x35, 0x55, 0xd0, 0x24, 0xda, 0x81, 0x13, 0xf6,
	0xc7, 0x3e, 0x31, 0xfd, 0x48, 0xd3, 0x8f, 0x3c, 0x34, 0xfd, 0xe2, 0xdd, 0x38, 0x1e, 0x20, 0xf7,
	0x06, 0x58, 0x9a, 0x29, 0xaf, 0x13, 0x38, 0x61, 0x27, 0xb6, 0xec, 0x5b, 0x7f, 0x34, 0xba, 0x6b,
	0x74, 0xc3, 0xf0, 0x10, 0xf5, 0xa6, 0x72, 0x46, 0x85, 0x60, 0x90, 0x78, 0x7f, 0x03, 0x27, 0xec,
	0xc6, 0x3f, 0xc2, 0xb8, 0x40, 0xff, 0x9b, 0x7d, 0xee, 0x41, 0x3c, 0xb3, 0x15, 0xe0, 0x05, 0xea,
	0x6d, 0xcf, 0x80, 0x3d, 0xd2, 0xbc, 0x0f, 0xf9, 0x7d, 0x1b, 0x7f, 0x40, 0xb6, 0xe7, 0x8f, 0x2d,
	0x98, 0x50, 0x09, 0xa3, 0xc3, 0xd7, 0xf7, 0x8f, 0xb7, 0xd6, 0xfe, 0x68, 0x2f, 0x6a, 0xfe, 0xb4,
	0xe0, 0xcc, 0x39, 0x99, 0x74, 0x17, 0xae, 0x21, 0x4b, 0x57, 0xef, 0x79, 0xfa, 0x15, 0x00, 0x00,
	0xff, 0xff, 0xf4, 0x2e, 0xd7, 0x01, 0x1b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PersonalServiceClient is the client API for PersonalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PersonalServiceClient interface {
	AddPerson(ctx context.Context, in *AddPersonModel, opts ...grpc.CallOption) (*response.ResponseBase, error)
}

type personalServiceClient struct {
	cc *grpc.ClientConn
}

func NewPersonalServiceClient(cc *grpc.ClientConn) PersonalServiceClient {
	return &personalServiceClient{cc}
}

func (c *personalServiceClient) AddPerson(ctx context.Context, in *AddPersonModel, opts ...grpc.CallOption) (*response.ResponseBase, error) {
	out := new(response.ResponseBase)
	err := c.cc.Invoke(ctx, "/personal.PersonalService/AddPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonalServiceServer is the server API for PersonalService service.
type PersonalServiceServer interface {
	AddPerson(context.Context, *AddPersonModel) (*response.ResponseBase, error)
}

// UnimplementedPersonalServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPersonalServiceServer struct {
}

func (*UnimplementedPersonalServiceServer) AddPerson(ctx context.Context, req *AddPersonModel) (*response.ResponseBase, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPerson not implemented")
}

func RegisterPersonalServiceServer(s *grpc.Server, srv PersonalServiceServer) {
	s.RegisterService(&_PersonalService_serviceDesc, srv)
}

func _PersonalService_AddPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPersonModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonalServiceServer).AddPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/personal.PersonalService/AddPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonalServiceServer).AddPerson(ctx, req.(*AddPersonModel))
	}
	return interceptor(ctx, in, info, handler)
}

var _PersonalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "personal.PersonalService",
	HandlerType: (*PersonalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPerson",
			Handler:    _PersonalService_AddPerson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/person/person.proto",
}