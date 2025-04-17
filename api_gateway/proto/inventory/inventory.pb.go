// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: api_gateway/proto/inventory/inventory.proto

package inventory

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// Product message represents a product in the inventory
type Product struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CategoryId    string                 `protobuf:"bytes,4,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Price         float64                `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	Stock         int32                  `protobuf:"varint,6,opt,name=stock,proto3" json:"stock,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[0]
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
	return file_api_gateway_proto_inventory_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
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

func (x *Product) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *Product) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *Product) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Product) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// ProductFilter for filtering products
type ProductFilter struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CategoryId    string                 `protobuf:"bytes,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	MinPrice      float64                `protobuf:"fixed64,3,opt,name=min_price,json=minPrice,proto3" json:"min_price,omitempty"`
	MaxPrice      float64                `protobuf:"fixed64,4,opt,name=max_price,json=maxPrice,proto3" json:"max_price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductFilter) Reset() {
	*x = ProductFilter{}
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductFilter) ProtoMessage() {}

func (x *ProductFilter) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductFilter.ProtoReflect.Descriptor instead.
func (*ProductFilter) Descriptor() ([]byte, []int) {
	return file_api_gateway_proto_inventory_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *ProductFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductFilter) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *ProductFilter) GetMinPrice() float64 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *ProductFilter) GetMaxPrice() float64 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

// CreateProductRequest for creating a new product
type CreateProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CategoryId    string                 `protobuf:"bytes,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Stock         int32                  `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_api_gateway_proto_inventory_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *CreateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *CreateProductRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateProductRequest) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

// UpdateProductRequest for updating an existing product
type UpdateProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CategoryId    string                 `protobuf:"bytes,4,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Price         float64                `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	Stock         int32                  `protobuf:"varint,6,opt,name=stock,proto3" json:"stock,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProductRequest) Reset() {
	*x = UpdateProductRequest{}
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductRequest) ProtoMessage() {}

func (x *UpdateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductRequest) Descriptor() ([]byte, []int) {
	return file_api_gateway_proto_inventory_inventory_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateProductRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *UpdateProductRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateProductRequest) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

// GetProductRequest for getting a product by id
type GetProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_api_gateway_proto_inventory_inventory_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DeleteProductRequest for deleting a product by id
type DeleteProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProductRequest) Reset() {
	*x = DeleteProductRequest{}
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductRequest) ProtoMessage() {}

func (x *DeleteProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductRequest) Descriptor() ([]byte, []int) {
	return file_api_gateway_proto_inventory_inventory_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// ListProductsRequest for listing products with filtering and pagination
type ListProductsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Filter        *ProductFilter         `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	Page          int32                  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProductsRequest) Reset() {
	*x = ListProductsRequest{}
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsRequest) ProtoMessage() {}

