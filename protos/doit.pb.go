// Code generated by protoc-gen-go.
// source: doit.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	doit.proto

It has these top-level messages:
	RegisterRequest
	RegisterReply
*/
package protos

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type RegisterRequest struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Address          *string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}

func (m *RegisterRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *RegisterRequest) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

type RegisterReply struct {
	Success          *bool  `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RegisterReply) Reset()         { *m = RegisterReply{} }
func (m *RegisterReply) String() string { return proto.CompactTextString(m) }
func (*RegisterReply) ProtoMessage()    {}

func (m *RegisterReply) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func init() {
}
