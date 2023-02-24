// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

// Package pbl contains generated bindings for API file pbl.api.
//
// Contents:
//
//	1 enum
//	2 structs
//	6 messages
package pbl

import (
	"strconv"

	fib_types "github.com/projectcalico/vpp-dataplane/v3/vpplink/binapi/vppapi/fib_types"
	_ "github.com/projectcalico/vpp-dataplane/v3/vpplink/binapi/vppapi/interface_types"
	ip_types "github.com/projectcalico/vpp-dataplane/v3/vpplink/binapi/vppapi/ip_types"
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "pbl"
	APIVersion = "0.1.0"
	VersionCrc = 0x87fa66f8
)

// PblClientFlags defines enum 'pbl_client_flags'.
type PblClientFlags uint32

const (
	PBL_API_FLAG_EXCLUSIVE PblClientFlags = 1
)

var (
	PblClientFlags_name = map[uint32]string{
		1: "PBL_API_FLAG_EXCLUSIVE",
	}
	PblClientFlags_value = map[string]uint32{
		"PBL_API_FLAG_EXCLUSIVE": 1,
	}
)

func (x PblClientFlags) String() string {
	s, ok := PblClientFlags_name[uint32(x)]
	if ok {
		return s
	}
	str := func(n uint32) string {
		s, ok := PblClientFlags_name[uint32(n)]
		if ok {
			return s
		}
		return "PblClientFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint32(0); i <= 32; i++ {
		val := uint32(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint32(x))
	}
	return s
}

// PblClient defines type 'pbl_client'.
type PblClient struct {
	ID         uint32            `binapi:"u32,name=id,default=4294967295" json:"id,omitempty"`
	Addr       ip_types.Address  `binapi:"address,name=addr" json:"addr,omitempty"`
	Paths      fib_types.FibPath `binapi:"fib_path,name=paths" json:"paths,omitempty"`
	Flags      uint8             `binapi:"u8,name=flags" json:"flags,omitempty"`
	TableID    uint32            `binapi:"u32,name=table_id" json:"table_id,omitempty"`
	NPorts     uint32            `binapi:"u32,name=n_ports" json:"-"`
	PortRanges []PblPortRange    `binapi:"pbl_port_range[n_ports],name=port_ranges" json:"port_ranges,omitempty"`
}

// PblPortRange defines type 'pbl_port_range'.
type PblPortRange struct {
	Start  uint16           `binapi:"u16,name=start" json:"start,omitempty"`
	End    uint16           `binapi:"u16,name=end" json:"end,omitempty"`
	Iproto ip_types.IPProto `binapi:"ip_proto,name=iproto" json:"iproto,omitempty"`
}

// PblClientDel defines message 'pbl_client_del'.
// InProgress: the message form may change in the future versions
type PblClientDel struct {
	ID uint32 `binapi:"u32,name=id" json:"id,omitempty"`
}

func (m *PblClientDel) Reset()               { *m = PblClientDel{} }
func (*PblClientDel) GetMessageName() string { return "pbl_client_del" }
func (*PblClientDel) GetCrcString() string   { return "3a91bde5" }
func (*PblClientDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PblClientDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.ID
	return size
}
func (m *PblClientDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.ID)
	return buf.Bytes(), nil
}
func (m *PblClientDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ID = buf.DecodeUint32()
	return nil
}

