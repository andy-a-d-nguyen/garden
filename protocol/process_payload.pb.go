// Code generated by protoc-gen-gogo.
// source: process_payload.proto
// DO NOT EDIT!

package warden

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type ProcessPayload_Source int32

const (
	ProcessPayload_stdin  ProcessPayload_Source = 0
	ProcessPayload_stdout ProcessPayload_Source = 1
	ProcessPayload_stderr ProcessPayload_Source = 2
)

var ProcessPayload_Source_name = map[int32]string{
	0: "stdin",
	1: "stdout",
	2: "stderr",
}
var ProcessPayload_Source_value = map[string]int32{
	"stdin":  0,
	"stdout": 1,
	"stderr": 2,
}

func (x ProcessPayload_Source) Enum() *ProcessPayload_Source {
	p := new(ProcessPayload_Source)
	*p = x
	return p
}
func (x ProcessPayload_Source) String() string {
	return proto.EnumName(ProcessPayload_Source_name, int32(x))
}
func (x *ProcessPayload_Source) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProcessPayload_Source_value, data, "ProcessPayload_Source")
	if err != nil {
		return err
	}
	*x = ProcessPayload_Source(value)
	return nil
}

type ProcessPayload struct {
	ProcessId        *uint32                `protobuf:"varint,1,req,name=process_id" json:"process_id,omitempty"`
	Source           *ProcessPayload_Source `protobuf:"varint,2,opt,name=source,enum=warden.ProcessPayload_Source" json:"source,omitempty"`
	Data             *string                `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	ExitStatus       *uint32                `protobuf:"varint,4,opt,name=exit_status" json:"exit_status,omitempty"`
	Error            *string                `protobuf:"bytes,5,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ProcessPayload) Reset()         { *m = ProcessPayload{} }
func (m *ProcessPayload) String() string { return proto.CompactTextString(m) }
func (*ProcessPayload) ProtoMessage()    {}

func (m *ProcessPayload) GetProcessId() uint32 {
	if m != nil && m.ProcessId != nil {
		return *m.ProcessId
	}
	return 0
}

func (m *ProcessPayload) GetSource() ProcessPayload_Source {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return ProcessPayload_stdin
}

func (m *ProcessPayload) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

func (m *ProcessPayload) GetExitStatus() uint32 {
	if m != nil && m.ExitStatus != nil {
		return *m.ExitStatus
	}
	return 0
}

func (m *ProcessPayload) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func init() {
	proto.RegisterEnum("warden.ProcessPayload_Source", ProcessPayload_Source_name, ProcessPayload_Source_value)
}
