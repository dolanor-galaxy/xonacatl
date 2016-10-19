// Code generated by protoc-gen-go.
// source: mapnik_vector/vector_tile.proto
// DO NOT EDIT!

/*
Package mapnik_vector is a generated protocol buffer package.

It is generated from these files:
	mapnik_vector/vector_tile.proto

It has these top-level messages:
	Tile
*/
package mapnik_vector

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

type Tile_GeomType int32

const (
	Tile_Unknown    Tile_GeomType = 0
	Tile_Point      Tile_GeomType = 1
	Tile_LineString Tile_GeomType = 2
	Tile_Polygon    Tile_GeomType = 3
)

var Tile_GeomType_name = map[int32]string{
	0: "Unknown",
	1: "Point",
	2: "LineString",
	3: "Polygon",
}
var Tile_GeomType_value = map[string]int32{
	"Unknown":    0,
	"Point":      1,
	"LineString": 2,
	"Polygon":    3,
}

func (x Tile_GeomType) Enum() *Tile_GeomType {
	p := new(Tile_GeomType)
	*p = x
	return p
}
func (x Tile_GeomType) String() string {
	return proto.EnumName(Tile_GeomType_name, int32(x))
}
func (x *Tile_GeomType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Tile_GeomType_value, data, "Tile_GeomType")
	if err != nil {
		return err
	}
	*x = Tile_GeomType(value)
	return nil
}
func (Tile_GeomType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Tile struct {
	Layers                       []*TileLayer `protobuf:"bytes,3,rep,name=layers" json:"layers,omitempty"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *Tile) Reset()                    { *m = Tile{} }
func (m *Tile) String() string            { return proto.CompactTextString(m) }
func (*Tile) ProtoMessage()               {}
func (*Tile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

var extRange_Tile = []proto.ExtensionRange{
	{16, 8191},
}

func (*Tile) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Tile
}

func (m *Tile) GetLayers() []*TileLayer {
	if m != nil {
		return m.Layers
	}
	return nil
}

// Variant type encoding
type TileValue struct {
	// Exactly one of these values may be present in a valid message
	StringValue                  *string  `protobuf:"bytes,1,opt,name=string_value" json:"string_value,omitempty"`
	FloatValue                   *float32 `protobuf:"fixed32,2,opt,name=float_value" json:"float_value,omitempty"`
	DoubleValue                  *float64 `protobuf:"fixed64,3,opt,name=double_value" json:"double_value,omitempty"`
	IntValue                     *int64   `protobuf:"varint,4,opt,name=int_value" json:"int_value,omitempty"`
	UintValue                    *uint64  `protobuf:"varint,5,opt,name=uint_value" json:"uint_value,omitempty"`
	SintValue                    *int64   `protobuf:"zigzag64,6,opt,name=sint_value" json:"sint_value,omitempty"`
	BoolValue                    *bool    `protobuf:"varint,7,opt,name=bool_value" json:"bool_value,omitempty"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *TileValue) Reset()                    { *m = TileValue{} }
func (m *TileValue) String() string            { return proto.CompactTextString(m) }
func (*TileValue) ProtoMessage()               {}
func (*TileValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

var extRange_TileValue = []proto.ExtensionRange{
	{8, 536870911},
}

func (*TileValue) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_TileValue
}

func (m *TileValue) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *TileValue) GetFloatValue() float32 {
	if m != nil && m.FloatValue != nil {
		return *m.FloatValue
	}
	return 0
}

func (m *TileValue) GetDoubleValue() float64 {
	if m != nil && m.DoubleValue != nil {
		return *m.DoubleValue
	}
	return 0
}

func (m *TileValue) GetIntValue() int64 {
	if m != nil && m.IntValue != nil {
		return *m.IntValue
	}
	return 0
}

func (m *TileValue) GetUintValue() uint64 {
	if m != nil && m.UintValue != nil {
		return *m.UintValue
	}
	return 0
}

func (m *TileValue) GetSintValue() int64 {
	if m != nil && m.SintValue != nil {
		return *m.SintValue
	}
	return 0
}

func (m *TileValue) GetBoolValue() bool {
	if m != nil && m.BoolValue != nil {
		return *m.BoolValue
	}
	return false
}

