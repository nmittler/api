// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/v1/template/standard_types.proto

package istio_mixer_v1_template

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Value is used inside templates for fields that have dynamic types. The actual datatype
// of the field depends on the datatype of the expression used in the operator configuration.
type Value struct {
}

func (m *Value) Reset()                    { *m = Value{} }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptorStandardTypes, []int{0} }

// IPAddress is used inside templates for fields that are of ValueType "IP_ADDRESS"
type IPAddress struct {
}

func (m *IPAddress) Reset()                    { *m = IPAddress{} }
func (*IPAddress) ProtoMessage()               {}
func (*IPAddress) Descriptor() ([]byte, []int) { return fileDescriptorStandardTypes, []int{1} }

// Duration is used inside templates for fields that are of ValueType "DURATION"
type Duration struct {
}

func (m *Duration) Reset()                    { *m = Duration{} }
func (*Duration) ProtoMessage()               {}
func (*Duration) Descriptor() ([]byte, []int) { return fileDescriptorStandardTypes, []int{2} }

// TimeStamp is used inside templates for fields that are of ValueType "TIMESTAMP"
type TimeStamp struct {
}

func (m *TimeStamp) Reset()                    { *m = TimeStamp{} }
func (*TimeStamp) ProtoMessage()               {}
func (*TimeStamp) Descriptor() ([]byte, []int) { return fileDescriptorStandardTypes, []int{3} }

// DNSName is used inside templates for fields that are of ValueType "DNS_NAME"
type DNSName struct {
}

func (m *DNSName) Reset()                    { *m = DNSName{} }
func (*DNSName) ProtoMessage()               {}
func (*DNSName) Descriptor() ([]byte, []int) { return fileDescriptorStandardTypes, []int{4} }

// EmailAddress is used inside templates for fields that are of ValueType "EMAIL_ADDRESS"
// DO NOT USE !! Under Development
type EmailAddress struct {
}

func (m *EmailAddress) Reset()                    { *m = EmailAddress{} }
func (*EmailAddress) ProtoMessage()               {}
func (*EmailAddress) Descriptor() ([]byte, []int) { return fileDescriptorStandardTypes, []int{5} }

// Uri is used inside templates for fields that are of ValueType "URI"
// DO NOT USE ! Under Development
type Uri struct {
}

func (m *Uri) Reset()                    { *m = Uri{} }
func (*Uri) ProtoMessage()               {}
func (*Uri) Descriptor() ([]byte, []int) { return fileDescriptorStandardTypes, []int{6} }

