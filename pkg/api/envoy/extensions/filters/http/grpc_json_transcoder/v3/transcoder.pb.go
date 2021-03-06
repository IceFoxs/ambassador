// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/extensions/filters/http/grpc_json_transcoder/v3/transcoder.proto

package envoy_extensions_filters_http_grpc_json_transcoder_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// [#next-free-field: 10]
type GrpcJsonTranscoder struct {
	// Types that are valid to be assigned to DescriptorSet:
	//	*GrpcJsonTranscoder_ProtoDescriptor
	//	*GrpcJsonTranscoder_ProtoDescriptorBin
	DescriptorSet isGrpcJsonTranscoder_DescriptorSet `protobuf_oneof:"descriptor_set"`
	// A list of strings that
	// supplies the fully qualified service names (i.e. "package_name.service_name") that
	// the transcoder will translate. If the service name doesn't exist in ``proto_descriptor``,
	// Envoy will fail at startup. The ``proto_descriptor`` may contain more services than
	// the service names specified here, but they won't be translated.
	Services []string `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty"`
	// Control options for response JSON. These options are passed directly to
	// `JsonPrintOptions <https://developers.google.com/protocol-buffers/docs/reference/cpp/
	// google.protobuf.util.json_util#JsonPrintOptions>`_.
	PrintOptions *GrpcJsonTranscoder_PrintOptions `protobuf:"bytes,3,opt,name=print_options,json=printOptions,proto3" json:"print_options,omitempty"`
	// Whether to keep the incoming request route after the outgoing headers have been transformed to
	// the match the upstream gRPC service. Note: This means that routes for gRPC services that are
	// not transcoded cannot be used in combination with *match_incoming_request_route*.
	MatchIncomingRequestRoute bool `protobuf:"varint,5,opt,name=match_incoming_request_route,json=matchIncomingRequestRoute,proto3" json:"match_incoming_request_route,omitempty"`
	// A list of query parameters to be ignored for transcoding method mapping.
	// By default, the transcoder filter will not transcode a request if there are any
	// unknown/invalid query parameters.
	//
	// Example :
	//
	// .. code-block:: proto
	//
	//     service Bookstore {
	//       rpc GetShelf(GetShelfRequest) returns (Shelf) {
	//         option (google.api.http) = {
	//           get: "/shelves/{shelf}"
	//         };
	//       }
	//     }
	//
	//     message GetShelfRequest {
	//       int64 shelf = 1;
	//     }
	//
	//     message Shelf {}
	//
	// The request ``/shelves/100?foo=bar`` will not be mapped to ``GetShelf``` because variable
	// binding for ``foo`` is not defined. Adding ``foo`` to ``ignored_query_parameters`` will allow
	// the same request to be mapped to ``GetShelf``.
	IgnoredQueryParameters []string `protobuf:"bytes,6,rep,name=ignored_query_parameters,json=ignoredQueryParameters,proto3" json:"ignored_query_parameters,omitempty"`
	// Whether to route methods without the ``google.api.http`` option.
	//
	// Example :
	//
	// .. code-block:: proto
	//
	//     package bookstore;
	//
	//     service Bookstore {
	//       rpc GetShelf(GetShelfRequest) returns (Shelf) {}
	//     }
	//
	//     message GetShelfRequest {
	//       int64 shelf = 1;
	//     }
	//
	//     message Shelf {}
	//
	// The client could ``post`` a json body ``{"shelf": 1234}`` with the path of
	// ``/bookstore.Bookstore/GetShelfRequest`` to call ``GetShelfRequest``.
	AutoMapping bool `protobuf:"varint,7,opt,name=auto_mapping,json=autoMapping,proto3" json:"auto_mapping,omitempty"`
	// Whether to ignore query parameters that cannot be mapped to a corresponding
	// protobuf field. Use this if you cannot control the query parameters and do
	// not know them beforehand. Otherwise use ``ignored_query_parameters``.
	// Defaults to false.
	IgnoreUnknownQueryParameters bool `protobuf:"varint,8,opt,name=ignore_unknown_query_parameters,json=ignoreUnknownQueryParameters,proto3" json:"ignore_unknown_query_parameters,omitempty"`
	// Whether to convert gRPC status headers to JSON.
	// When trailer indicates a gRPC error and there was no HTTP body, take ``google.rpc.Status``
	// from the ``grpc-status-details-bin`` header and use it as JSON body.
	// If there was no such header, make ``google.rpc.Status`` out of the ``grpc-status`` and
	// ``grpc-message`` headers.
	// The error details types must be present in the ``proto_descriptor``.
	//
	// For example, if an upstream server replies with headers:
	//
	// .. code-block:: none
	//
	//     grpc-status: 5
	//     grpc-status-details-bin:
	//         CAUaMwoqdHlwZS5nb29nbGVhcGlzLmNvbS9nb29nbGUucnBjLlJlcXVlc3RJbmZvEgUKA3ItMQ
	//
	// The ``grpc-status-details-bin`` header contains a base64-encoded protobuf message
	// ``google.rpc.Status``. It will be transcoded into:
	//
	// .. code-block:: none
	//
	//     HTTP/1.1 404 Not Found
	//     content-type: application/json
	//
	//     {"code":5,"details":[{"@type":"type.googleapis.com/google.rpc.RequestInfo","requestId":"r-1"}]}
	//
	//  In order to transcode the message, the ``google.rpc.RequestInfo`` type from
	//  the ``google/rpc/error_details.proto`` should be included in the configured
	//  :ref:`proto descriptor set <config_grpc_json_generate_proto_descriptor_set>`.
	ConvertGrpcStatus    bool     `protobuf:"varint,9,opt,name=convert_grpc_status,json=convertGrpcStatus,proto3" json:"convert_grpc_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcJsonTranscoder) Reset()         { *m = GrpcJsonTranscoder{} }
func (m *GrpcJsonTranscoder) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder) ProtoMessage()    {}
func (*GrpcJsonTranscoder) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dffdaf96e302174, []int{0}
}
func (m *GrpcJsonTranscoder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GrpcJsonTranscoder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GrpcJsonTranscoder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GrpcJsonTranscoder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder.Merge(m, src)
}
func (m *GrpcJsonTranscoder) XXX_Size() int {
	return m.Size()
}
func (m *GrpcJsonTranscoder) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder proto.InternalMessageInfo

type isGrpcJsonTranscoder_DescriptorSet interface {
	isGrpcJsonTranscoder_DescriptorSet()
	MarshalTo([]byte) (int, error)
	Size() int
}

type GrpcJsonTranscoder_ProtoDescriptor struct {
	ProtoDescriptor string `protobuf:"bytes,1,opt,name=proto_descriptor,json=protoDescriptor,proto3,oneof" json:"proto_descriptor,omitempty"`
}
type GrpcJsonTranscoder_ProtoDescriptorBin struct {
	ProtoDescriptorBin []byte `protobuf:"bytes,4,opt,name=proto_descriptor_bin,json=protoDescriptorBin,proto3,oneof" json:"proto_descriptor_bin,omitempty"`
}

func (*GrpcJsonTranscoder_ProtoDescriptor) isGrpcJsonTranscoder_DescriptorSet()    {}
func (*GrpcJsonTranscoder_ProtoDescriptorBin) isGrpcJsonTranscoder_DescriptorSet() {}

func (m *GrpcJsonTranscoder) GetDescriptorSet() isGrpcJsonTranscoder_DescriptorSet {
	if m != nil {
		return m.DescriptorSet
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetProtoDescriptor() string {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptor); ok {
		return x.ProtoDescriptor
	}
	return ""
}

func (m *GrpcJsonTranscoder) GetProtoDescriptorBin() []byte {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptorBin); ok {
		return x.ProtoDescriptorBin
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetPrintOptions() *GrpcJsonTranscoder_PrintOptions {
	if m != nil {
		return m.PrintOptions
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetMatchIncomingRequestRoute() bool {
	if m != nil {
		return m.MatchIncomingRequestRoute
	}
	return false
}

func (m *GrpcJsonTranscoder) GetIgnoredQueryParameters() []string {
	if m != nil {
		return m.IgnoredQueryParameters
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetAutoMapping() bool {
	if m != nil {
		return m.AutoMapping
	}
	return false
}

func (m *GrpcJsonTranscoder) GetIgnoreUnknownQueryParameters() bool {
	if m != nil {
		return m.IgnoreUnknownQueryParameters
	}
	return false
}

func (m *GrpcJsonTranscoder) GetConvertGrpcStatus() bool {
	if m != nil {
		return m.ConvertGrpcStatus
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GrpcJsonTranscoder) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GrpcJsonTranscoder_ProtoDescriptor)(nil),
		(*GrpcJsonTranscoder_ProtoDescriptorBin)(nil),
	}
}

type GrpcJsonTranscoder_PrintOptions struct {
	// Whether to add spaces, line breaks and indentation to make the JSON
	// output easy to read. Defaults to false.
	AddWhitespace bool `protobuf:"varint,1,opt,name=add_whitespace,json=addWhitespace,proto3" json:"add_whitespace,omitempty"`
	// Whether to always print primitive fields. By default primitive
	// fields with default values will be omitted in JSON output. For
	// example, an int32 field set to 0 will be omitted. Setting this flag to
	// true will override the default behavior and print primitive fields
	// regardless of their values. Defaults to false.
	AlwaysPrintPrimitiveFields bool `protobuf:"varint,2,opt,name=always_print_primitive_fields,json=alwaysPrintPrimitiveFields,proto3" json:"always_print_primitive_fields,omitempty"`
	// Whether to always print enums as ints. By default they are rendered
	// as strings. Defaults to false.
	AlwaysPrintEnumsAsInts bool `protobuf:"varint,3,opt,name=always_print_enums_as_ints,json=alwaysPrintEnumsAsInts,proto3" json:"always_print_enums_as_ints,omitempty"`
	// Whether to preserve proto field names. By default protobuf will
	// generate JSON field names using the ``json_name`` option, or lower camel case,
	// in that order. Setting this flag will preserve the original field names. Defaults to false.
	PreserveProtoFieldNames bool     `protobuf:"varint,4,opt,name=preserve_proto_field_names,json=preserveProtoFieldNames,proto3" json:"preserve_proto_field_names,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *GrpcJsonTranscoder_PrintOptions) Reset()         { *m = GrpcJsonTranscoder_PrintOptions{} }
func (m *GrpcJsonTranscoder_PrintOptions) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder_PrintOptions) ProtoMessage()    {}
func (*GrpcJsonTranscoder_PrintOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dffdaf96e302174, []int{0, 0}
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Merge(m, src)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Size() int {
	return m.Size()
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder_PrintOptions proto.InternalMessageInfo

func (m *GrpcJsonTranscoder_PrintOptions) GetAddWhitespace() bool {
	if m != nil {
		return m.AddWhitespace
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintPrimitiveFields() bool {
	if m != nil {
		return m.AlwaysPrintPrimitiveFields
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintEnumsAsInts() bool {
	if m != nil {
		return m.AlwaysPrintEnumsAsInts
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetPreserveProtoFieldNames() bool {
	if m != nil {
		return m.PreserveProtoFieldNames
	}
	return false
}

func init() {
	proto.RegisterType((*GrpcJsonTranscoder)(nil), "envoy.extensions.filters.http.grpc_json_transcoder.v3.GrpcJsonTranscoder")
	proto.RegisterType((*GrpcJsonTranscoder_PrintOptions)(nil), "envoy.extensions.filters.http.grpc_json_transcoder.v3.GrpcJsonTranscoder.PrintOptions")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/grpc_json_transcoder/v3/transcoder.proto", fileDescriptor_6dffdaf96e302174)
}

var fileDescriptor_6dffdaf96e302174 = []byte{
	// 648 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0x4e, 0x14, 0x31,
	0x1c, 0x67, 0x40, 0x60, 0x28, 0x0b, 0x68, 0x55, 0x18, 0x37, 0xb8, 0x2e, 0x46, 0x93, 0x4d, 0x4c,
	0x66, 0x12, 0x88, 0x89, 0xc2, 0x41, 0x19, 0x05, 0xc1, 0x04, 0x5d, 0xc7, 0xaf, 0x63, 0x53, 0x66,
	0xca, 0x52, 0xdd, 0x69, 0x4b, 0xdb, 0x19, 0xd8, 0xf8, 0x02, 0x9e, 0x3d, 0xfa, 0x34, 0x5e, 0x4c,
	0xf4, 0xe6, 0x23, 0x18, 0x9e, 0xc2, 0x78, 0x32, 0x6d, 0x87, 0x65, 0x05, 0x13, 0x93, 0xbd, 0x6d,
	0xfb, 0xfb, 0xda, 0xfe, 0xe7, 0xf7, 0x07, 0x9b, 0x84, 0x95, 0xbc, 0x17, 0x91, 0x23, 0x4d, 0x98,
	0xa2, 0x9c, 0xa9, 0x68, 0x8f, 0x76, 0x35, 0x91, 0x2a, 0xda, 0xd7, 0x5a, 0x44, 0x1d, 0x29, 0x52,
	0xf4, 0x4e, 0x71, 0x86, 0xb4, 0xc4, 0x4c, 0xa5, 0x3c, 0x23, 0x32, 0x2a, 0x57, 0xa2, 0xd3, 0x53,
	0x28, 0x24, 0xd7, 0x1c, 0xde, 0xb5, 0x3e, 0xe1, 0xa9, 0x4f, 0x58, 0xf9, 0x84, 0xc6, 0x27, 0xfc,
	0x97, 0x4f, 0x58, 0xae, 0xd4, 0x97, 0x8a, 0x4c, 0xe0, 0x08, 0x33, 0xc6, 0x35, 0xd6, 0x36, 0xbe,
	0x24, 0xd2, 0xe8, 0x29, 0xeb, 0x38, 0xe7, 0xfa, 0x42, 0x89, 0xbb, 0x34, 0xc3, 0x9a, 0x44, 0x27,
	0x3f, 0x1c, 0x70, 0xf3, 0xfb, 0x24, 0x80, 0x4f, 0xa4, 0x48, 0x9f, 0x2a, 0xce, 0x5e, 0xf5, 0x5d,
	0xe1, 0x1d, 0x70, 0xd1, 0xe2, 0x28, 0x23, 0x2a, 0x95, 0x54, 0x68, 0x2e, 0x03, 0xaf, 0xe9, 0xb5,
	0xa6, 0xb6, 0x46, 0x92, 0x39, 0x8b, 0x3c, 0xee, 0x03, 0x70, 0x19, 0x5c, 0x39, 0x4b, 0x46, 0xbb,
	0x94, 0x05, 0x17, 0x9a, 0x5e, 0xab, 0xb6, 0x35, 0x92, 0xc0, 0x33, 0x82, 0x98, 0x32, 0x78, 0x0b,
	0xf8, 0x8a, 0xc8, 0x92, 0xa6, 0x44, 0x05, 0xa3, 0xcd, 0xb1, 0xd6, 0x54, 0xec, 0xff, 0x8e, 0xc7,
	0x3f, 0x79, 0xa3, 0xbe, 0x97, 0xf4, 0x11, 0xf8, 0x01, 0xcc, 0x08, 0x49, 0x99, 0x46, 0x5c, 0xd8,
	0x87, 0x05, 0x63, 0x4d, 0xaf, 0x35, 0xbd, 0xfc, 0x26, 0x1c, 0x6a, 0x50, 0xe1, 0xf9, 0x87, 0x86,
	0x6d, 0x63, 0xff, 0xdc, 0xb9, 0x27, 0x35, 0x31, 0x70, 0x82, 0x0f, 0xc0, 0x62, 0x8e, 0x75, 0xba,
	0x8f, 0x28, 0x4b, 0x79, 0x4e, 0x59, 0x07, 0x49, 0x72, 0x50, 0x10, 0xa5, 0x91, 0xe4, 0x85, 0x26,
	0xc1, 0x78, 0xd3, 0x6b, 0xf9, 0xc9, 0x35, 0xcb, 0xd9, 0xae, 0x28, 0x89, 0x63, 0x24, 0x86, 0x00,
	0xef, 0x81, 0x80, 0x76, 0x18, 0x97, 0x24, 0x43, 0x07, 0x05, 0x91, 0x3d, 0x24, 0xb0, 0xc4, 0x39,
	0x31, 0xff, 0x33, 0x98, 0x30, 0x6f, 0x4e, 0xe6, 0x2b, 0xfc, 0x85, 0x81, 0xdb, 0x7d, 0x14, 0x2e,
	0x81, 0x1a, 0x2e, 0x34, 0x47, 0x39, 0x16, 0x82, 0xb2, 0x4e, 0x30, 0x69, 0xa3, 0xa6, 0xcd, 0xdd,
	0x8e, 0xbb, 0x82, 0x1b, 0xe0, 0x86, 0x13, 0xa3, 0x82, 0xbd, 0x67, 0xfc, 0x90, 0x9d, 0xcf, 0xf0,
	0xad, 0x6a, 0xd1, 0xd1, 0x5e, 0x3b, 0xd6, 0xd9, 0xa4, 0x10, 0x5c, 0x4e, 0x39, 0x2b, 0x89, 0xd4,
	0xc8, 0x4e, 0x4d, 0x69, 0xac, 0x0b, 0x15, 0x4c, 0x59, 0xe9, 0xa5, 0x0a, 0x32, 0x73, 0x7b, 0x69,
	0x81, 0xfa, 0x97, 0x51, 0x50, 0x1b, 0x9c, 0x19, 0xbc, 0x0d, 0x66, 0x71, 0x96, 0xa1, 0xc3, 0x7d,
	0xaa, 0x89, 0x12, 0x38, 0x25, 0xb6, 0x27, 0x7e, 0x32, 0x83, 0xb3, 0xec, 0x6d, 0xff, 0x12, 0xae,
	0x83, 0xeb, 0xb8, 0x7b, 0x88, 0x7b, 0x0a, 0xb9, 0x0f, 0x2a, 0x24, 0xcd, 0xa9, 0xa6, 0x25, 0x41,
	0x7b, 0x94, 0x74, 0x33, 0x53, 0x02, 0xa3, 0xaa, 0x3b, 0x92, 0x4d, 0x68, 0x9f, 0x50, 0x36, 0x2d,
	0x03, 0xae, 0x82, 0xfa, 0x5f, 0x16, 0x84, 0x15, 0xb9, 0x42, 0x58, 0x21, 0xca, 0xb4, 0x6b, 0x86,
	0x9f, 0xcc, 0x0f, 0xe8, 0x37, 0x0c, 0xbe, 0xae, 0xb6, 0x99, 0x56, 0x70, 0x0d, 0xd4, 0x85, 0x24,
	0xa6, 0x57, 0x04, 0xb9, 0xae, 0xda, 0x58, 0xc4, 0x70, 0x4e, 0x94, 0x2d, 0xaa, 0x9f, 0x2c, 0x9c,
	0x30, 0xda, 0x86, 0x60, 0x43, 0x9f, 0x19, 0x78, 0x75, 0xe7, 0xf3, 0xd7, 0x8f, 0x8d, 0xad, 0x6a,
	0xcb, 0xc3, 0x94, 0xb3, 0x3d, 0xda, 0xa9, 0x0a, 0xe7, 0xfa, 0x36, 0xd8, 0xb2, 0xe5, 0xff, 0xb5,
	0x6c, 0xf5, 0xa1, 0xb1, 0x5b, 0x03, 0xf7, 0x87, 0xb6, 0x8b, 0xaf, 0x82, 0xd9, 0x81, 0x55, 0x53,
	0x44, 0xc3, 0xb1, 0x5f, 0xb1, 0x17, 0x67, 0xdf, 0x8e, 0x1b, 0xde, 0x8f, 0xe3, 0x86, 0xf7, 0xf3,
	0xb8, 0xe1, 0x81, 0x47, 0x94, 0xbb, 0x35, 0x11, 0x92, 0x1f, 0xf5, 0x86, 0xdb, 0x98, 0x78, 0xee,
	0x34, 0xd5, 0x4e, 0xa5, 0xed, 0xed, 0x4e, 0xd8, 0xf9, 0xad, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff,
	0x04, 0x54, 0x32, 0x94, 0xf5, 0x04, 0x00, 0x00,
}

func (m *GrpcJsonTranscoder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrpcJsonTranscoder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GrpcJsonTranscoder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ConvertGrpcStatus {
		i--
		if m.ConvertGrpcStatus {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.IgnoreUnknownQueryParameters {
		i--
		if m.IgnoreUnknownQueryParameters {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.AutoMapping {
		i--
		if m.AutoMapping {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.IgnoredQueryParameters) > 0 {
		for iNdEx := len(m.IgnoredQueryParameters) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.IgnoredQueryParameters[iNdEx])
			copy(dAtA[i:], m.IgnoredQueryParameters[iNdEx])
			i = encodeVarintTranscoder(dAtA, i, uint64(len(m.IgnoredQueryParameters[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.MatchIncomingRequestRoute {
		i--
		if m.MatchIncomingRequestRoute {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.DescriptorSet != nil {
		{
			size := m.DescriptorSet.Size()
			i -= size
			if _, err := m.DescriptorSet.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.PrintOptions != nil {
		{
			size, err := m.PrintOptions.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTranscoder(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Services) > 0 {
		for iNdEx := len(m.Services) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Services[iNdEx])
			copy(dAtA[i:], m.Services[iNdEx])
			i = encodeVarintTranscoder(dAtA, i, uint64(len(m.Services[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func (m *GrpcJsonTranscoder_ProtoDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GrpcJsonTranscoder_ProtoDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.ProtoDescriptor)
	copy(dAtA[i:], m.ProtoDescriptor)
	i = encodeVarintTranscoder(dAtA, i, uint64(len(m.ProtoDescriptor)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *GrpcJsonTranscoder_ProtoDescriptorBin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GrpcJsonTranscoder_ProtoDescriptorBin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ProtoDescriptorBin != nil {
		i -= len(m.ProtoDescriptorBin)
		copy(dAtA[i:], m.ProtoDescriptorBin)
		i = encodeVarintTranscoder(dAtA, i, uint64(len(m.ProtoDescriptorBin)))
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *GrpcJsonTranscoder_PrintOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrpcJsonTranscoder_PrintOptions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GrpcJsonTranscoder_PrintOptions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.PreserveProtoFieldNames {
		i--
		if m.PreserveProtoFieldNames {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.AlwaysPrintEnumsAsInts {
		i--
		if m.AlwaysPrintEnumsAsInts {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.AlwaysPrintPrimitiveFields {
		i--
		if m.AlwaysPrintPrimitiveFields {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.AddWhitespace {
		i--
		if m.AddWhitespace {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTranscoder(dAtA []byte, offset int, v uint64) int {
	offset -= sovTranscoder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GrpcJsonTranscoder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DescriptorSet != nil {
		n += m.DescriptorSet.Size()
	}
	if len(m.Services) > 0 {
		for _, s := range m.Services {
			l = len(s)
			n += 1 + l + sovTranscoder(uint64(l))
		}
	}
	if m.PrintOptions != nil {
		l = m.PrintOptions.Size()
		n += 1 + l + sovTranscoder(uint64(l))
	}
	if m.MatchIncomingRequestRoute {
		n += 2
	}
	if len(m.IgnoredQueryParameters) > 0 {
		for _, s := range m.IgnoredQueryParameters {
			l = len(s)
			n += 1 + l + sovTranscoder(uint64(l))
		}
	}
	if m.AutoMapping {
		n += 2
	}
	if m.IgnoreUnknownQueryParameters {
		n += 2
	}
	if m.ConvertGrpcStatus {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GrpcJsonTranscoder_ProtoDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ProtoDescriptor)
	n += 1 + l + sovTranscoder(uint64(l))
	return n
}
func (m *GrpcJsonTranscoder_ProtoDescriptorBin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProtoDescriptorBin != nil {
		l = len(m.ProtoDescriptorBin)
		n += 1 + l + sovTranscoder(uint64(l))
	}
	return n
}
func (m *GrpcJsonTranscoder_PrintOptions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AddWhitespace {
		n += 2
	}
	if m.AlwaysPrintPrimitiveFields {
		n += 2
	}
	if m.AlwaysPrintEnumsAsInts {
		n += 2
	}
	if m.PreserveProtoFieldNames {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTranscoder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTranscoder(x uint64) (n int) {
	return sovTranscoder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GrpcJsonTranscoder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTranscoder
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GrpcJsonTranscoder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrpcJsonTranscoder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtoDescriptor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTranscoder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTranscoder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DescriptorSet = &GrpcJsonTranscoder_ProtoDescriptor{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Services", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTranscoder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTranscoder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Services = append(m.Services, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrintOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTranscoder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTranscoder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PrintOptions == nil {
				m.PrintOptions = &GrpcJsonTranscoder_PrintOptions{}
			}
			if err := m.PrintOptions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtoDescriptorBin", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTranscoder
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTranscoder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.DescriptorSet = &GrpcJsonTranscoder_ProtoDescriptorBin{v}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MatchIncomingRequestRoute", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MatchIncomingRequestRoute = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IgnoredQueryParameters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTranscoder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTranscoder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IgnoredQueryParameters = append(m.IgnoredQueryParameters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AutoMapping", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AutoMapping = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IgnoreUnknownQueryParameters", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IgnoreUnknownQueryParameters = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConvertGrpcStatus", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ConvertGrpcStatus = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTranscoder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTranscoder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTranscoder
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GrpcJsonTranscoder_PrintOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTranscoder
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PrintOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrintOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddWhitespace", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AddWhitespace = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlwaysPrintPrimitiveFields", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AlwaysPrintPrimitiveFields = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlwaysPrintEnumsAsInts", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AlwaysPrintEnumsAsInts = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreserveProtoFieldNames", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PreserveProtoFieldNames = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTranscoder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTranscoder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTranscoder
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTranscoder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTranscoder
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
					return 0, ErrIntOverflowTranscoder
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTranscoder
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
			if length < 0 {
				return 0, ErrInvalidLengthTranscoder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTranscoder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTranscoder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTranscoder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTranscoder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTranscoder = fmt.Errorf("proto: unexpected end of group")
)
