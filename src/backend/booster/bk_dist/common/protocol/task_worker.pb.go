// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task_worker.proto

package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PBCompressType int32

const (
	PBCompressType_NONE PBCompressType = 0
	PBCompressType_LZO  PBCompressType = 1
	PBCompressType_LZ4  PBCompressType = 2
)

var PBCompressType_name = map[int32]string{
	0: "NONE",
	1: "LZO",
	2: "LZ4",
}
var PBCompressType_value = map[string]int32{
	"NONE": 0,
	"LZO":  1,
	"LZ4":  2,
}

func (x PBCompressType) Enum() *PBCompressType {
	p := new(PBCompressType)
	*p = x
	return p
}
func (x PBCompressType) String() string {
	return proto.EnumName(PBCompressType_name, int32(x))
}
func (x *PBCompressType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PBCompressType_value, data, "PBCompressType")
	if err != nil {
		return err
	}
	*x = PBCompressType(value)
	return nil
}
func (PBCompressType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_task_worker_6fa3dbe88e12366c, []int{0}
}

type PBCacheStatus int32

const (
	PBCacheStatus_NOFOUND           PBCacheStatus = 0
	PBCacheStatus_SUCCESS           PBCacheStatus = 1
	PBCacheStatus_ERRORWHILEFINDING PBCacheStatus = 2
	PBCacheStatus_ERRORWHILESAVING  PBCacheStatus = 3
)

var PBCacheStatus_name = map[int32]string{
	0: "NOFOUND",
	1: "SUCCESS",
	2: "ERRORWHILEFINDING",
	3: "ERRORWHILESAVING",
}
var PBCacheStatus_value = map[string]int32{
	"NOFOUND":           0,
	"SUCCESS":           1,
	"ERRORWHILEFINDING": 2,
	"ERRORWHILESAVING":  3,
}

func (x PBCacheStatus) Enum() *PBCacheStatus {
	p := new(PBCacheStatus)
	*p = x
	return p
}
func (x PBCacheStatus) String() string {
	return proto.EnumName(PBCacheStatus_name, int32(x))
}
func (x *PBCacheStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PBCacheStatus_value, data, "PBCacheStatus")
	if err != nil {
		return err
	}
	*x = PBCacheStatus(value)
	return nil
}
func (PBCacheStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_task_worker_6fa3dbe88e12366c, []int{1}
}

type PBCmdType int32

const (
	PBCmdType_DISPATCHTASKREQ PBCmdType = 0
	PBCmdType_DISPATCHTASKRSP PBCmdType = 1
	PBCmdType_SYNCTIMEREQ     PBCmdType = 2
	PBCmdType_SYNCTIMERSP     PBCmdType = 3
	PBCmdType_SENDFILEREQ     PBCmdType = 4
	PBCmdType_SENDFILERSP     PBCmdType = 5
	PBCmdType_CHECKCACHEREQ   PBCmdType = 6
	PBCmdType_CHECKCACHERSP   PBCmdType = 7
	PBCmdType_UNKNOWN         PBCmdType = 9999
)

var PBCmdType_name = map[int32]string{
	0:    "DISPATCHTASKREQ",
	1:    "DISPATCHTASKRSP",
	2:    "SYNCTIMEREQ",
	3:    "SYNCTIMERSP",
	4:    "SENDFILEREQ",
	5:    "SENDFILERSP",
	6:    "CHECKCACHEREQ",
	7:    "CHECKCACHERSP",
	9999: "UNKNOWN",
}
var PBCmdType_value = map[string]int32{
	"DISPATCHTASKREQ": 0,
	"DISPATCHTASKRSP": 1,
	"SYNCTIMEREQ":     2,
	"SYNCTIMERSP":     3,
	"SENDFILEREQ":     4,
	"SENDFILERSP":     5,
	"CHECKCACHEREQ":   6,
	"CHECKCACHERSP":   7,
	"UNKNOWN":         9999,
}

func (x PBCmdType) Enum() *PBCmdType {
	p := new(PBCmdType)
	*p = x
	return p
}
func (x PBCmdType) String() string {
	return proto.EnumName(PBCmdType_name, int32(x))
}
func (x *PBCmdType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PBCmdType_value, data, "PBCmdType")
	if err != nil {
		return err
	}
	*x = PBCmdType(value)
	return nil
}
func (PBCmdType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_task_worker_6fa3dbe88e12366c, []int{2}
}

