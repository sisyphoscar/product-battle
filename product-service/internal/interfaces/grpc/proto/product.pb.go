// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: product.proto

package product_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Product struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl      string                 `protobuf:"bytes,4,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type ProductList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []*Product             `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductList) Reset() {
	*x = ProductList{}
	mi := &file_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductList) ProtoMessage() {}

func (x *ProductList) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductList.ProtoReflect.Descriptor instead.
func (*ProductList) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductList) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_product_proto protoreflect.FileDescriptor

const file_product_proto_rawDesc = "" +
	"\n" +
	"\rproduct.proto\x12\rproduct_proto\x1a\x1bgoogle/protobuf/empty.proto\"k\n" +
	"\aProduct\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x04R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x1a\n" +
	"\bimageUrl\x18\x04 \x01(\tR\bimageUrl\"A\n" +
	"\vProductList\x122\n" +
	"\bproducts\x18\x01 \x03(\v2\x16.product_proto.ProductR\bproducts2V\n" +
	"\x0eProductService\x12D\n" +
	"\x0eGetAllProducts\x12\x16.google.protobuf.Empty\x1a\x1a.product_proto.ProductListBMZK/product-service/internal/interfaces/grpc/proto/product_proto;product_protob\x06proto3"

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData []byte
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_product_proto_rawDesc), len(file_product_proto_rawDesc)))
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_product_proto_goTypes = []any{
	(*Product)(nil),       // 0: product_proto.Product
	(*ProductList)(nil),   // 1: product_proto.ProductList
	(*emptypb.Empty)(nil), // 2: google.protobuf.Empty
}
var file_product_proto_depIdxs = []int32{
	0, // 0: product_proto.ProductList.products:type_name -> product_proto.Product
	2, // 1: product_proto.ProductService.GetAllProducts:input_type -> google.protobuf.Empty
	1, // 2: product_proto.ProductService.GetAllProducts:output_type -> product_proto.ProductList
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_product_proto_rawDesc), len(file_product_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}