func (x *ListProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsRequest.ProtoReflect.Descriptor instead.
func (*ListProductsRequest) Descriptor() ([]byte, []int) {
	return file_api_gateway_proto_inventory_inventory_proto_rawDescGZIP(), []int{6}
}

func (x *ListProductsRequest) GetFilter() *ProductFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ListProductsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListProductsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

// ListProductsResponse for returning a list of products
type ListProductsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []*Product             `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProductsResponse) Reset() {
	*x = ListProductsResponse{}
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsResponse) ProtoMessage() {}

func (x *ListProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsResponse.ProtoReflect.Descriptor instead.
func (*ListProductsResponse) Descriptor() ([]byte, []int) {
	return file_api_gateway_proto_inventory_inventory_proto_rawDescGZIP(), []int{7}
}

func (x *ListProductsResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

// Category message represents a product category
type Category struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Category) Reset() {
	*x = Category{}
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_api_gateway_proto_inventory_inventory_proto_rawDescGZIP(), []int{8}
}

func (x *Category) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// CreateCategoryRequest for creating a new category
type CreateCategoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCategoryRequest) Reset() {
	*x = CreateCategoryRequest{}
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryRequest) ProtoMessage() {}

func (x *CreateCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryRequest.ProtoReflect.Descriptor instead.
func (*CreateCategoryRequest) Descriptor() ([]byte, []int) {
	return file_api_gateway_proto_inventory_inventory_proto_rawDescGZIP(), []int{9}
}

func (x *CreateCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCategoryRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// UpdateCategoryRequest for updating an existing category
type UpdateCategoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCategoryRequest) Reset() {
	*x = UpdateCategoryRequest{}
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCategoryRequest) ProtoMessage() {}

func (x *UpdateCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCategoryRequest.ProtoReflect.Descriptor instead.
func (*UpdateCategoryRequest) Descriptor() ([]byte, []int) {
	return file_api_gateway_proto_inventory_inventory_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateCategoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCategoryRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// GetCategoryRequest for getting a category by id
type GetCategoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCategoryRequest) Reset() {
	*x = GetCategoryRequest{}
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryRequest) ProtoMessage() {}

func (x *GetCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryRequest.ProtoReflect.Descriptor instead.
func (*GetCategoryRequest) Descriptor() ([]byte, []int) {
	return file_api_gateway_proto_inventory_inventory_proto_rawDescGZIP(), []int{11}
}

func (x *GetCategoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// DeleteCategoryRequest for deleting a category by id
type DeleteCategoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCategoryRequest) Reset() {
	*x = DeleteCategoryRequest{}
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCategoryRequest) ProtoMessage() {}

func (x *DeleteCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCategoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteCategoryRequest) Descriptor() ([]byte, []int) {
	return file_api_gateway_proto_inventory_inventory_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteCategoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// ListCategoriesResponse for returning a list of categories
type ListCategoriesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Categories    []*Category            `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCategoriesResponse) Reset() {
	*x = ListCategoriesResponse{}
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoriesResponse) ProtoMessage() {}