type PBFileDesc struct {
	Fullpath       *string         `protobuf:"bytes,1,req,name=fullpath" json:"fullpath,omitempty"`
	Size           *int64          `protobuf:"varint,2,req,name=size" json:"size,omitempty"`
	Md5            *string         `protobuf:"bytes,3,req,name=md5" json:"md5,omitempty"`
	Compresstype   *PBCompressType `protobuf:"varint,4,req,name=compresstype,enum=protocol.PBCompressType" json:"compresstype,omitempty"`
	Compressedsize *int64          `protobuf:"varint,5,req,name=compressedsize" json:"compressedsize,omitempty"`
	Buffer         []byte          `protobuf:"bytes,6,opt,name=buffer" json:"buffer,omitempty"`
	// to specified relative path in target
	Targetrelativepath   *string  `protobuf:"bytes,7,opt,name=targetrelativepath" json:"targetrelativepath,omitempty"`
	Filemode             *uint32  `protobuf:"varint,8,opt,name=filemode" json:"filemode,omitempty"`
	Linktarget           []byte   `protobuf:"bytes,9,opt,name=linktarget" json:"linktarget,omitempty"`
	Modifytime           *int64   `protobuf:"varint,10,opt,name=modifytime" json:"modifytime,omitempty"`
	Accesstime           *int64   `protobuf:"varint,11,opt,name=accesstime" json:"accesstime,omitempty"`
	Createtime           *int64   `protobuf:"varint,12,opt,name=createtime" json:"createtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBFileDesc) Reset()         { *m = PBFileDesc{} }
func (m *PBFileDesc) String() string { return proto.CompactTextString(m) }
func (*PBFileDesc) ProtoMessage()    {}
func (*PBFileDesc) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_worker_6fa3dbe88e12366c, []int{0}
}
func (m *PBFileDesc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBFileDesc.Unmarshal(m, b)
}
func (m *PBFileDesc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBFileDesc.Marshal(b, m, deterministic)
}
func (dst *PBFileDesc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBFileDesc.Merge(dst, src)
}
func (m *PBFileDesc) XXX_Size() int {
	return xxx_messageInfo_PBFileDesc.Size(m)
}
func (m *PBFileDesc) XXX_DiscardUnknown() {
	xxx_messageInfo_PBFileDesc.DiscardUnknown(m)
}

var xxx_messageInfo_PBFileDesc proto.InternalMessageInfo

func (m *PBFileDesc) GetFullpath() string {
	if m != nil && m.Fullpath != nil {
		return *m.Fullpath
	}
	return ""
}

func (m *PBFileDesc) GetSize() int64 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

func (m *PBFileDesc) GetMd5() string {
	if m != nil && m.Md5 != nil {
		return *m.Md5
	}
	return ""
}

func (m *PBFileDesc) GetCompresstype() PBCompressType {
	if m != nil && m.Compresstype != nil {
		return *m.Compresstype
	}
	return PBCompressType_NONE
}

func (m *PBFileDesc) GetCompressedsize() int64 {
	if m != nil && m.Compressedsize != nil {
		return *m.Compressedsize
	}
	return 0
}

func (m *PBFileDesc) GetBuffer() []byte {
	if m != nil {
		return m.Buffer
	}
	return nil
}

func (m *PBFileDesc) GetTargetrelativepath() string {
	if m != nil && m.Targetrelativepath != nil {
		return *m.Targetrelativepath
	}
	return ""
}

func (m *PBFileDesc) GetFilemode() uint32 {
	if m != nil && m.Filemode != nil {
		return *m.Filemode
	}
	return 0
}

func (m *PBFileDesc) GetLinktarget() []byte {
	if m != nil {
		return m.Linktarget
	}
	return nil
}

func (m *PBFileDesc) GetModifytime() int64 {
	if m != nil && m.Modifytime != nil {
		return *m.Modifytime
	}
	return 0
}

func (m *PBFileDesc) GetAccesstime() int64 {
	if m != nil && m.Accesstime != nil {
		return *m.Accesstime
	}
	return 0
}

func (m *PBFileDesc) GetCreatetime() int64 {
	if m != nil && m.Createtime != nil {
		return *m.Createtime
	}
	return 0
}

type PBFileResult struct {
	Fullpath             *string  `protobuf:"bytes,1,req,name=fullpath" json:"fullpath,omitempty"`
	Retcode              *int32   `protobuf:"varint,2,req,name=retcode" json:"retcode,omitempty"`
	Targetrelativepath   *string  `protobuf:"bytes,3,opt,name=targetrelativepath" json:"targetrelativepath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBFileResult) Reset()         { *m = PBFileResult{} }
func (m *PBFileResult) String() string { return proto.CompactTextString(m) }
func (*PBFileResult) ProtoMessage()    {}
func (*PBFileResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_worker_6fa3dbe88e12366c, []int{1}
}
func (m *PBFileResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBFileResult.Unmarshal(m, b)
}
func (m *PBFileResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBFileResult.Marshal(b, m, deterministic)
}
func (dst *PBFileResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBFileResult.Merge(dst, src)
}
func (m *PBFileResult) XXX_Size() int {
	return xxx_messageInfo_PBFileResult.Size(m)
}
func (m *PBFileResult) XXX_DiscardUnknown() {
	xxx_messageInfo_PBFileResult.DiscardUnknown(m)
}

var xxx_messageInfo_PBFileResult proto.InternalMessageInfo

func (m *PBFileResult) GetFullpath() string {
	if m != nil && m.Fullpath != nil {
		return *m.Fullpath
	}
	return ""
}

func (m *PBFileResult) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *PBFileResult) GetTargetrelativepath() string {
	if m != nil && m.Targetrelativepath != nil {
		return *m.Targetrelativepath
	}
	return ""
}