func init() {
	proto.RegisterType((*Value)(nil), "istio.mixer.v1.template.Value")
	proto.RegisterType((*IPAddress)(nil), "istio.mixer.v1.template.IPAddress")
	proto.RegisterType((*Duration)(nil), "istio.mixer.v1.template.Duration")
	proto.RegisterType((*TimeStamp)(nil), "istio.mixer.v1.template.TimeStamp")
	proto.RegisterType((*DNSName)(nil), "istio.mixer.v1.template.DNSName")
	proto.RegisterType((*EmailAddress)(nil), "istio.mixer.v1.template.EmailAddress")
	proto.RegisterType((*Uri)(nil), "istio.mixer.v1.template.Uri")
}
func (this *Value) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Value)
	if !ok {
		that2, ok := that.(Value)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *IPAddress) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*IPAddress)
	if !ok {
		that2, ok := that.(IPAddress)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *Duration) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Duration)
	if !ok {
		that2, ok := that.(Duration)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *TimeStamp) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*TimeStamp)
	if !ok {
		that2, ok := that.(TimeStamp)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *DNSName) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*DNSName)
	if !ok {
		that2, ok := that.(DNSName)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *EmailAddress) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*EmailAddress)
	if !ok {
		that2, ok := that.(EmailAddress)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *Uri) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Uri)
	if !ok {
		that2, ok := that.(Uri)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *Value) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&istio_mixer_v1_template.Value{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *IPAddress) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&istio_mixer_v1_template.IPAddress{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Duration) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&istio_mixer_v1_template.Duration{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TimeStamp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&istio_mixer_v1_template.TimeStamp{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DNSName) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&istio_mixer_v1_template.DNSName{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *EmailAddress) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&istio_mixer_v1_template.EmailAddress{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Uri) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&istio_mixer_v1_template.Uri{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringStandardTypes(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Value) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Value) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *IPAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IPAddress) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *Duration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Duration) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *TimeStamp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimeStamp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *DNSName) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DNSName) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *EmailAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EmailAddress) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *Uri) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Uri) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintStandardTypes(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Value) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *IPAddress) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *Duration) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *TimeStamp) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *DNSName) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *EmailAddress) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *Uri) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovStandardTypes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStandardTypes(x uint64) (n int) {
	return sovStandardTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Value) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Value{`,
		`}`,
	}, "")
	return s
}
func (this *IPAddress) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IPAddress{`,
		`}`,
	}, "")
	return s
}
func (this *Duration) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Duration{`,
		`}`,
	}, "")
	return s
}
func (this *TimeStamp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TimeStamp{`,
		`}`,
	}, "")
	return s
}
func (this *DNSName) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DNSName{`,
		`}`,
	}, "")
	return s
}
func (this *EmailAddress) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EmailAddress{`,
		`}`,
	}, "")
	return s
}
func (this *Uri) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Uri{`,
		`}`,
	}, "")
	return s
}
func valueToStringStandardTypes(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Value) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStandardTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Value: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Value: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipStandardTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStandardTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IPAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStandardTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IPAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IPAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipStandardTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStandardTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Duration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStandardTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Duration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Duration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipStandardTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStandardTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TimeStamp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStandardTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TimeStamp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeStamp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipStandardTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStandardTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DNSName) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStandardTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DNSName: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DNSName: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipStandardTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStandardTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EmailAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStandardTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EmailAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmailAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipStandardTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStandardTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Uri) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStandardTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Uri: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Uri: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipStandardTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStandardTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStandardTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStandardTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStandardTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStandardTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthStandardTypes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStandardTypes
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipStandardTypes(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthStandardTypes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStandardTypes   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mixer/v1/template/standard_types.proto", fileDescriptorStandardTypes) }

var fileDescriptorStandardTypes = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcb, 0xcd, 0xac, 0x48,
	0x2d, 0xd2, 0x2f, 0x33, 0xd4, 0x2f, 0x49, 0xcd, 0x2d, 0xc8, 0x49, 0x2c, 0x49, 0xd5, 0x2f, 0x2e,
	0x49, 0xcc, 0x4b, 0x49, 0x2c, 0x4a, 0x89, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0xcf, 0x2c, 0x2e, 0xc9, 0xcc, 0xd7, 0x03, 0xab, 0xd6, 0x2b, 0x33, 0xd4,
	0x83, 0xa9, 0x56, 0x62, 0xe7, 0x62, 0x0d, 0x4b, 0xcc, 0x29, 0x4d, 0x55, 0xe2, 0xe6, 0xe2, 0xf4,
	0x0c, 0x70, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x56, 0xe2, 0xe2, 0xe2, 0x70, 0x29, 0x2d, 0x4a,
	0x2c, 0xc9, 0xcc, 0xcf, 0x03, 0x49, 0x84, 0x64, 0xe6, 0xa6, 0x06, 0x97, 0x24, 0xe6, 0x16, 0x28,
	0x71, 0x72, 0xb1, 0xbb, 0xf8, 0x05, 0xfb, 0x25, 0xe6, 0xa6, 0x2a, 0xf1, 0x71, 0xf1, 0xb8, 0xe6,
	0x26, 0x66, 0xe6, 0xc0, 0xf4, 0xb0, 0x72, 0x31, 0x87, 0x16, 0x65, 0x3a, 0xe9, 0x5c, 0x78, 0x28,
	0xc7, 0x70, 0xe3, 0xa1, 0x1c, 0xc3, 0x87, 0x87, 0x72, 0x8c, 0x0d, 0x8f, 0xe4, 0x18, 0x57, 0x3c,
	0x92, 0x63, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x5f,
	0x3c, 0x92, 0x63, 0xf8, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x24, 0x36, 0xb0, 0xf3,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xbf, 0x05, 0x83, 0xc8, 0x00, 0x00, 0x00,
}