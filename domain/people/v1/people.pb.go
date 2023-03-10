// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: domain/people/v1/people.proto

package peoplepb

import (
	v1 "github.com/kubediscovery/apis/entity/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Domain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity   *v1.Entity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	DomainID string     `protobuf:"bytes,2,opt,name=domainID,proto3" json:"domainID,omitempty"`
	TeamID   string     `protobuf:"bytes,3,opt,name=teamID,proto3" json:"teamID,omitempty"`
	Owner    bool       `protobuf:"varint,4,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *Domain) Reset() {
	*x = Domain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_people_v1_people_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Domain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Domain) ProtoMessage() {}

func (x *Domain) ProtoReflect() protoreflect.Message {
	mi := &file_domain_people_v1_people_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Domain.ProtoReflect.Descriptor instead.
func (*Domain) Descriptor() ([]byte, []int) {
	return file_domain_people_v1_people_proto_rawDescGZIP(), []int{0}
}

func (x *Domain) GetEntity() *v1.Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *Domain) GetDomainID() string {
	if x != nil {
		return x.DomainID
	}
	return ""
}

func (x *Domain) GetTeamID() string {
	if x != nil {
		return x.TeamID
	}
	return ""
}

func (x *Domain) GetOwner() bool {
	if x != nil {
		return x.Owner
	}
	return false
}

var File_domain_people_v1_people_proto protoreflect.FileDescriptor

var file_domain_people_v1_people_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a,
	0x16, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x29, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x61, 0x6d,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x3b, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_people_v1_people_proto_rawDescOnce sync.Once
	file_domain_people_v1_people_proto_rawDescData = file_domain_people_v1_people_proto_rawDesc
)

func file_domain_people_v1_people_proto_rawDescGZIP() []byte {
	file_domain_people_v1_people_proto_rawDescOnce.Do(func() {
		file_domain_people_v1_people_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_people_v1_people_proto_rawDescData)
	})
	return file_domain_people_v1_people_proto_rawDescData
}

var file_domain_people_v1_people_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_domain_people_v1_people_proto_goTypes = []interface{}{
	(*Domain)(nil),    // 0: people.v1beta1.Domain
	(*v1.Entity)(nil), // 1: entity.v1.Entity
}
var file_domain_people_v1_people_proto_depIdxs = []int32{
	1, // 0: people.v1beta1.Domain.entity:type_name -> entity.v1.Entity
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_domain_people_v1_people_proto_init() }
func file_domain_people_v1_people_proto_init() {
	if File_domain_people_v1_people_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_people_v1_people_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Domain); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_domain_people_v1_people_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_people_v1_people_proto_goTypes,
		DependencyIndexes: file_domain_people_v1_people_proto_depIdxs,
		MessageInfos:      file_domain_people_v1_people_proto_msgTypes,
	}.Build()
	File_domain_people_v1_people_proto = out.File
	file_domain_people_v1_people_proto_rawDesc = nil
	file_domain_people_v1_people_proto_goTypes = nil
	file_domain_people_v1_people_proto_depIdxs = nil
}