type PBCommand struct {
	Workdir              *string       `protobuf:"bytes,1,req,name=workdir" json:"workdir,omitempty"`
	Exepath              *string       `protobuf:"bytes,2,req,name=exepath" json:"exepath,omitempty"`
	Exename              *string       `protobuf:"bytes,3,req,name=exename" json:"exename,omitempty"`
	Params               []string      `protobuf:"bytes,4,rep,name=params" json:"params,omitempty"`
	Inputfiles           []*PBFileDesc `protobuf:"bytes,5,rep,name=inputfiles" json:"inputfiles,omitempty"`
	Resultfiles          []string      `protobuf:"bytes,6,rep,name=resultfiles" json:"resultfiles,omitempty"`
	Env                  [][]byte      `protobuf:"bytes,7,rep,name=env" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PBCommand) Reset()         { *m = PBCommand{} }
func (m *PBCommand) String() string { return proto.CompactTextString(m) }
func (*PBCommand) ProtoMessage()    {}
func (*PBCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_worker_6fa3dbe88e12366c, []int{2}
}
func (m *PBCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBCommand.Unmarshal(m, b)
}
func (m *PBCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBCommand.Marshal(b, m, deterministic)
}
func (dst *PBCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBCommand.Merge(dst, src)
}
func (m *PBCommand) XXX_Size() int {
	return xxx_messageInfo_PBCommand.Size(m)
}
func (m *PBCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_PBCommand.DiscardUnknown(m)
}

var xxx_messageInfo_PBCommand proto.InternalMessageInfo

func (m *PBCommand) GetWorkdir() string {
	if m != nil && m.Workdir != nil {
		return *m.Workdir
	}
	return ""
}

func (m *PBCommand) GetExepath() string {
	if m != nil && m.Exepath != nil {
		return *m.Exepath
	}
	return ""
}

func (m *PBCommand) GetExename() string {
	if m != nil && m.Exename != nil {
		return *m.Exename
	}
	return ""
}

func (m *PBCommand) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *PBCommand) GetInputfiles() []*PBFileDesc {
	if m != nil {
		return m.Inputfiles
	}
	return nil
}

func (m *PBCommand) GetResultfiles() []string {
	if m != nil {
		return m.Resultfiles
	}
	return nil
}

func (m *PBCommand) GetEnv() [][]byte {
	if m != nil {
		return m.Env
	}
	return nil
}

type PBStatEntry struct {
	Key                  *string  `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Time                 *int64   `protobuf:"varint,2,req,name=time" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBStatEntry) Reset()         { *m = PBStatEntry{} }
func (m *PBStatEntry) String() string { return proto.CompactTextString(m) }
func (*PBStatEntry) ProtoMessage()    {}
func (*PBStatEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_worker_6fa3dbe88e12366c, []int{3}
}
func (m *PBStatEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBStatEntry.Unmarshal(m, b)
}
func (m *PBStatEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBStatEntry.Marshal(b, m, deterministic)
}
func (dst *PBStatEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBStatEntry.Merge(dst, src)
}
func (m *PBStatEntry) XXX_Size() int {
	return xxx_messageInfo_PBStatEntry.Size(m)
}
func (m *PBStatEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_PBStatEntry.DiscardUnknown(m)
}

var xxx_messageInfo_PBStatEntry proto.InternalMessageInfo

func (m *PBStatEntry) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *PBStatEntry) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

type PBResult struct {
	Cmd                  *PBCommand     `protobuf:"bytes,1,req,name=cmd" json:"cmd,omitempty"`
	Retcode              *int32         `protobuf:"varint,2,req,name=retcode" json:"retcode,omitempty"`
	Outputmessage        *string        `protobuf:"bytes,3,req,name=outputmessage" json:"outputmessage,omitempty"`
	Errormessage         *string        `protobuf:"bytes,4,req,name=errormessage" json:"errormessage,omitempty"`
	Resultfiles          []*PBFileDesc  `protobuf:"bytes,5,rep,name=resultfiles" json:"resultfiles,omitempty"`
	Stats                []*PBStatEntry `protobuf:"bytes,6,rep,name=stats" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PBResult) Reset()         { *m = PBResult{} }
func (m *PBResult) String() string { return proto.CompactTextString(m) }
func (*PBResult) ProtoMessage()    {}
func (*PBResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_worker_6fa3dbe88e12366c, []int{4}
}
func (m *PBResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBResult.Unmarshal(m, b)
}
func (m *PBResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBResult.Marshal(b, m, deterministic)
}
func (dst *PBResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBResult.Merge(dst, src)
}
func (m *PBResult) XXX_Size() int {
	return xxx_messageInfo_PBResult.Size(m)
}
func (m *PBResult) XXX_DiscardUnknown() {
	xxx_messageInfo_PBResult.DiscardUnknown(m)
}

var xxx_messageInfo_PBResult proto.InternalMessageInfo

func (m *PBResult) GetCmd() *PBCommand {
	if m != nil {
		return m.Cmd
	}
	return nil
}

func (m *PBResult) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *PBResult) GetOutputmessage() string {
	if m != nil && m.Outputmessage != nil {
		return *m.Outputmessage
	}
	return ""
}

func (m *PBResult) GetErrormessage() string {
	if m != nil && m.Errormessage != nil {
		return *m.Errormessage
	}
	return ""
}

func (m *PBResult) GetResultfiles() []*PBFileDesc {
	if m != nil {
		return m.Resultfiles
	}
	return nil
}

func (m *PBResult) GetStats() []*PBStatEntry {
	if m != nil {
		return m.Stats
	}
	return nil
}

type PBCacheParam struct {
	Name                 []byte   `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Md5                  []byte   `protobuf:"bytes,2,req,name=md5" json:"md5,omitempty"`
	Size                 *int64   `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	Target               []byte   `protobuf:"bytes,4,req,name=target" json:"target,omitempty"`
	Overwrite            *int32   `protobuf:"varint,5,opt,name=overwrite" json:"overwrite,omitempty"`
	Filemode             *uint32  `protobuf:"varint,6,opt,name=filemode" json:"filemode,omitempty"`
	Linktarget           []byte   `protobuf:"bytes,7,opt,name=linktarget" json:"linktarget,omitempty"`
	Modifytime           *int64   `protobuf:"varint,8,opt,name=modifytime" json:"modifytime,omitempty"`
	Accesstime           *int64   `protobuf:"varint,9,opt,name=accesstime" json:"accesstime,omitempty"`
	Createtime           *int64   `protobuf:"varint,10,opt,name=createtime" json:"createtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBCacheParam) Reset()         { *m = PBCacheParam{} }
func (m *PBCacheParam) String() string { return proto.CompactTextString(m) }
func (*PBCacheParam) ProtoMessage()    {}
func (*PBCacheParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_worker_6fa3dbe88e12366c, []int{5}
}
func (m *PBCacheParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBCacheParam.Unmarshal(m, b)
}
func (m *PBCacheParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBCacheParam.Marshal(b, m, deterministic)
}
func (dst *PBCacheParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBCacheParam.Merge(dst, src)
}
func (m *PBCacheParam) XXX_Size() int {
	return xxx_messageInfo_PBCacheParam.Size(m)
}
func (m *PBCacheParam) XXX_DiscardUnknown() {
	xxx_messageInfo_PBCacheParam.DiscardUnknown(m)
}

var xxx_messageInfo_PBCacheParam proto.InternalMessageInfo

func (m *PBCacheParam) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *PBCacheParam) GetMd5() []byte {
	if m != nil {
		return m.Md5
	}
	return nil
}

func (m *PBCacheParam) GetSize() int64 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

func (m *PBCacheParam) GetTarget() []byte {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *PBCacheParam) GetOverwrite() int32 {
	if m != nil && m.Overwrite != nil {
		return *m.Overwrite
	}
	return 0
}

func (m *PBCacheParam) GetFilemode() uint32 {
	if m != nil && m.Filemode != nil {
		return *m.Filemode
	}
	return 0
}

func (m *PBCacheParam) GetLinktarget() []byte {
	if m != nil {
		return m.Linktarget
	}
	return nil
}

func (m *PBCacheParam) GetModifytime() int64 {
	if m != nil && m.Modifytime != nil {
		return *m.Modifytime
	}
	return 0
}

func (m *PBCacheParam) GetAccesstime() int64 {
	if m != nil && m.Accesstime != nil {
		return *m.Accesstime
	}
	return 0
}

func (m *PBCacheParam) GetCreatetime() int64 {
	if m != nil && m.Createtime != nil {
		return *m.Createtime
	}
	return 0
}

type PBCacheResult struct {
	Status               *PBCacheStatus `protobuf:"varint,1,req,name=status,enum=protocol.PBCacheStatus" json:"status,omitempty"`
	Reason               []byte         `protobuf:"bytes,2,req,name=reason" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PBCacheResult) Reset()         { *m = PBCacheResult{} }
