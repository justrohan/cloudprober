// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/probes/udp/proto/config.proto

package proto

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

type ProbeConf struct {
	// Export stats after these many milliseconds
	// NOTE: Setting stats export interval using this field doesn't work anymore.
	// This field will be removed after the release v0.10.3. To set
	// stats_export_interval, please modify the outer probe configuration.
	// probe {
	//   type: UDP
	//   stats_export_interval_msec: 10000
	//   udp_probe {}
	// }
	StatsExportIntervalMsec *int32 `protobuf:"varint,2,opt,name=stats_export_interval_msec,json=statsExportIntervalMsec,def=10000" json:"stats_export_interval_msec,omitempty"`
	// Port to send UDP Ping to (UDP Echo).  Should be same as
	// ProberConfig.udp_echo_server_port.
	// TODO(_): Can we just read this from ProberConfig?
	Port *int32 `protobuf:"varint,3,opt,name=port,def=31122" json:"port,omitempty"`
	// Number of sending side ports to use.
	NumTxPorts *int32 `protobuf:"varint,4,opt,name=num_tx_ports,json=numTxPorts,def=16" json:"num_tx_ports,omitempty"`
	// message max to account for MTU.
	MaxLength *int32 `protobuf:"varint,5,opt,name=max_length,json=maxLength,def=1300" json:"max_length,omitempty"`
	// IP proto: v4|v6.
	// This field is now deprecated and will be removed after the release v0.10.3.
	// ip_version can be configured in the outer layer of the config:
	// probe {
	//   type: UDP
	//   ip_version: 6
	//   udp_probe {}
	// }
	IpVersion *int32 `protobuf:"varint,6,opt,name=ip_version,json=ipVersion" json:"ip_version,omitempty"`
	// Changes the exported monitoring streams to be per port:
	// 1. Changes the streams names to total-per-port, success-per-port etc.
	// 2. Adds src_port and dst_port as stream labels.
	// Note that the field name is experimental and may change in the future.
	ExportMetricsByPort *bool `protobuf:"varint,7,opt,name=export_metrics_by_port,json=exportMetricsByPort,def=0" json:"export_metrics_by_port,omitempty"`
	// Whether to use all transmit ports per probe, per target.
	// Default is to probe each target once per probe and round-robin through the
	// source ports.
	// Setting this field to true changes the behavior to send traffic from all
	// ports to all targets in each probe.
	// For example, if num_tx_ports is set to 16, in every probe cycle, we'll send
	// 16 packets to every target (1 per tx port).
	// Note that setting this field to true will increase the probe traffic.
	UseAllTxPortsPerProbe *bool    `protobuf:"varint,8,opt,name=use_all_tx_ports_per_probe,json=useAllTxPortsPerProbe,def=0" json:"use_all_tx_ports_per_probe,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *ProbeConf) Reset()         { *m = ProbeConf{} }
func (m *ProbeConf) String() string { return proto.CompactTextString(m) }
func (*ProbeConf) ProtoMessage()    {}
func (*ProbeConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_f1551a718a2bffdc, []int{0}
}
func (m *ProbeConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeConf.Unmarshal(m, b)
}
func (m *ProbeConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeConf.Marshal(b, m, deterministic)
}
func (dst *ProbeConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeConf.Merge(dst, src)
}
func (m *ProbeConf) XXX_Size() int {
	return xxx_messageInfo_ProbeConf.Size(m)
}
func (m *ProbeConf) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeConf.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeConf proto.InternalMessageInfo

const Default_ProbeConf_StatsExportIntervalMsec int32 = 10000
const Default_ProbeConf_Port int32 = 31122
const Default_ProbeConf_NumTxPorts int32 = 16
const Default_ProbeConf_MaxLength int32 = 1300
const Default_ProbeConf_ExportMetricsByPort bool = false
const Default_ProbeConf_UseAllTxPortsPerProbe bool = false

func (m *ProbeConf) GetStatsExportIntervalMsec() int32 {
	if m != nil && m.StatsExportIntervalMsec != nil {
		return *m.StatsExportIntervalMsec
	}
	return Default_ProbeConf_StatsExportIntervalMsec
}

func (m *ProbeConf) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return Default_ProbeConf_Port
}

func (m *ProbeConf) GetNumTxPorts() int32 {
	if m != nil && m.NumTxPorts != nil {
		return *m.NumTxPorts
	}
	return Default_ProbeConf_NumTxPorts
}

func (m *ProbeConf) GetMaxLength() int32 {
	if m != nil && m.MaxLength != nil {
		return *m.MaxLength
	}
	return Default_ProbeConf_MaxLength
}

func (m *ProbeConf) GetIpVersion() int32 {
	if m != nil && m.IpVersion != nil {
		return *m.IpVersion
	}
	return 0
}

func (m *ProbeConf) GetExportMetricsByPort() bool {
	if m != nil && m.ExportMetricsByPort != nil {
		return *m.ExportMetricsByPort
	}
	return Default_ProbeConf_ExportMetricsByPort
}

func (m *ProbeConf) GetUseAllTxPortsPerProbe() bool {
	if m != nil && m.UseAllTxPortsPerProbe != nil {
		return *m.UseAllTxPortsPerProbe
	}
	return Default_ProbeConf_UseAllTxPortsPerProbe
}

func init() {
	proto.RegisterType((*ProbeConf)(nil), "cloudprober.probes.udp.ProbeConf")
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/probes/udp/proto/config.proto", fileDescriptor_config_f1551a718a2bffdc)
}

var fileDescriptor_config_f1551a718a2bffdc = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd0, 0x4d, 0x6b, 0x22, 0x31,
	0x18, 0x07, 0x70, 0x74, 0x9d, 0x5d, 0x0d, 0x7b, 0xca, 0xb2, 0x36, 0x15, 0x0a, 0xd2, 0xf6, 0xe0,
	0x69, 0x5e, 0x14, 0x7a, 0xb0, 0x97, 0xd6, 0xd2, 0x43, 0xa1, 0x82, 0x0c, 0xa5, 0xd7, 0x30, 0x8e,
	0x8f, 0x63, 0x20, 0x6f, 0xe4, 0x45, 0xc6, 0xaf, 0xd9, 0x4f, 0x54, 0x92, 0x91, 0xe2, 0x69, 0xe6,
	0xf9, 0xf3, 0x0b, 0xf9, 0xe7, 0x41, 0x8f, 0x0d, 0x73, 0x07, 0xbf, 0x4d, 0x6b, 0x25, 0xb2, 0x46,
	0xa9, 0x86, 0x43, 0x56, 0x73, 0xe5, 0x77, 0xda, 0xa8, 0x2d, 0x98, 0x2c, 0x7e, 0x6c, 0xe6, 0x77,
	0x3a, 0xfc, 0x3a, 0x95, 0xd5, 0x4a, 0xee, 0x59, 0x93, 0xc6, 0x01, 0x8f, 0x2f, 0x68, 0xda, 0xd1,
	0xd4, 0xef, 0xf4, 0xed, 0x57, 0x1f, 0x8d, 0x36, 0x61, 0x7c, 0x51, 0x72, 0x8f, 0x57, 0x68, 0x62,
	0x5d, 0xe5, 0x2c, 0x85, 0x56, 0x2b, 0xe3, 0x28, 0x93, 0x0e, 0xcc, 0xb1, 0xe2, 0x54, 0x58, 0xa8,
	0x49, 0x7f, 0xda, 0x9b, 0x25, 0xcb, 0xa4, 0xc8, 0xf3, 0x3c, 0x2f, 0xaf, 0x22, 0x7c, 0x8d, 0xee,
	0xed, 0xcc, 0xd6, 0x16, 0x6a, 0x7c, 0x8d, 0x06, 0x21, 0x23, 0xbf, 0x3a, 0xbd, 0x28, 0x8a, 0xf9,
	0xbc, 0x8c, 0x11, 0xbe, 0x47, 0x7f, 0xa5, 0x17, 0xd4, 0xb5, 0x34, 0x8c, 0x96, 0x0c, 0x22, 0xe9,
	0x17, 0x0f, 0x25, 0x92, 0x5e, 0x7c, 0xb4, 0x9b, 0x90, 0xe2, 0x3b, 0x84, 0x44, 0xd5, 0x52, 0x0e,
	0xb2, 0x71, 0x07, 0x92, 0x44, 0x33, 0x28, 0x16, 0x79, 0x5e, 0x8e, 0x44, 0xd5, 0xbe, 0xc7, 0x18,
	0xdf, 0x20, 0xc4, 0x34, 0x3d, 0x82, 0xb1, 0x4c, 0x49, 0xf2, 0x3b, 0xa0, 0x72, 0xc4, 0xf4, 0x67,
	0x17, 0xe0, 0x25, 0x1a, 0x9f, 0x9f, 0x20, 0xc0, 0x19, 0x56, 0x5b, 0xba, 0x3d, 0xc5, 0x4b, 0xc9,
	0x9f, 0x69, 0x6f, 0x36, 0x5c, 0x26, 0xfb, 0x8a, 0x5b, 0x28, 0xff, 0x75, 0x68, 0xdd, 0x99, 0xd5,
	0x29, 0x14, 0xc0, 0x4f, 0x68, 0xe2, 0x2d, 0xd0, 0x8a, 0xf3, 0x9f, 0xa6, 0x54, 0x83, 0xa1, 0x71,
	0x6b, 0x64, 0x78, 0x79, 0xfe, 0xbf, 0xb7, 0xf0, 0xcc, 0xf9, 0xb9, 0xf9, 0x06, 0x4c, 0x5c, 0xe5,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14, 0x17, 0xcd, 0xbf, 0xaa, 0x01, 0x00, 0x00,
}