type TileFeature struct {
	Id *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Tags of this feature. Even numbered values refer to the nth
	// value in the keys list on the tile message, odd numbered
	// values refer to the nth value in the values list on the tile
	// message.
	Tags []uint32 `protobuf:"varint,2,rep,packed,name=tags" json:"tags,omitempty"`
	// The type of geometry stored in this feature.
	Type *Tile_GeomType `protobuf:"varint,3,opt,name=type,enum=mapnik.vector.Tile_GeomType,def=0" json:"type,omitempty"`
	// Contains a stream of commands and parameters (vertices). The
	// repeat count is shifted to the left by 3 bits. This means
	// that the command has 3 bits (0-7). The repeat count
	// indicates how often this command is to be repeated. Defined
	// commands are:
	// - MoveTo:    1   (2 parameters follow)
	// - LineTo:    2   (2 parameters follow)
	// - ClosePath: 7   (no parameters follow)
	//
	// Ex.: MoveTo(3, 6), LineTo(8, 12), LineTo(20, 34), ClosePath
	// Encoded as: [ 9 3 6 18 5 6 12 22 15 ]
	//                                  == command type 7 (ClosePath), length 1
	//                             ===== relative LineTo(+12, +22) == LineTo(20, 34)
	//                         === relative LineTo(+5, +6) == LineTo(8, 12)
	//                      == [00010 010] = command type 2 (LineTo), length 2
	//                  === relative MoveTo(+3, +6)
	//              == [00001 001] = command type 1 (MoveTo), length 1
	// Commands are encoded as uint32 varints, vertex parameters are
	// encoded as sint32 varints (zigzag). Vertex parameters are
	// also encoded as deltas to the previous position. The original
	// position is (0,0)
	Geometry         []uint32 `protobuf:"varint,4,rep,packed,name=geometry" json:"geometry,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TileFeature) Reset()                    { *m = TileFeature{} }
func (m *TileFeature) String() string            { return proto.CompactTextString(m) }
func (*TileFeature) ProtoMessage()               {}
func (*TileFeature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

const Default_TileFeature_Type Tile_GeomType = Tile_Unknown

func (m *TileFeature) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TileFeature) GetTags() []uint32 {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *TileFeature) GetType() Tile_GeomType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_TileFeature_Type
}

func (m *TileFeature) GetGeometry() []uint32 {
	if m != nil {
		return m.Geometry
	}
	return nil
}

type TileLayer struct {
	// Any compliant implementation must first read the version
	// number encoded in this message and choose the correct
	// implementation for this version number before proceeding to
	// decode other parts of this message.
	Version *uint32 `protobuf:"varint,15,req,name=version,def=1" json:"version,omitempty"`
	Name    *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// The actual features in this tile.
	Features []*TileFeature `protobuf:"bytes,2,rep,name=features" json:"features,omitempty"`
	// Dictionary encoding for keys
	Keys []string `protobuf:"bytes,3,rep,name=keys" json:"keys,omitempty"`
	// Dictionary encoding for values
	Values []*TileValue `protobuf:"bytes,4,rep,name=values" json:"values,omitempty"`
	// The bounding box in this tile spans from 0..4095 units
	Extent                       *uint32 `protobuf:"varint,5,opt,name=extent,def=4096" json:"extent,omitempty"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *TileLayer) Reset()                    { *m = TileLayer{} }
func (m *TileLayer) String() string            { return proto.CompactTextString(m) }
func (*TileLayer) ProtoMessage()               {}
func (*TileLayer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

var extRange_TileLayer = []proto.ExtensionRange{
	{16, 536870911},
}

func (*TileLayer) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_TileLayer
}

const Default_TileLayer_Version uint32 = 1
const Default_TileLayer_Extent uint32 = 4096

func (m *TileLayer) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return Default_TileLayer_Version
}

func (m *TileLayer) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TileLayer) GetFeatures() []*TileFeature {
	if m != nil {
		return m.Features
	}
	return nil
}

func (m *TileLayer) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *TileLayer) GetValues() []*TileValue {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *TileLayer) GetExtent() uint32 {
	if m != nil && m.Extent != nil {
		return *m.Extent
	}
	return Default_TileLayer_Extent
}

func init() {
	proto.RegisterType((*Tile)(nil), "mapnik.vector.tile")
	proto.RegisterType((*TileValue)(nil), "mapnik.vector.tile.value")
	proto.RegisterType((*TileFeature)(nil), "mapnik.vector.tile.feature")
	proto.RegisterType((*TileLayer)(nil), "mapnik.vector.tile.layer")
	proto.RegisterEnum("mapnik.vector.Tile_GeomType", Tile_GeomType_name, Tile_GeomType_value)
}