func (x *ListCategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_proto_inventory_inventory_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoriesResponse.ProtoReflect.Descriptor instead.
func (*ListCategoriesResponse) Descriptor() ([]byte, []int) {
	return file_api_gateway_proto_inventory_inventory_proto_rawDescGZIP(), []int{13}
}

func (x *ListCategoriesResponse) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

var File_api_gateway_proto_inventory_inventory_proto protoreflect.FileDescriptor

const file_api_gateway_proto_inventory_inventory_proto_rawDesc = "" +
	"\n" +
	"+api_gateway/proto/inventory/inventory.proto\x12\tinventory\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1bgoogle/protobuf/empty.proto\"\x92\x02\n" +
	"\aProduct\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x1f\n" +
	"\vcategory_id\x18\x04 \x01(\tR\n" +
	"categoryId\x12\x14\n" +
	"\x05price\x18\x05 \x01(\x01R\x05price\x12\x14\n" +
	"\x05stock\x18\x06 \x01(\x05R\x05stock\x129\n" +
	"\n" +
	"created_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"~\n" +
	"\rProductFilter\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1f\n" +
	"\vcategory_id\x18\x02 \x01(\tR\n" +
	"categoryId\x12\x1b\n" +
	"\tmin_price\x18\x03 \x01(\x01R\bminPrice\x12\x1b\n" +
	"\tmax_price\x18\x04 \x01(\x01R\bmaxPrice\"\x99\x01\n" +
	"\x14CreateProductRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x1f\n" +
	"\vcategory_id\x18\x03 \x01(\tR\n" +
	"categoryId\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x01R\x05price\x12\x14\n" +
	"\x05stock\x18\x05 \x01(\x05R\x05stock\"\xa9\x01\n" +
	"\x14UpdateProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x1f\n" +
	"\vcategory_id\x18\x04 \x01(\tR\n" +
	"categoryId\x12\x14\n" +
	"\x05price\x18\x05 \x01(\x01R\x05price\x12\x14\n" +
	"\x05stock\x18\x06 \x01(\x05R\x05stock\"#\n" +
	"\x11GetProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"&\n" +
	"\x14DeleteProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"q\n" +
	"\x13ListProductsRequest\x120\n" +
	"\x06filter\x18\x01 \x01(\v2\x18.inventory.ProductFilterR\x06filter\x12\x12\n" +
	"\x04page\x18\x02 \x01(\x05R\x04page\x12\x14\n" +
	"\x05limit\x18\x03 \x01(\x05R\x05limit\"F\n" +
	"\x14ListProductsResponse\x12.\n" +
	"\bproducts\x18\x01 \x03(\v2\x12.inventory.ProductR\bproducts\"P\n" +
	"\bCategory\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\"M\n" +
	"\x15CreateCategoryRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\"]\n" +
	"\x15UpdateCategoryRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\"$\n" +
	"\x12GetCategoryRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"'\n" +
	"\x15DeleteCategoryRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"M\n" +
	"\x16ListCategoriesResponse\x123\n" +
	"\n" +
	"categories\x18\x01 \x03(\v2\x13.inventory.CategoryR\n" +
	"categories2\xf7\x02\n" +
	"\x0eProductService\x12D\n" +
	"\rCreateProduct\x12\x1f.inventory.CreateProductRequest\x1a\x12.inventory.Product\x12>\n" +
	"\n" +
	"GetProduct\x12\x1c.inventory.GetProductRequest\x1a\x12.inventory.Product\x12D\n" +
	"\rUpdateProduct\x12\x1f.inventory.UpdateProductRequest\x1a\x12.inventory.Product\x12H\n" +
	"\rDeleteProduct\x12\x1f.inventory.DeleteProductRequest\x1a\x16.google.protobuf.Empty\x12O\n" +
	"\fListProducts\x12\x1e.inventory.ListProductsRequest\x1a\x1f.inventory.ListProductsResponse2\xff\x02\n" +
	"\x0fCategoryService\x12G\n" +
	"\x0eCreateCategory\x12 .inventory.CreateCategoryRequest\x1a\x13.inventory.Category\x12A\n" +
	"\vGetCategory\x12\x1d.inventory.GetCategoryRequest\x1a\x13.inventory.Category\x12G\n" +
	"\x0eUpdateCategory\x12 .inventory.UpdateCategoryRequest\x1a\x13.inventory.Category\x12J\n" +
	"\x0eDeleteCategory\x12 .inventory.DeleteCategoryRequest\x1a\x16.google.protobuf.Empty\x12K\n" +
	"\x0eListCategories\x12\x16.google.protobuf.Empty\x1a!.inventory.ListCategoriesResponseB\x1dZ\x1bapi_gateway/proto/inventoryb\x06proto3"

var (
	file_api_gateway_proto_inventory_inventory_proto_rawDescOnce sync.Once
	file_api_gateway_proto_inventory_inventory_proto_rawDescData []byte
)

func file_api_gateway_proto_inventory_inventory_proto_rawDescGZIP() []byte {
	file_api_gateway_proto_inventory_inventory_proto_rawDescOnce.Do(func() {
		file_api_gateway_proto_inventory_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_gateway_proto_inventory_inventory_proto_rawDesc), len(file_api_gateway_proto_inventory_inventory_proto_rawDesc)))
	})
	return file_api_gateway_proto_inventory_inventory_proto_rawDescData
}