func (m *PBCacheResult) String() string { return proto.CompactTextString(m) }
func (*PBCacheResult) ProtoMessage()    {}
func (*PBCacheResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_worker_6fa3dbe88e12366c, []int{6}
}
func (m *PBCacheResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBCacheResult.Unmarshal(m, b)
}
func (m *PBCacheResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBCacheResult.Marshal(b, m, deterministic)
}
func (dst *PBCacheResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBCacheResult.Merge(dst, src)
}
func (m *PBCacheResult) XXX_Size() int {
	return xxx_messageInfo_PBCacheResult.Size(m)
}
func (m *PBCacheResult) XXX_DiscardUnknown() {
	xxx_messageInfo_PBCacheResult.DiscardUnknown(m)
}

var xxx_messageInfo_PBCacheResult proto.InternalMessageInfo

func (m *PBCacheResult) GetStatus() PBCacheStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return PBCacheStatus_NOFOUND
}

func (m *PBCacheResult) GetReason() []byte {
	if m != nil {
		return m.Reason
	}
	return nil
}

type PBHead struct {
	Version              *string    `protobuf:"bytes,1,req,name=version" json:"version,omitempty"`
	Magic                *string    `protobuf:"bytes,2,req,name=magic" json:"magic,omitempty"`
	Bodylen              *int32     `protobuf:"varint,3,req,name=bodylen" json:"bodylen,omitempty"`
	Buflen               *int64     `protobuf:"varint,4,req,name=buflen" json:"buflen,omitempty"`
	Cmdtype              *PBCmdType `protobuf:"varint,5,req,name=cmdtype,enum=protocol.PBCmdType" json:"cmdtype,omitempty"`
	Business             *string    `protobuf:"bytes,6,opt,name=business" json:"business,omitempty"`
	Taskid               *string    `protobuf:"bytes,7,opt,name=taskid" json:"taskid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PBHead) Reset()         { *m = PBHead{} }
func (m *PBHead) String() string { return proto.CompactTextString(m) }
func (*PBHead) ProtoMessage()    {}
func (*PBHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_worker_6fa3dbe88e12366c, []int{7}
}
func (m *PBHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBHead.Unmarshal(m, b)
}
func (m *PBHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBHead.Marshal(b, m, deterministic)
}
func (dst *PBHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBHead.Merge(dst, src)
}
func (m *PBHead) XXX_Size() int {
	return xxx_messageInfo_PBHead.Size(m)
}
func (m *PBHead) XXX_DiscardUnknown() {
	xxx_messageInfo_PBHead.DiscardUnknown(m)
}

var xxx_messageInfo_PBHead proto.InternalMessageInfo

func (m *PBHead) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *PBHead) GetMagic() string {
	if m != nil && m.Magic != nil {
		return *m.Magic
	}
	return ""
}

func (m *PBHead) GetBodylen() int32 {
	if m != nil && m.Bodylen != nil {
		return *m.Bodylen
	}
	return 0
}

func (m *PBHead) GetBuflen() int64 {
	if m != nil && m.Buflen != nil {
		return *m.Buflen
	}
	return 0
}

func (m *PBHead) GetCmdtype() PBCmdType {
	if m != nil && m.Cmdtype != nil {
		return *m.Cmdtype
	}
	return PBCmdType_DISPATCHTASKREQ
}

func (m *PBHead) GetBusiness() string {
	if m != nil && m.Business != nil {
		return *m.Business
	}
	return ""
}

func (m *PBHead) GetTaskid() string {
	if m != nil && m.Taskid != nil {
		return *m.Taskid
	}
	return ""
}

type PBBodyDispatchTaskReq struct {
	Cmds                 []*PBCommand `protobuf:"bytes,1,rep,name=cmds" json:"cmds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PBBodyDispatchTaskReq) Reset()         { *m = PBBodyDispatchTaskReq{} }
func (m *PBBodyDispatchTaskReq) String() string { return proto.CompactTextString(m) }
func (*PBBodyDispatchTaskReq) ProtoMessage()    {}
func (*PBBodyDispatchTaskReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_worker_6fa3dbe88e12366c, []int{8}
}
func (m *PBBodyDispatchTaskReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBBodyDispatchTaskReq.Unmarshal(m, b)
}
func (m *PBBodyDispatchTaskReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBBodyDispatchTaskReq.Marshal(b, m, deterministic)
}
func (dst *PBBodyDispatchTaskReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBBodyDispatchTaskReq.Merge(dst, src)
}
func (m *PBBodyDispatchTaskReq) XXX_Size() int {
	return xxx_messageInfo_PBBodyDispatchTaskReq.Size(m)
}
func (m *PBBodyDispatchTaskReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PBBodyDispatchTaskReq.DiscardUnknown(m)
}

var xxx_messageInfo_PBBodyDispatchTaskReq proto.InternalMessageInfo

func (m *PBBodyDispatchTaskReq) GetCmds() []*PBCommand {
	if m != nil {
		return m.Cmds
	}
	return nil
}

type PBBodyDispatchTaskRsp struct {
	Results              []*PBResult `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PBBodyDispatchTaskRsp) Reset()         { *m = PBBodyDispatchTaskRsp{} }
func (m *PBBodyDispatchTaskRsp) String() string { return proto.CompactTextString(m) }
func (*PBBodyDispatchTaskRsp) ProtoMessage()    {}
func (*PBBodyDispatchTaskRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_worker_6fa3dbe88e12366c, []int{9}
}
func (m *PBBodyDispatchTaskRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBBodyDispatchTaskRsp.Unmarshal(m, b)
}
func (m *PBBodyDispatchTaskRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBBodyDispatchTaskRsp.Marshal(b, m, deterministic)
}
func (dst *PBBodyDispatchTaskRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBBodyDispatchTaskRsp.Merge(dst, src)
}
func (m *PBBodyDispatchTaskRsp) XXX_Size() int {
	return xxx_messageInfo_PBBodyDispatchTaskRsp.Size(m)
}
func (m *PBBodyDispatchTaskRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_PBBodyDispatchTaskRsp.DiscardUnknown(m)
}

var xxx_messageInfo_PBBodyDispatchTaskRsp proto.InternalMessageInfo

func (m *PBBodyDispatchTaskRsp) GetResults() []*PBResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type PBBodySyncTimeRsp struct {
	Timenanosecond       *int64   `protobuf:"varint,1,req,name=timenanosecond" json:"timenanosecond,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBBodySyncTimeRsp) Reset()         { *m = PBBodySyncTimeRsp{} }
func (m *PBBodySyncTimeRsp) String() string { return proto.CompactTextString(m) }
func (*PBBodySyncTimeRsp) ProtoMessage()    {}
func (*PBBodySyncTimeRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_worker_6fa3dbe88e12366c, []int{10}
}
func (m *PBBodySyncTimeRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBBodySyncTimeRsp.Unmarshal(m, b)
}
func (m *PBBodySyncTimeRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBBodySyncTimeRsp.Marshal(b, m, deterministic)
}
func (dst *PBBodySyncTimeRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBBodySyncTimeRsp.Merge(dst, src)
}
func (m *PBBodySyncTimeRsp) XXX_Size() int {
	return xxx_messageInfo_PBBodySyncTimeRsp.Size(m)
}
func (m *PBBodySyncTimeRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_PBBodySyncTimeRsp.DiscardUnknown(m)
}

var xxx_messageInfo_PBBodySyncTimeRsp proto.InternalMessageInfo

func (m *PBBodySyncTimeRsp) GetTimenanosecond() int64 {
	if m != nil && m.Timenanosecond != nil {
		return *m.Timenanosecond
	}
	return 0
}

type PBBodySendFileReq struct {
	Inputfiles           []*PBFileDesc `protobuf:"bytes,1,rep,name=inputfiles" json:"inputfiles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PBBodySendFileReq) Reset()         { *m = PBBodySendFileReq{} }
func (m *PBBodySendFileReq) String() string { return proto.CompactTextString(m) }
func (*PBBodySendFileReq) ProtoMessage()    {}
func (*PBBodySendFileReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_worker_6fa3dbe88e12366c, []int{11}
}
func (m *PBBodySendFileReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBBodySendFileReq.Unmarshal(m, b)
}
func (m *PBBodySendFileReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBBodySendFileReq.Marshal(b, m, deterministic)
}
func (dst *PBBodySendFileReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBBodySendFileReq.Merge(dst, src)
}
func (m *PBBodySendFileReq) XXX_Size() int {
	return xxx_messageInfo_PBBodySendFileReq.Size(m)
}
func (m *PBBodySendFileReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PBBodySendFileReq.DiscardUnknown(m)
}

var xxx_messageInfo_PBBodySendFileReq proto.InternalMessageInfo

func (m *PBBodySendFileReq) GetInputfiles() []*PBFileDesc {
	if m != nil {
		return m.Inputfiles
	}
	return nil
}

type PBBodySendFileRsp struct {
	Results              []*PBFileResult `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PBBodySendFileRsp) Reset()         { *m = PBBodySendFileRsp{} }
func (m *PBBodySendFileRsp) String() string { return proto.CompactTextString(m) }
func (*PBBodySendFileRsp) ProtoMessage()    {}
func (*PBBodySendFileRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_worker_6fa3dbe88e12366c, []int{12}
}
func (m *PBBodySendFileRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBBodySendFileRsp.Unmarshal(m, b)
}
func (m *PBBodySendFileRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBBodySendFileRsp.Marshal(b, m, deterministic)
}
func (dst *PBBodySendFileRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBBodySendFileRsp.Merge(dst, src)
}
func (m *PBBodySendFileRsp) XXX_Size() int {
	return xxx_messageInfo_PBBodySendFileRsp.Size(m)
}
func (m *PBBodySendFileRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_PBBodySendFileRsp.DiscardUnknown(m)
}

var xxx_messageInfo_PBBodySendFileRsp proto.InternalMessageInfo

func (m *PBBodySendFileRsp) GetResults() []*PBFileResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type PBBodyCheckCacheReq struct {
	Params               []*PBCacheParam `protobuf:"bytes,1,rep,name=params" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PBBodyCheckCacheReq) Reset()         { *m = PBBodyCheckCacheReq{} }
func (m *PBBodyCheckCacheReq) String() string { return proto.CompactTextString(m) }
func (*PBBodyCheckCacheReq) ProtoMessage()    {}
func (*PBBodyCheckCacheReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_worker_6fa3dbe88e12366c, []int{13}
}
func (m *PBBodyCheckCacheReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBBodyCheckCacheReq.Unmarshal(m, b)
}
func (m *PBBodyCheckCacheReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBBodyCheckCacheReq.Marshal(b, m, deterministic)
}
func (dst *PBBodyCheckCacheReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBBodyCheckCacheReq.Merge(dst, src)
}
func (m *PBBodyCheckCacheReq) XXX_Size() int {
	return xxx_messageInfo_PBBodyCheckCacheReq.Size(m)
}
func (m *PBBodyCheckCacheReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PBBodyCheckCacheReq.DiscardUnknown(m)
}

var xxx_messageInfo_PBBodyCheckCacheReq proto.InternalMessageInfo

func (m *PBBodyCheckCacheReq) GetParams() []*PBCacheParam {
	if m != nil {
		return m.Params
	}
	return nil
}

type PBBodyCheckCacheRsp struct {
	Results              []*PBCacheResult `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PBBodyCheckCacheRsp) Reset()         { *m = PBBodyCheckCacheRsp{} }
func (m *PBBodyCheckCacheRsp) String() string { return proto.CompactTextString(m) }
func (*PBBodyCheckCacheRsp) ProtoMessage()    {}
func (*PBBodyCheckCacheRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_worker_6fa3dbe88e12366c, []int{14}
}
func (m *PBBodyCheckCacheRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBBodyCheckCacheRsp.Unmarshal(m, b)
}
func (m *PBBodyCheckCacheRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBBodyCheckCacheRsp.Marshal(b, m, deterministic)
}
func (dst *PBBodyCheckCacheRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBBodyCheckCacheRsp.Merge(dst, src)
}
func (m *PBBodyCheckCacheRsp) XXX_Size() int {
	return xxx_messageInfo_PBBodyCheckCacheRsp.Size(m)
}
func (m *PBBodyCheckCacheRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_PBBodyCheckCacheRsp.DiscardUnknown(m)
}

var xxx_messageInfo_PBBodyCheckCacheRsp proto.InternalMessageInfo

func (m *PBBodyCheckCacheRsp) GetResults() []*PBCacheResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*PBFileDesc)(nil), "protocol.PBFileDesc")
	proto.RegisterType((*PBFileResult)(nil), "protocol.PBFileResult")
	proto.RegisterType((*PBCommand)(nil), "protocol.PBCommand")
	proto.RegisterType((*PBStatEntry)(nil), "protocol.PBStatEntry")
	proto.RegisterType((*PBResult)(nil), "protocol.PBResult")
	proto.RegisterType((*PBCacheParam)(nil), "protocol.PBCacheParam")
	proto.RegisterType((*PBCacheResult)(nil), "protocol.PBCacheResult")
	proto.RegisterType((*PBHead)(nil), "protocol.PBHead")
	proto.RegisterType((*PBBodyDispatchTaskReq)(nil), "protocol.PBBodyDispatchTaskReq")
	proto.RegisterType((*PBBodyDispatchTaskRsp)(nil), "protocol.PBBodyDispatchTaskRsp")
	proto.RegisterType((*PBBodySyncTimeRsp)(nil), "protocol.PBBodySyncTimeRsp")
	proto.RegisterType((*PBBodySendFileReq)(nil), "protocol.PBBodySendFileReq")
	proto.RegisterType((*PBBodySendFileRsp)(nil), "protocol.PBBodySendFileRsp")
	proto.RegisterType((*PBBodyCheckCacheReq)(nil), "protocol.PBBodyCheckCacheReq")
	proto.RegisterType((*PBBodyCheckCacheRsp)(nil), "protocol.PBBodyCheckCacheRsp")
	proto.RegisterEnum("protocol.PBCompressType", PBCompressType_name, PBCompressType_value)
	proto.RegisterEnum("protocol.PBCacheStatus", PBCacheStatus_name, PBCacheStatus_value)
	proto.RegisterEnum("protocol.PBCmdType", PBCmdType_name, PBCmdType_value)
}