func init() { proto.RegisterFile("mapnik_vector/vector_tile.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x51, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x8e, 0xd3, 0x3a, 0xd3, 0x76, 0x09, 0xc3, 0x1e, 0x42, 0x41, 0xc2, 0xda, 0x93, 0xa9,
	0x44, 0x81, 0x82, 0x90, 0xe8, 0x05, 0xb4, 0x17, 0x38, 0x70, 0x58, 0x09, 0x38, 0x57, 0x59, 0x76,
	0xb6, 0x8a, 0x9a, 0xda, 0x55, 0xe2, 0x16, 0x72, 0xeb, 0xab, 0xf0, 0x08, 0x3c, 0x04, 0xef, 0x85,
	0xec, 0x24, 0x5a, 0x21, 0x95, 0x93, 0xe5, 0x6f, 0x3e, 0xcd, 0x7c, 0x3f, 0xf0, 0x74, 0x9b, 0xef,
	0x4c, 0xb1, 0x59, 0x1d, 0xe8, 0xbb, 0xb3, 0xd5, 0x8b, 0xf6, 0x59, 0xb9, 0xa2, 0xa4, 0xf9, 0xae,
	0xb2, 0xce, 0xe2, 0xa4, 0x25, 0xcc, 0xdb, 0xc9, 0xc5, 0x1f, 0x01, 0xc2, 0x4f, 0xf1, 0x19, 0x0c,
	0xca, 0xbc, 0xa1, 0xaa, 0xce, 0x22, 0x15, 0xe9, 0xd1, 0xe2, 0xd1, 0xfc, 0x1f, 0xe2, 0x3c, 0xac,
	0x08, 0x8c, 0xe9, 0x2f, 0x06, 0xf1, 0x21, 0x2f, 0xf7, 0x84, 0xe7, 0x30, 0xae, 0x5d, 0x55, 0x98,
	0xf5, 0x2a, 0xfc, 0x33, 0xa6, 0x98, 0x4e, 0xf0, 0x21, 0x8c, 0x6e, 0x4b, 0x9b, 0xbb, 0x0e, 0xe4,
	0x8a, 0x69, 0xee, 0xa9, 0x37, 0x76, 0x7f, 0x5d, 0x52, 0x87, 0x46, 0x8a, 0x69, 0x86, 0x0f, 0x20,
	0x29, 0x4c, 0x4f, 0x14, 0x8a, 0xe9, 0x08, 0x11, 0x60, 0x7f, 0x87, 0xc5, 0x8a, 0x69, 0xe1, 0xb1,
	0xfa, 0x0e, 0x1b, 0x28, 0xa6, 0xd1, 0x63, 0xd7, 0xd6, 0x96, 0x1d, 0x36, 0x54, 0x4c, 0xcb, 0x99,
	0x94, 0x32, 0x3d, 0x1e, 0x8f, 0x47, 0x3e, 0xad, 0x60, 0x78, 0x4b, 0xb9, 0xdb, 0x57, 0x84, 0x00,
	0xbc, 0xb8, 0x09, 0xd2, 0x04, 0xa6, 0x20, 0x5c, 0xbe, 0xae, 0x33, 0xae, 0x22, 0x3d, 0xb9, 0xe4,
	0x29, 0xc3, 0xd7, 0x20, 0x5c, 0xb3, 0x6b, 0xf5, 0x9c, 0x2d, 0x9e, 0x9c, 0x72, 0xfd, 0x91, 0xec,
	0xf6, 0x6b, 0xb3, 0xa3, 0xe5, 0xf0, 0x9b, 0xd9, 0x18, 0xfb, 0xc3, 0xe0, 0x39, 0xc8, 0x35, 0xd9,
	0x2d, 0xb9, 0xaa, 0xc9, 0x44, 0xbf, 0x6a, 0xfa, 0x9b, 0x41, 0x1c, 0x12, 0x42, 0x84, 0xe1, 0x81,
	0xaa, 0xba, 0xb0, 0x26, 0xbb, 0xaf, 0xb8, 0x9e, 0x2c, 0xd9, 0x2b, 0x1c, 0x83, 0x30, 0xf9, 0xd6,
	0x67, 0xc4, 0x75, 0x82, 0xcf, 0x41, 0x76, 0xfa, 0x5a, 0x31, 0xa3, 0xc5, 0xe3, 0x53, 0xa7, 0x7b,
	0x0f, 0x63, 0x10, 0x1b, 0x6a, 0xda, 0x6e, 0x12, 0xdf, 0x55, 0x70, 0x5d, 0x87, 0xe3, 0xff, 0xe9,
	0xaa, 0x6f, 0x68, 0x40, 0x3f, 0x1d, 0x19, 0x17, 0x92, 0x9c, 0x2c, 0xc5, 0x9b, 0x97, 0xef, 0xde,
	0xce, 0xa4, 0x4c, 0xdb, 0x9c, 0x2e, 0xde, 0x83, 0xec, 0xed, 0xe1, 0x08, 0x7a, 0x83, 0xe9, 0x3d,
	0x4c, 0x20, 0xbe, 0xb2, 0x85, 0x71, 0x29, 0xc3, 0x33, 0x80, 0xcf, 0x85, 0xa1, 0x2f, 0xa1, 0xe9,
	0x94, 0x7b, 0xde, 0x95, 0x2d, 0x9b, 0xb5, 0x35, 0x69, 0x34, 0x8b, 0xfd, 0xaa, 0x0f, 0x97, 0xfc,
	0x53, 0xf4, 0x37, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x13, 0xd2, 0x8b, 0x7c, 0x02, 0x00, 0x00,
}