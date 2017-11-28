// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routing/v1alpha1/ingress_rule.proto

package istio_routing_v1alpha1_config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ingress rules are routing rules applied to the ingress proxy pool. The
// ingress proxes serve as the receiving edge proxy for the entire mesh, but
// can also be addressed from inside the mesh.  Each ingress rule defines a
// destination service and port. Rules that do not resolve to a service or a
// port in the mesh should be ignored.
//
// The routing rules for the destination service are applied at the ingress
// proxy. That means the routing rule match conditions are composed and its
// actions are enforced. The traffic splitting for the destination service is
// also effective.
//
// WARNING: This API is experimental and under active development
type IngressRule struct {
	// REQUIRED: Port on which the ingress proxy listens and applies the rule.
	Port int32 `protobuf:"varint,1,opt,name=port" json:"port,omitempty"`
	// Optional TLS secret path to apply server-side TLS context on the port.
	// It is up to the underlying secret store to interpret the path to the secret.
	TlsSecret string `protobuf:"bytes,2,opt,name=tls_secret,json=tlsSecret" json:"tls_secret,omitempty"`
	// RECOMMENDED. Precedence is used to disambiguate the order of application
	// of rules. A higher number takes priority. If not specified, the value is
	// assumed to be 0.  The order of application for rules with the same
	// precedence is unspecified.
	Precedence int32 `protobuf:"varint,3,opt,name=precedence" json:"precedence,omitempty"`
	// Match conditions to be satisfied for the ingress rule to be
	// activated.
	Match *MatchCondition `protobuf:"bytes,4,opt,name=match" json:"match,omitempty"`
	// REQUIRED: Destination uniquely identifies the destination service.
	//
	// *Note:* The ingress rule destination specification represents all version
	// of the service and therefore the IstioService's labels field MUST be empty.
	//
	Destination *IstioService `protobuf:"bytes,5,opt,name=destination" json:"destination,omitempty"`
	// REQUIRED: Destination port identifies a port on the destination service for routing.
	//
	// Types that are valid to be assigned to DestinationServicePort:
	//	*IngressRule_DestinationPort
	//	*IngressRule_DestinationPortName
	DestinationServicePort isIngressRule_DestinationServicePort `protobuf_oneof:"destination_service_port"`
}

func (m *IngressRule) Reset()                    { *m = IngressRule{} }
func (m *IngressRule) String() string            { return proto.CompactTextString(m) }
func (*IngressRule) ProtoMessage()               {}
func (*IngressRule) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type isIngressRule_DestinationServicePort interface {
	isIngressRule_DestinationServicePort()
}

type IngressRule_DestinationPort struct {
	DestinationPort int32 `protobuf:"varint,6,opt,name=destination_port,json=destinationPort,oneof"`
}
type IngressRule_DestinationPortName struct {
	DestinationPortName string `protobuf:"bytes,7,opt,name=destination_port_name,json=destinationPortName,oneof"`
}

func (*IngressRule_DestinationPort) isIngressRule_DestinationServicePort()     {}
func (*IngressRule_DestinationPortName) isIngressRule_DestinationServicePort() {}

func (m *IngressRule) GetDestinationServicePort() isIngressRule_DestinationServicePort {
	if m != nil {
		return m.DestinationServicePort
	}
	return nil
}

func (m *IngressRule) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *IngressRule) GetTlsSecret() string {
	if m != nil {
		return m.TlsSecret
	}
	return ""
}

func (m *IngressRule) GetPrecedence() int32 {
	if m != nil {
		return m.Precedence
	}
	return 0
}

func (m *IngressRule) GetMatch() *MatchCondition {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *IngressRule) GetDestination() *IstioService {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *IngressRule) GetDestinationPort() int32 {
	if x, ok := m.GetDestinationServicePort().(*IngressRule_DestinationPort); ok {
		return x.DestinationPort
	}
	return 0
}

func (m *IngressRule) GetDestinationPortName() string {
	if x, ok := m.GetDestinationServicePort().(*IngressRule_DestinationPortName); ok {
		return x.DestinationPortName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*IngressRule) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _IngressRule_OneofMarshaler, _IngressRule_OneofUnmarshaler, _IngressRule_OneofSizer, []interface{}{
		(*IngressRule_DestinationPort)(nil),
		(*IngressRule_DestinationPortName)(nil),
	}
}

func _IngressRule_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*IngressRule)
	// destination_service_port
	switch x := m.DestinationServicePort.(type) {
	case *IngressRule_DestinationPort:
		b.EncodeVarint(6<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.DestinationPort))
	case *IngressRule_DestinationPortName:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.DestinationPortName)
	case nil:
	default:
		return fmt.Errorf("IngressRule.DestinationServicePort has unexpected type %T", x)
	}
	return nil
}