func init() { proto.RegisterFile("task_worker.proto", fileDescriptor_task_worker_6fa3dbe88e12366c) }

var fileDescriptor_task_worker_6fa3dbe88e12366c = []byte{
	// 957 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdd, 0x72, 0xe3, 0x34,
	0x18, 0xad, 0xed, 0xfc, 0x7e, 0x4e, 0x53, 0x57, 0xfd, 0xc1, 0xc3, 0x95, 0x31, 0x3b, 0xac, 0x29,
	0x4c, 0x2f, 0x3a, 0x70, 0xc3, 0xec, 0x0e, 0xd3, 0x38, 0x29, 0xc9, 0xb4, 0xeb, 0x98, 0x38, 0xdd,
	0x1d, 0xb8, 0x61, 0x54, 0x5b, 0x6d, 0x3d, 0x89, 0x7f, 0x2a, 0x29, 0x5d, 0xc2, 0x4b, 0xec, 0x70,
	0xcb, 0x3b, 0xf0, 0x04, 0xf0, 0x70, 0x8c, 0x64, 0x27, 0x69, 0xd2, 0x94, 0xab, 0x44, 0x47, 0xdf,
	0x91, 0xbf, 0x73, 0xbe, 0x23, 0xc1, 0x3e, 0xc7, 0x6c, 0xf2, 0xdb, 0xc7, 0x8c, 0x4e, 0x08, 0x3d,
	0xcd, 0x69, 0xc6, 0x33, 0xd4, 0x90, 0x3f, 0x61, 0x36, 0xb5, 0x3f, 0xa9, 0x00, 0x7e, 0xe7, 0x22,
	0x9e, 0x92, 0x2e, 0x61, 0x21, 0x32, 0xa0, 0x71, 0x3b, 0x9b, 0x4e, 0x73, 0xcc, 0xef, 0x4d, 0xc5,
	0x52, 0x9d, 0x26, 0x6a, 0x41, 0x85, 0xc5, 0x7f, 0x10, 0x53, 0xb5, 0x54, 0x47, 0x43, 0x3a, 0x68,
	0x49, 0xf4, 0xbd, 0xa9, 0xc9, 0xad, 0x53, 0x68, 0x85, 0x59, 0x92, 0x53, 0xc2, 0x18, 0x9f, 0xe7,
	0xc4, 0xac, 0x58, 0xaa, 0xd3, 0x3e, 0x33, 0x4f, 0x17, 0x87, 0x9f, 0xfa, 0x1d, 0xb7, 0xdc, 0x1f,
	0xcf, 0x73, 0x82, 0x8e, 0xa1, 0xbd, 0xa8, 0x27, 0x91, 0x3c, 0xb4, 0x2a, 0x0f, 0x6d, 0x43, 0xed,
	0x66, 0x76, 0x7b, 0x4b, 0xa8, 0x59, 0xb3, 0x14, 0xa7, 0x85, 0x3e, 0x07, 0xc4, 0x31, 0xbd, 0x23,
	0x9c, 0x92, 0x29, 0xe6, 0xf1, 0x23, 0x91, 0xed, 0xd4, 0x2d, 0xc5, 0x69, 0xca, 0x06, 0xe3, 0x29,
	0x49, 0xb2, 0x88, 0x98, 0x0d, 0x4b, 0x71, 0x76, 0x11, 0x02, 0x98, 0xc6, 0xe9, 0xa4, 0x60, 0x98,
	0x4d, 0x79, 0x02, 0x02, 0x48, 0xb2, 0x28, 0xbe, 0x9d, 0xf3, 0x38, 0x21, 0x26, 0x58, 0x8a, 0xa3,
	0x09, 0x0c, 0x87, 0xa1, 0xe8, 0x55, 0x60, 0xfa, 0x02, 0x0b, 0x29, 0xc1, 0x9c, 0x48, 0xac, 0x25,
	0x30, 0xfb, 0x1d, 0xb4, 0x0a, 0x43, 0x46, 0x84, 0xcd, 0xa6, 0x7c, 0x8b, 0x25, 0x7b, 0x50, 0xa7,
	0x84, 0x87, 0xa2, 0x05, 0xe1, 0x4a, 0xf5, 0x85, 0x86, 0x35, 0xd1, 0xb0, 0xfd, 0x97, 0x02, 0x4d,
	0xe9, 0x43, 0x82, 0xd3, 0x48, 0x50, 0xc5, 0x20, 0xa2, 0x98, 0xae, 0xce, 0x22, 0xbf, 0x17, 0xf5,
	0xea, 0x13, 0x20, 0xc5, 0x09, 0x29, 0x5d, 0x6e, 0x43, 0x2d, 0xc7, 0x14, 0x27, 0xcc, 0xac, 0x58,
	0x9a, 0xd3, 0x44, 0x0e, 0x40, 0x9c, 0xe6, 0x33, 0x2e, 0x6c, 0x60, 0x66, 0xd5, 0xd2, 0x1c, 0xfd,
	0xec, 0xf0, 0xa9, 0xe7, 0xcb, 0x61, 0x1e, 0x80, 0x4e, 0xa5, 0x86, 0xa2, 0xb4, 0x26, 0xe9, 0x3a,
	0x68, 0x24, 0x7d, 0x34, 0xeb, 0x96, 0xe6, 0xb4, 0x6c, 0x07, 0x74, 0xbf, 0x13, 0x70, 0xcc, 0x7b,
	0x29, 0xa7, 0x73, 0xb1, 0x37, 0x21, 0xf3, 0xd5, 0xe0, 0xa5, 0x2b, 0x72, 0xf0, 0xf6, 0x3f, 0x0a,
	0x34, 0xfc, 0x4e, 0x69, 0x89, 0x05, 0x5a, 0x98, 0x44, 0xb2, 0x4e, 0x3f, 0x3b, 0xd8, 0x98, 0xf7,
	0x42, 0xe7, 0xba, 0x45, 0x47, 0xb0, 0x9b, 0xcd, 0x78, 0x3e, 0xe3, 0x09, 0x61, 0x0c, 0xdf, 0x2d,
	0xc4, 0x1d, 0x42, 0x8b, 0x50, 0x9a, 0xd1, 0x05, 0x5a, 0x91, 0xe8, 0xd7, 0xeb, 0x8d, 0xff, 0x9f,
	0xc6, 0x57, 0x50, 0x65, 0x1c, 0xf3, 0x42, 0x9d, 0x7e, 0x76, 0xf4, 0xb4, 0x68, 0x29, 0xcc, 0xfe,
	0x57, 0x11, 0x43, 0x75, 0x71, 0x78, 0x4f, 0x7c, 0xe1, 0xa5, 0x10, 0x27, 0x2d, 0x16, 0x12, 0x5a,
	0x8b, 0x54, 0xab, 0x72, 0xb1, 0x08, 0xbc, 0x26, 0x13, 0xd2, 0x86, 0x5a, 0x99, 0xac, 0x8a, 0xdc,
	0xdd, 0x87, 0x66, 0xf6, 0x48, 0xe8, 0x47, 0x1a, 0x73, 0x11, 0x5f, 0xc5, 0xa9, 0xae, 0x45, 0xb2,
	0xb6, 0x25, 0x92, 0xf5, 0x2d, 0x91, 0x6c, 0x6c, 0x89, 0x64, 0x73, 0x4b, 0x24, 0x65, 0x74, 0xed,
	0x3e, 0xec, 0x96, 0xdd, 0x97, 0x03, 0x78, 0x0d, 0x35, 0xa1, 0x7a, 0xc6, 0xa4, 0x80, 0xf6, 0xd9,
	0x67, 0x6b, 0x33, 0x10, 0x85, 0x81, 0xdc, 0x16, 0xed, 0x53, 0x82, 0x59, 0x96, 0x16, 0xe2, 0xec,
	0x3f, 0x15, 0xa8, 0xf9, 0x9d, 0x3e, 0xc1, 0x72, 0x44, 0x8f, 0x84, 0xb2, 0x38, 0x4b, 0xcb, 0x81,
	0xef, 0x42, 0x35, 0xc1, 0x77, 0x71, 0xb8, 0x0a, 0xe2, 0x4d, 0x16, 0xcd, 0xa7, 0x24, 0x95, 0xb3,
	0xaa, 0x96, 0xd7, 0x54, 0xac, 0x2b, 0xf2, 0xda, 0xbe, 0x82, 0x7a, 0x98, 0x44, 0xf2, 0xe6, 0x57,
	0x65, 0x17, 0xeb, 0x49, 0x48, 0x22, 0x79, 0xe9, 0x0d, 0x68, 0xdc, 0xcc, 0x58, 0x9c, 0x12, 0xc6,
	0xa4, 0x3b, 0xcd, 0xc2, 0x52, 0x36, 0x89, 0xa3, 0xe2, 0x4a, 0xdb, 0x3f, 0xc0, 0x91, 0xdf, 0xe9,
	0x64, 0xd1, 0xbc, 0x1b, 0xb3, 0x1c, 0xf3, 0xf0, 0x7e, 0x8c, 0xd9, 0x64, 0x44, 0x1e, 0xd0, 0x17,
	0x50, 0x09, 0x93, 0x48, 0x68, 0xd4, 0x5e, 0xc8, 0x99, 0xfd, 0x66, 0x2b, 0x97, 0xe5, 0xe8, 0x4b,
	0x11, 0x40, 0xe1, 0xd5, 0x82, 0x8e, 0x9e, 0xd2, 0x0b, 0x1b, 0xed, 0x6f, 0x60, 0xbf, 0x60, 0x07,
	0xf3, 0x34, 0x1c, 0xc7, 0x09, 0x11, 0xcc, 0x63, 0x68, 0x0b, 0xeb, 0x53, 0x9c, 0x66, 0x8c, 0x84,
	0x59, 0x5a, 0xe4, 0x5c, 0xb3, 0xdf, 0x2e, 0x8b, 0x49, 0x1a, 0x15, 0xef, 0xc3, 0xc3, 0xc6, 0x65,
	0x54, 0x5e, 0x0e, 0xaa, 0xfd, 0xe6, 0x19, 0x9d, 0xe5, 0xe8, 0xf5, 0x66, 0x97, 0xc7, 0x9b, 0xdc,
	0xb2, 0xd3, 0xb7, 0x70, 0x50, 0xb0, 0xdd, 0x7b, 0x12, 0x4e, 0xca, 0x28, 0x3c, 0xa0, 0xaf, 0x96,
	0x6f, 0xc3, 0x16, 0xfa, 0x2a, 0xee, 0xf6, 0x8f, 0x5b, 0xe8, 0x2c, 0x47, 0xce, 0xe6, 0xe7, 0x9f,
	0xe7, 0xa8, 0xf8, 0xfe, 0xc9, 0xb7, 0xd0, 0xde, 0x78, 0xcc, 0x1b, 0x50, 0xf1, 0x86, 0x5e, 0xcf,
	0xd8, 0x41, 0x75, 0xd0, 0xae, 0x7e, 0x1d, 0x1a, 0x4a, 0xf1, 0xe7, 0x3b, 0x43, 0x3d, 0x79, 0xbf,
	0xcc, 0x6b, 0x19, 0x43, 0x1d, 0xea, 0xde, 0xf0, 0x62, 0x78, 0xed, 0x75, 0x8d, 0x1d, 0xb1, 0x08,
	0xae, 0x5d, 0xb7, 0x17, 0x04, 0x86, 0x82, 0x8e, 0x60, 0xbf, 0x37, 0x1a, 0x0d, 0x47, 0x1f, 0xfa,
	0x83, 0xab, 0xde, 0xc5, 0xc0, 0xeb, 0x0e, 0xbc, 0x9f, 0x0c, 0x15, 0x1d, 0x82, 0xb1, 0x82, 0x83,
	0xf3, 0xf7, 0x02, 0xd5, 0x4e, 0xfe, 0x2e, 0xde, 0xd2, 0x32, 0x59, 0x07, 0xb0, 0xd7, 0x1d, 0x04,
	0xfe, 0xf9, 0xd8, 0xed, 0x8f, 0xcf, 0x83, 0xcb, 0x51, 0xef, 0x67, 0x63, 0xe7, 0x19, 0x18, 0xf8,
	0x86, 0x82, 0xf6, 0x40, 0x0f, 0x7e, 0xf1, 0xdc, 0xf1, 0xe0, 0x5d, 0x4f, 0x54, 0xa9, 0x6b, 0x40,
	0xe0, 0x1b, 0x9a, 0x04, 0x7a, 0x5e, 0xf7, 0x62, 0x70, 0x25, 0x2b, 0x2a, 0x6b, 0x40, 0xe0, 0x1b,
	0x55, 0xb4, 0x0f, 0xbb, 0x6e, 0xbf, 0xe7, 0x5e, 0xba, 0xe7, 0x6e, 0x5f, 0xd6, 0xd4, 0x36, 0xa0,
	0xc0, 0x37, 0xea, 0xa8, 0x05, 0xf5, 0x6b, 0xef, 0xd2, 0x1b, 0x7e, 0xf0, 0x8c, 0x4f, 0xde, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x37, 0x25, 0x6b, 0x57, 0x7a, 0x07, 0x00, 0x00,
}