var file_api_gateway_proto_inventory_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_api_gateway_proto_inventory_inventory_proto_goTypes = []any{
	(*Product)(nil),                // 0: inventory.Product
	(*ProductFilter)(nil),          // 1: inventory.ProductFilter
	(*CreateProductRequest)(nil),   // 2: inventory.CreateProductRequest
	(*UpdateProductRequest)(nil),   // 3: inventory.UpdateProductRequest
	(*GetProductRequest)(nil),      // 4: inventory.GetProductRequest
	(*DeleteProductRequest)(nil),   // 5: inventory.DeleteProductRequest
	(*ListProductsRequest)(nil),    // 6: inventory.ListProductsRequest
	(*ListProductsResponse)(nil),   // 7: inventory.ListProductsResponse
	(*Category)(nil),               // 8: inventory.Category
	(*CreateCategoryRequest)(nil),  // 9: inventory.CreateCategoryRequest
	(*UpdateCategoryRequest)(nil),  // 10: inventory.UpdateCategoryRequest
	(*GetCategoryRequest)(nil),     // 11: inventory.GetCategoryRequest
	(*DeleteCategoryRequest)(nil),  // 12: inventory.DeleteCategoryRequest
	(*ListCategoriesResponse)(nil), // 13: inventory.ListCategoriesResponse
	(*timestamppb.Timestamp)(nil),  // 14: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),          // 15: google.protobuf.Empty
}
var file_api_gateway_proto_inventory_inventory_proto_depIdxs = []int32{
	14, // 0: inventory.Product.created_at:type_name -> google.protobuf.Timestamp
	14, // 1: inventory.Product.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 2: inventory.ListProductsRequest.filter:type_name -> inventory.ProductFilter
	0,  // 3: inventory.ListProductsResponse.products:type_name -> inventory.Product
	8,  // 4: inventory.ListCategoriesResponse.categories:type_name -> inventory.Category
	2,  // 5: inventory.ProductService.CreateProduct:input_type -> inventory.CreateProductRequest
	4,  // 6: inventory.ProductService.GetProduct:input_type -> inventory.GetProductRequest
	3,  // 7: inventory.ProductService.UpdateProduct:input_type -> inventory.UpdateProductRequest
	5,  // 8: inventory.ProductService.DeleteProduct:input_type -> inventory.DeleteProductRequest
	6,  // 9: inventory.ProductService.ListProducts:input_type -> inventory.ListProductsRequest
	9,  // 10: inventory.CategoryService.CreateCategory:input_type -> inventory.CreateCategoryRequest
	11, // 11: inventory.CategoryService.GetCategory:input_type -> inventory.GetCategoryRequest
	10, // 12: inventory.CategoryService.UpdateCategory:input_type -> inventory.UpdateCategoryRequest
	12, // 13: inventory.CategoryService.DeleteCategory:input_type -> inventory.DeleteCategoryRequest
	15, // 14: inventory.CategoryService.ListCategories:input_type -> google.protobuf.Empty
	0,  // 15: inventory.ProductService.CreateProduct:output_type -> inventory.Product
	0,  // 16: inventory.ProductService.GetProduct:output_type -> inventory.Product
	0,  // 17: inventory.ProductService.UpdateProduct:output_type -> inventory.Product
	15, // 18: inventory.ProductService.DeleteProduct:output_type -> google.protobuf.Empty
	7,  // 19: inventory.ProductService.ListProducts:output_type -> inventory.ListProductsResponse
	8,  // 20: inventory.CategoryService.CreateCategory:output_type -> inventory.Category
	8,  // 21: inventory.CategoryService.GetCategory:output_type -> inventory.Category
	8,  // 22: inventory.CategoryService.UpdateCategory:output_type -> inventory.Category
	15, // 23: inventory.CategoryService.DeleteCategory:output_type -> google.protobuf.Empty
	13, // 24: inventory.CategoryService.ListCategories:output_type -> inventory.ListCategoriesResponse
	15, // [15:25] is the sub-list for method output_type
	5,  // [5:15] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_gateway_proto_inventory_inventory_proto_init() }
func file_api_gateway_proto_inventory_inventory_proto_init() {
	if File_api_gateway_proto_inventory_inventory_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_gateway_proto_inventory_inventory_proto_rawDesc), len(file_api_gateway_proto_inventory_inventory_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_gateway_proto_inventory_inventory_proto_goTypes,
		DependencyIndexes: file_api_gateway_proto_inventory_inventory_proto_depIdxs,
		MessageInfos:      file_api_gateway_proto_inventory_inventory_proto_msgTypes,
	}.Build()
	File_api_gateway_proto_inventory_inventory_proto = out.File
	file_api_gateway_proto_inventory_inventory_proto_goTypes = nil
	file_api_gateway_proto_inventory_inventory_proto_depIdxs = nil
}