// PblClientDelReply defines message 'pbl_client_del_reply'.
// InProgress: the message form may change in the future versions
type PblClientDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *PblClientDelReply) Reset()               { *m = PblClientDelReply{} }
func (*PblClientDelReply) GetMessageName() string { return "pbl_client_del_reply" }
func (*PblClientDelReply) GetCrcString() string   { return "e8d4e804" }
func (*PblClientDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PblClientDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *PblClientDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *PblClientDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// PblClientDetails defines message 'pbl_client_details'.
// InProgress: the message form may change in the future versions
type PblClientDetails struct {
	Client PblClient `binapi:"pbl_client,name=client" json:"client,omitempty"`
}

func (m *PblClientDetails) Reset()               { *m = PblClientDetails{} }
func (*PblClientDetails) GetMessageName() string { return "pbl_client_details" }
func (*PblClientDetails) GetCrcString() string   { return "14278144" }
func (*PblClientDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PblClientDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Client.ID
	size += 1      // m.Client.Addr.Af
	size += 1 * 16 // m.Client.Addr.Un
	size += 4      // m.Client.Paths.SwIfIndex
	size += 4      // m.Client.Paths.TableID
	size += 4      // m.Client.Paths.RpfID
	size += 1      // m.Client.Paths.Weight
	size += 1      // m.Client.Paths.Preference
	size += 4      // m.Client.Paths.Type
	size += 4      // m.Client.Paths.Flags
	size += 4      // m.Client.Paths.Proto
	size += 1 * 16 // m.Client.Paths.Nh.Address
	size += 4      // m.Client.Paths.Nh.ViaLabel
	size += 4      // m.Client.Paths.Nh.ObjID
	size += 4      // m.Client.Paths.Nh.ClassifyTableIndex
	size += 1      // m.Client.Paths.NLabels
	for j3 := 0; j3 < 16; j3++ {
		size += 1 // m.Client.Paths.LabelStack[j3].IsUniform
		size += 4 // m.Client.Paths.LabelStack[j3].Label
		size += 1 // m.Client.Paths.LabelStack[j3].TTL
		size += 1 // m.Client.Paths.LabelStack[j3].Exp
	}
	size += 1 // m.Client.Flags
	size += 4 // m.Client.TableID
	size += 4 // m.Client.NPorts
	for j2 := 0; j2 < len(m.Client.PortRanges); j2++ {
		var s2 PblPortRange
		_ = s2
		if j2 < len(m.Client.PortRanges) {
			s2 = m.Client.PortRanges[j2]
		}
		size += 2 // s2.Start
		size += 2 // s2.End
		size += 1 // s2.Iproto
	}
	return size
}
func (m *PblClientDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Client.ID)
	buf.EncodeUint8(uint8(m.Client.Addr.Af))
	buf.EncodeBytes(m.Client.Addr.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.Client.Paths.SwIfIndex)
	buf.EncodeUint32(m.Client.Paths.TableID)
	buf.EncodeUint32(m.Client.Paths.RpfID)
	buf.EncodeUint8(m.Client.Paths.Weight)
	buf.EncodeUint8(m.Client.Paths.Preference)
	buf.EncodeUint32(uint32(m.Client.Paths.Type))
	buf.EncodeUint32(uint32(m.Client.Paths.Flags))
	buf.EncodeUint32(uint32(m.Client.Paths.Proto))
	buf.EncodeBytes(m.Client.Paths.Nh.Address.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.Client.Paths.Nh.ViaLabel)
	buf.EncodeUint32(m.Client.Paths.Nh.ObjID)
	buf.EncodeUint32(m.Client.Paths.Nh.ClassifyTableIndex)
	buf.EncodeUint8(m.Client.Paths.NLabels)
	for j2 := 0; j2 < 16; j2++ {
		buf.EncodeUint8(m.Client.Paths.LabelStack[j2].IsUniform)
		buf.EncodeUint32(m.Client.Paths.LabelStack[j2].Label)
		buf.EncodeUint8(m.Client.Paths.LabelStack[j2].TTL)
		buf.EncodeUint8(m.Client.Paths.LabelStack[j2].Exp)
	}
	buf.EncodeUint8(m.Client.Flags)
	buf.EncodeUint32(m.Client.TableID)
	buf.EncodeUint32(uint32(len(m.Client.PortRanges)))
	for j1 := 0; j1 < len(m.Client.PortRanges); j1++ {
		var v1 PblPortRange // PortRanges
		if j1 < len(m.Client.PortRanges) {
			v1 = m.Client.PortRanges[j1]
		}
		buf.EncodeUint16(v1.Start)
		buf.EncodeUint16(v1.End)
		buf.EncodeUint8(uint8(v1.Iproto))
	}
	return buf.Bytes(), nil
}
func (m *PblClientDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Client.ID = buf.DecodeUint32()
	m.Client.Addr.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Client.Addr.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Client.Paths.SwIfIndex = buf.DecodeUint32()
	m.Client.Paths.TableID = buf.DecodeUint32()
	m.Client.Paths.RpfID = buf.DecodeUint32()
	m.Client.Paths.Weight = buf.DecodeUint8()
	m.Client.Paths.Preference = buf.DecodeUint8()
	m.Client.Paths.Type = fib_types.FibPathType(buf.DecodeUint32())
	m.Client.Paths.Flags = fib_types.FibPathFlags(buf.DecodeUint32())
	m.Client.Paths.Proto = fib_types.FibPathNhProto(buf.DecodeUint32())
	copy(m.Client.Paths.Nh.Address.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Client.Paths.Nh.ViaLabel = buf.DecodeUint32()
	m.Client.Paths.Nh.ObjID = buf.DecodeUint32()
	m.Client.Paths.Nh.ClassifyTableIndex = buf.DecodeUint32()
	m.Client.Paths.NLabels = buf.DecodeUint8()
	for j2 := 0; j2 < 16; j2++ {
		m.Client.Paths.LabelStack[j2].IsUniform = buf.DecodeUint8()
		m.Client.Paths.LabelStack[j2].Label = buf.DecodeUint32()
		m.Client.Paths.LabelStack[j2].TTL = buf.DecodeUint8()
		m.Client.Paths.LabelStack[j2].Exp = buf.DecodeUint8()
	}
	m.Client.Flags = buf.DecodeUint8()
	m.Client.TableID = buf.DecodeUint32()
	m.Client.NPorts = buf.DecodeUint32()
	m.Client.PortRanges = make([]PblPortRange, m.Client.NPorts)
	for j1 := 0; j1 < len(m.Client.PortRanges); j1++ {
		m.Client.PortRanges[j1].Start = buf.DecodeUint16()
		m.Client.PortRanges[j1].End = buf.DecodeUint16()
		m.Client.PortRanges[j1].Iproto = ip_types.IPProto(buf.DecodeUint8())
	}
	return nil
}

// PblClientDump defines message 'pbl_client_dump'.
// InProgress: the message form may change in the future versions
type PblClientDump struct{}

func (m *PblClientDump) Reset()               { *m = PblClientDump{} }
func (*PblClientDump) GetMessageName() string { return "pbl_client_dump" }
func (*PblClientDump) GetCrcString() string   { return "51077d14" }
func (*PblClientDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PblClientDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *PblClientDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *PblClientDump) Unmarshal(b []byte) error {
	return nil
}

// PblClientUpdate defines message 'pbl_client_update'.
// InProgress: the message form may change in the future versions
type PblClientUpdate struct {
	Client PblClient `binapi:"pbl_client,name=client" json:"client,omitempty"`
}

func (m *PblClientUpdate) Reset()               { *m = PblClientUpdate{} }
func (*PblClientUpdate) GetMessageName() string { return "pbl_client_update" }
func (*PblClientUpdate) GetCrcString() string   { return "d83d6e65" }
func (*PblClientUpdate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *PblClientUpdate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Client.ID
	size += 1      // m.Client.Addr.Af
	size += 1 * 16 // m.Client.Addr.Un
	size += 4      // m.Client.Paths.SwIfIndex
	size += 4      // m.Client.Paths.TableID
	size += 4      // m.Client.Paths.RpfID
	size += 1      // m.Client.Paths.Weight
	size += 1      // m.Client.Paths.Preference
	size += 4      // m.Client.Paths.Type
	size += 4      // m.Client.Paths.Flags
	size += 4      // m.Client.Paths.Proto
	size += 1 * 16 // m.Client.Paths.Nh.Address
	size += 4      // m.Client.Paths.Nh.ViaLabel
	size += 4      // m.Client.Paths.Nh.ObjID
	size += 4      // m.Client.Paths.Nh.ClassifyTableIndex
	size += 1      // m.Client.Paths.NLabels
	for j3 := 0; j3 < 16; j3++ {
		size += 1 // m.Client.Paths.LabelStack[j3].IsUniform
		size += 4 // m.Client.Paths.LabelStack[j3].Label
		size += 1 // m.Client.Paths.LabelStack[j3].TTL
		size += 1 // m.Client.Paths.LabelStack[j3].Exp
	}
	size += 1 // m.Client.Flags
	size += 4 // m.Client.TableID
	size += 4 // m.Client.NPorts
	for j2 := 0; j2 < len(m.Client.PortRanges); j2++ {
		var s2 PblPortRange
		_ = s2
		if j2 < len(m.Client.PortRanges) {
			s2 = m.Client.PortRanges[j2]
		}
		size += 2 // s2.Start
		size += 2 // s2.End
		size += 1 // s2.Iproto
	}
	return size
}
func (m *PblClientUpdate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Client.ID)
	buf.EncodeUint8(uint8(m.Client.Addr.Af))
	buf.EncodeBytes(m.Client.Addr.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.Client.Paths.SwIfIndex)
	buf.EncodeUint32(m.Client.Paths.TableID)
	buf.EncodeUint32(m.Client.Paths.RpfID)
	buf.EncodeUint8(m.Client.Paths.Weight)
	buf.EncodeUint8(m.Client.Paths.Preference)
	buf.EncodeUint32(uint32(m.Client.Paths.Type))
	buf.EncodeUint32(uint32(m.Client.Paths.Flags))
	buf.EncodeUint32(uint32(m.Client.Paths.Proto))
	buf.EncodeBytes(m.Client.Paths.Nh.Address.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.Client.Paths.Nh.ViaLabel)
	buf.EncodeUint32(m.Client.Paths.Nh.ObjID)
	buf.EncodeUint32(m.Client.Paths.Nh.ClassifyTableIndex)
	buf.EncodeUint8(m.Client.Paths.NLabels)
	for j2 := 0; j2 < 16; j2++ {
		buf.EncodeUint8(m.Client.Paths.LabelStack[j2].IsUniform)
		buf.EncodeUint32(m.Client.Paths.LabelStack[j2].Label)
		buf.EncodeUint8(m.Client.Paths.LabelStack[j2].TTL)
		buf.EncodeUint8(m.Client.Paths.LabelStack[j2].Exp)
	}
	buf.EncodeUint8(m.Client.Flags)
	buf.EncodeUint32(m.Client.TableID)
	buf.EncodeUint32(uint32(len(m.Client.PortRanges)))
	for j1 := 0; j1 < len(m.Client.PortRanges); j1++ {
		var v1 PblPortRange // PortRanges
		if j1 < len(m.Client.PortRanges) {
			v1 = m.Client.PortRanges[j1]
		}
		buf.EncodeUint16(v1.Start)
		buf.EncodeUint16(v1.End)
		buf.EncodeUint8(uint8(v1.Iproto))
	}
	return buf.Bytes(), nil
}
func (m *PblClientUpdate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Client.ID = buf.DecodeUint32()
	m.Client.Addr.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Client.Addr.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Client.Paths.SwIfIndex = buf.DecodeUint32()
	m.Client.Paths.TableID = buf.DecodeUint32()
	m.Client.Paths.RpfID = buf.DecodeUint32()
	m.Client.Paths.Weight = buf.DecodeUint8()
	m.Client.Paths.Preference = buf.DecodeUint8()
	m.Client.Paths.Type = fib_types.FibPathType(buf.DecodeUint32())
	m.Client.Paths.Flags = fib_types.FibPathFlags(buf.DecodeUint32())
	m.Client.Paths.Proto = fib_types.FibPathNhProto(buf.DecodeUint32())
	copy(m.Client.Paths.Nh.Address.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Client.Paths.Nh.ViaLabel = buf.DecodeUint32()
	m.Client.Paths.Nh.ObjID = buf.DecodeUint32()
	m.Client.Paths.Nh.ClassifyTableIndex = buf.DecodeUint32()
	m.Client.Paths.NLabels = buf.DecodeUint8()
	for j2 := 0; j2 < 16; j2++ {
		m.Client.Paths.LabelStack[j2].IsUniform = buf.DecodeUint8()
		m.Client.Paths.LabelStack[j2].Label = buf.DecodeUint32()
		m.Client.Paths.LabelStack[j2].TTL = buf.DecodeUint8()
		m.Client.Paths.LabelStack[j2].Exp = buf.DecodeUint8()
	}
	m.Client.Flags = buf.DecodeUint8()
	m.Client.TableID = buf.DecodeUint32()
	m.Client.NPorts = buf.DecodeUint32()
	m.Client.PortRanges = make([]PblPortRange, m.Client.NPorts)
	for j1 := 0; j1 < len(m.Client.PortRanges); j1++ {
		m.Client.PortRanges[j1].Start = buf.DecodeUint16()
		m.Client.PortRanges[j1].End = buf.DecodeUint16()
		m.Client.PortRanges[j1].Iproto = ip_types.IPProto(buf.DecodeUint8())
	}
	return nil
}

// PblClientUpdateReply defines message 'pbl_client_update_reply'.
// InProgress: the message form may change in the future versions
type PblClientUpdateReply struct {
	Retval int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	ID     uint32 `binapi:"u32,name=id" json:"id,omitempty"`
}

func (m *PblClientUpdateReply) Reset()               { *m = PblClientUpdateReply{} }
func (*PblClientUpdateReply) GetMessageName() string { return "pbl_client_update_reply" }
func (*PblClientUpdateReply) GetCrcString() string   { return "e2fc8294" }
func (*PblClientUpdateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *PblClientUpdateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.ID
	return size
}
func (m *PblClientUpdateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.ID)
	return buf.Bytes(), nil
}
func (m *PblClientUpdateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.ID = buf.DecodeUint32()
	return nil
}

func init() { file_pbl_binapi_init() }
func file_pbl_binapi_init() {
	api.RegisterMessage((*PblClientDel)(nil), "pbl_client_del_3a91bde5")
	api.RegisterMessage((*PblClientDelReply)(nil), "pbl_client_del_reply_e8d4e804")
	api.RegisterMessage((*PblClientDetails)(nil), "pbl_client_details_14278144")
	api.RegisterMessage((*PblClientDump)(nil), "pbl_client_dump_51077d14")
	api.RegisterMessage((*PblClientUpdate)(nil), "pbl_client_update_d83d6e65")
	api.RegisterMessage((*PblClientUpdateReply)(nil), "pbl_client_update_reply_e2fc8294")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*PblClientDel)(nil),
		(*PblClientDelReply)(nil),
		(*PblClientDetails)(nil),
		(*PblClientDump)(nil),
		(*PblClientUpdate)(nil),
		(*PblClientUpdateReply)(nil),
	}
}