func _IngressRule_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*IngressRule)
	switch tag {
	case 6: // destination_service_port.destination_port
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.DestinationServicePort = &IngressRule_DestinationPort{int32(x)}
		return true, err
	case 7: // destination_service_port.destination_port_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.DestinationServicePort = &IngressRule_DestinationPortName{x}
		return true, err
	default:
		return false, nil
	}
}

func _IngressRule_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*IngressRule)
	// destination_service_port
	switch x := m.DestinationServicePort.(type) {
	case *IngressRule_DestinationPort:
		n += proto.SizeVarint(6<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.DestinationPort))
	case *IngressRule_DestinationPortName:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.DestinationPortName)))
		n += len(x.DestinationPortName)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*IngressRule)(nil), "istio.routing.v1alpha1.config.IngressRule")
}

func init() { proto.RegisterFile("routing/v1alpha1/ingress_rule.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xbb, 0xb5, 0xad, 0x74, 0x7a, 0x50, 0x22, 0x42, 0x28, 0x54, 0xaa, 0x5e, 0x0a, 0xc5,
	0x94, 0xaa, 0x4f, 0x60, 0x2f, 0xed, 0xa1, 0x22, 0xe9, 0x03, 0x2c, 0x31, 0x1d, 0xb7, 0x81, 0xdd,
	0x64, 0x49, 0xb2, 0x7d, 0x71, 0x5f, 0x40, 0x92, 0x55, 0x08, 0x15, 0xf4, 0x96, 0xfc, 0xf3, 0x7f,
	0x3f, 0xff, 0x0c, 0xdc, 0x5b, 0xd3, 0x78, 0xa5, 0x8b, 0xc5, 0x71, 0x29, 0xca, 0xfa, 0x20, 0x96,
	0x0b, 0xa5, 0x0b, 0x8b, 0xce, 0xe5, 0xb6, 0x29, 0x91, 0xd5, 0xd6, 0x78, 0x43, 0x26, 0xca, 0x79,
	0x65, 0xd8, 0xb7, 0x95, 0xfd, 0x58, 0x99, 0x34, 0xfa, 0x43, 0x15, 0xe3, 0xdb, 0x5f, 0x19, 0x41,
	0xc0, 0x24, 0xe1, 0xee, 0xb3, 0x0b, 0xa3, 0x4d, 0x1b, 0xcc, 0x9b, 0x12, 0x09, 0x81, 0x5e, 0x6d,
	0xac, 0xa7, 0xd9, 0x34, 0x9b, 0xf5, 0x79, 0x7c, 0x93, 0x09, 0x80, 0x2f, 0x5d, 0xee, 0x50, 0x5a,
	0xf4, 0xb4, 0x3b, 0xcd, 0x66, 0x43, 0x3e, 0xf4, 0xa5, 0xdb, 0x45, 0x81, 0xdc, 0x00, 0xd4, 0x16,
	0x25, 0xee, 0x51, 0x4b, 0xa4, 0x67, 0x11, 0x4c, 0x14, 0xb2, 0x82, 0x7e, 0x25, 0xbc, 0x3c, 0xd0,
	0xde, 0x34, 0x9b, 0x8d, 0x1e, 0x1f, 0xd8, 0x9f, 0xa5, 0xd9, 0x36, 0x78, 0x57, 0x46, 0xef, 0x95,
	0x57, 0x46, 0xf3, 0x96, 0x25, 0x5b, 0x18, 0xed, 0xd1, 0x79, 0xa5, 0x45, 0x50, 0x69, 0x3f, 0x46,
	0xcd, 0xff, 0x89, 0xda, 0x84, 0xe9, 0x0e, 0xed, 0x51, 0x49, 0xe4, 0x29, 0x4f, 0xe6, 0x70, 0x99,
	0x7c, 0xf3, 0xb8, 0xf2, 0x20, 0x34, 0x5f, 0x77, 0xf8, 0x45, 0x32, 0x79, 0x0b, 0xfb, 0x3f, 0xc3,
	0xf5, 0xa9, 0x39, 0xd7, 0xa2, 0x42, 0x7a, 0x1e, 0x4e, 0xb1, 0xee, 0xf0, 0xab, 0x13, 0xe2, 0x55,
	0x54, 0xf8, 0x32, 0x06, 0x9a, 0x52, 0xae, 0xad, 0x11, 0xe9, 0xf7, 0x41, 0x3c, 0xfe, 0xd3, 0x57,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x71, 0xfe, 0x53, 0x1c, 0xe5, 0x01, 0x00, 0x00,
}