// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/assetsDefinition.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AssetCategory int32

const (
	AssetCategory_Vehicular     AssetCategory = 0
	AssetCategory_Proffessional AssetCategory = 1
	AssetCategory_Housing       AssetCategory = 2
	AssetCategory_Computing     AssetCategory = 3
)

var AssetCategory_name = map[int32]string{
	0: "Vehicular",
	1: "Proffessional",
	2: "Housing",
	3: "Computing",
}

var AssetCategory_value = map[string]int32{
	"Vehicular":     0,
	"Proffessional": 1,
	"Housing":       2,
	"Computing":     3,
}

func (x AssetCategory) String() string {
	return proto.EnumName(AssetCategory_name, int32(x))
}

func (AssetCategory) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2cb015c5e1f6559e, []int{0}
}

type PersonCategory int32

const (
	PersonCategory_Lessee PersonCategory = 0
	PersonCategory_Lessor PersonCategory = 1
)

var PersonCategory_name = map[int32]string{
	0: "Lessee",
	1: "Lessor",
}

var PersonCategory_value = map[string]int32{
	"Lessee": 0,
	"Lessor": 1,
}

func (x PersonCategory) String() string {
	return proto.EnumName(PersonCategory_name, int32(x))
}

func (PersonCategory) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2cb015c5e1f6559e, []int{1}
}

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2cb015c5e1f6559e, []int{2, 0}
}

type Asset struct {
	Name                      string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                        int32         `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Owner                     *Person       `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Functionality_Description string        `protobuf:"bytes,4,opt,name=Functionality_Description,json=FunctionalityDescription,proto3" json:"Functionality_Description,omitempty"`
	Availability              bool          `protobuf:"varint,5,opt,name=Availability,proto3" json:"Availability,omitempty"`
	Timeslots                 *Timeslots    `protobuf:"bytes,6,opt,name=timeslots,proto3" json:"timeslots,omitempty"`
	Category                  AssetCategory `protobuf:"varint,7,opt,name=category,proto3,enum=AssetCategory" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}      `json:"-"`
	XXX_unrecognized          []byte        `json:"-"`
	XXX_sizecache             int32         `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cb015c5e1f6559e, []int{0}
}

func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Asset) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Asset) GetOwner() *Person {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Asset) GetFunctionality_Description() string {
	if m != nil {
		return m.Functionality_Description
	}
	return ""
}

func (m *Asset) GetAvailability() bool {
	if m != nil {
		return m.Availability
	}
	return false
}

func (m *Asset) GetTimeslots() *Timeslots {
	if m != nil {
		return m.Timeslots
	}
	return nil
}

func (m *Asset) GetCategory() AssetCategory {
	if m != nil {
		return m.Category
	}
	return AssetCategory_Vehicular
}

type Timeslots struct {
	Month                string   `protobuf:"bytes,1,opt,name=Month,proto3" json:"Month,omitempty"`
	Day                  string   `protobuf:"bytes,2,opt,name=Day,proto3" json:"Day,omitempty"`
	Hours                float32  `protobuf:"fixed32,3,opt,name=Hours,proto3" json:"Hours,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timeslots) Reset()         { *m = Timeslots{} }
func (m *Timeslots) String() string { return proto.CompactTextString(m) }
func (*Timeslots) ProtoMessage()    {}
func (*Timeslots) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cb015c5e1f6559e, []int{1}
}

func (m *Timeslots) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timeslots.Unmarshal(m, b)
}
func (m *Timeslots) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timeslots.Marshal(b, m, deterministic)
}
func (m *Timeslots) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timeslots.Merge(m, src)
}
func (m *Timeslots) XXX_Size() int {
	return xxx_messageInfo_Timeslots.Size(m)
}
func (m *Timeslots) XXX_DiscardUnknown() {
	xxx_messageInfo_Timeslots.DiscardUnknown(m)
}

var xxx_messageInfo_Timeslots proto.InternalMessageInfo

func (m *Timeslots) GetMonth() string {
	if m != nil {
		return m.Month
	}
	return ""
}

func (m *Timeslots) GetDay() string {
	if m != nil {
		return m.Day
	}
	return ""
}

func (m *Timeslots) GetHours() float32 {
	if m != nil {
		return m.Hours
	}
	return 0
}

type Person struct {
	FirstName            string                `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string                `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Id                   int32                 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Email                string                `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,5,rep,name=phones,proto3" json:"phones,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cb015c5e1f6559e, []int{2}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Person) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cb015c5e1f6559e, []int{2, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// remove property person to allow only for asset name query(Asset category search)
type GetAssetAvailabilityRequest struct {
	Assetcategory        AssetCategory `protobuf:"varint,2,opt,name=assetcategory,proto3,enum=AssetCategory" json:"assetcategory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetAssetAvailabilityRequest) Reset()         { *m = GetAssetAvailabilityRequest{} }
func (m *GetAssetAvailabilityRequest) String() string { return proto.CompactTextString(m) }
func (*GetAssetAvailabilityRequest) ProtoMessage()    {}
func (*GetAssetAvailabilityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cb015c5e1f6559e, []int{3}
}

func (m *GetAssetAvailabilityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAssetAvailabilityRequest.Unmarshal(m, b)
}
func (m *GetAssetAvailabilityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAssetAvailabilityRequest.Marshal(b, m, deterministic)
}
func (m *GetAssetAvailabilityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAssetAvailabilityRequest.Merge(m, src)
}
func (m *GetAssetAvailabilityRequest) XXX_Size() int {
	return xxx_messageInfo_GetAssetAvailabilityRequest.Size(m)
}
func (m *GetAssetAvailabilityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAssetAvailabilityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAssetAvailabilityRequest proto.InternalMessageInfo

func (m *GetAssetAvailabilityRequest) GetAssetcategory() AssetCategory {
	if m != nil {
		return m.Assetcategory
	}
	return AssetCategory_Vehicular
}

// return only repeted assets that belong to queried category
type AssetsofReqCategory struct {
	Avalability          *Asset   `protobuf:"bytes,1,opt,name=Avalability,proto3" json:"Avalability,omitempty"`
	Timeslots            *Asset   `protobuf:"bytes,2,opt,name=timeslots,proto3" json:"timeslots,omitempty"`
	Name                 *Asset   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Owner                *Asset   `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetsofReqCategory) Reset()         { *m = AssetsofReqCategory{} }
func (m *AssetsofReqCategory) String() string { return proto.CompactTextString(m) }
func (*AssetsofReqCategory) ProtoMessage()    {}
func (*AssetsofReqCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cb015c5e1f6559e, []int{4}
}

func (m *AssetsofReqCategory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetsofReqCategory.Unmarshal(m, b)
}
func (m *AssetsofReqCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetsofReqCategory.Marshal(b, m, deterministic)
}
func (m *AssetsofReqCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetsofReqCategory.Merge(m, src)
}
func (m *AssetsofReqCategory) XXX_Size() int {
	return xxx_messageInfo_AssetsofReqCategory.Size(m)
}
func (m *AssetsofReqCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetsofReqCategory.DiscardUnknown(m)
}

var xxx_messageInfo_AssetsofReqCategory proto.InternalMessageInfo

func (m *AssetsofReqCategory) GetAvalability() *Asset {
	if m != nil {
		return m.Avalability
	}
	return nil
}

func (m *AssetsofReqCategory) GetTimeslots() *Asset {
	if m != nil {
		return m.Timeslots
	}
	return nil
}

func (m *AssetsofReqCategory) GetName() *Asset {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AssetsofReqCategory) GetOwner() *Asset {
	if m != nil {
		return m.Owner
	}
	return nil
}

type AssetAvailabilityResponse struct {
	Assetsrequested      []*AssetsofReqCategory `protobuf:"bytes,1,rep,name=assetsrequested,proto3" json:"assetsrequested,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AssetAvailabilityResponse) Reset()         { *m = AssetAvailabilityResponse{} }
func (m *AssetAvailabilityResponse) String() string { return proto.CompactTextString(m) }
func (*AssetAvailabilityResponse) ProtoMessage()    {}
func (*AssetAvailabilityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cb015c5e1f6559e, []int{5}
}

func (m *AssetAvailabilityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetAvailabilityResponse.Unmarshal(m, b)
}
func (m *AssetAvailabilityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetAvailabilityResponse.Marshal(b, m, deterministic)
}
func (m *AssetAvailabilityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetAvailabilityResponse.Merge(m, src)
}
func (m *AssetAvailabilityResponse) XXX_Size() int {
	return xxx_messageInfo_AssetAvailabilityResponse.Size(m)
}
func (m *AssetAvailabilityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetAvailabilityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AssetAvailabilityResponse proto.InternalMessageInfo

func (m *AssetAvailabilityResponse) GetAssetsrequested() []*AssetsofReqCategory {
	if m != nil {
		return m.Assetsrequested
	}
	return nil
}

func init() {
	proto.RegisterEnum("AssetCategory", AssetCategory_name, AssetCategory_value)
	proto.RegisterEnum("PersonCategory", PersonCategory_name, PersonCategory_value)
	proto.RegisterEnum("Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
	proto.RegisterType((*Asset)(nil), "Asset")
	proto.RegisterType((*Timeslots)(nil), "Timeslots")
	proto.RegisterType((*Person)(nil), "Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "Person.PhoneNumber")
	proto.RegisterType((*GetAssetAvailabilityRequest)(nil), "getAssetAvailabilityRequest")
	proto.RegisterType((*AssetsofReqCategory)(nil), "AssetsofReqCategory")
	proto.RegisterType((*AssetAvailabilityResponse)(nil), "AssetAvailabilityResponse")
}

func init() { proto.RegisterFile("protos/assetsDefinition.proto", fileDescriptor_2cb015c5e1f6559e) }

var fileDescriptor_2cb015c5e1f6559e = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x4d, 0x4f, 0xdb, 0x40,
	0x10, 0xc5, 0x4e, 0x6c, 0xe2, 0x49, 0x93, 0x9a, 0x01, 0x55, 0x26, 0x80, 0x14, 0x59, 0xad, 0x64,
	0x81, 0xe4, 0x4a, 0x69, 0x6f, 0x48, 0x95, 0xf8, 0x14, 0xa8, 0x04, 0xd0, 0x96, 0xb6, 0x52, 0x7b,
	0x40, 0x26, 0x6c, 0xc8, 0xaa, 0x89, 0xd7, 0x78, 0xd7, 0x54, 0x39, 0xf6, 0x2f, 0xf4, 0x0f, 0xf4,
	0xaf, 0x56, 0xde, 0x75, 0x8c, 0xdd, 0xd2, 0x9e, 0x32, 0xf3, 0xde, 0xd3, 0x64, 0xe7, 0xcd, 0x8c,
	0x61, 0x2b, 0x49, 0xb9, 0xe4, 0xe2, 0x75, 0x24, 0x04, 0x95, 0xe2, 0x90, 0x8e, 0x59, 0xcc, 0x24,
	0xe3, 0x71, 0xa8, 0x70, 0xff, 0x87, 0x09, 0xd6, 0x5e, 0x4e, 0x21, 0x42, 0x33, 0x8e, 0x66, 0xd4,
	0x33, 0xfa, 0x46, 0xe0, 0x10, 0x15, 0x63, 0x17, 0x4c, 0x76, 0xeb, 0x99, 0x7d, 0x23, 0xb0, 0x88,
	0xc9, 0x6e, 0x71, 0x0b, 0x2c, 0xfe, 0x3d, 0xa6, 0xa9, 0xd7, 0xe8, 0x1b, 0x41, 0x7b, 0xb0, 0x1c,
	0x5e, 0xd2, 0x54, 0xf0, 0x98, 0x68, 0x14, 0x77, 0x61, 0xfd, 0x38, 0x8b, 0x47, 0x79, 0xf9, 0x68,
	0xca, 0xe4, 0xfc, 0xfa, 0x90, 0x8a, 0x51, 0xca, 0x92, 0x1c, 0xf0, 0x9a, 0xaa, 0xae, 0x57, 0x13,
	0x54, 0x78, 0xf4, 0xe1, 0xd9, 0xde, 0x43, 0xc4, 0xa6, 0xd1, 0x0d, 0xcb, 0x29, 0xcf, 0xea, 0x1b,
	0x41, 0x8b, 0xd4, 0x30, 0x0c, 0xc0, 0x91, 0x6c, 0x46, 0xc5, 0x94, 0x4b, 0xe1, 0xd9, 0xea, 0x0d,
	0x10, 0x5e, 0x2d, 0x10, 0xf2, 0x48, 0xe2, 0x36, 0xb4, 0x46, 0x91, 0xa4, 0x77, 0x3c, 0x9d, 0x7b,
	0xcb, 0x7d, 0x23, 0xe8, 0x0e, 0xba, 0xa1, 0xea, 0xf3, 0xa0, 0x40, 0x49, 0xc9, 0xfb, 0xa7, 0xe0,
	0x94, 0x35, 0x70, 0x0d, 0xac, 0x21, 0x8f, 0xe5, 0xa4, 0xf0, 0x41, 0x27, 0xe8, 0x42, 0xe3, 0x30,
	0x9a, 0x2b, 0x27, 0x1c, 0x92, 0x87, 0xb9, 0xee, 0x84, 0x67, 0xa9, 0x50, 0x56, 0x98, 0x44, 0x27,
	0xfe, 0x4f, 0x13, 0x6c, 0xed, 0x09, 0x6e, 0x01, 0x8c, 0x59, 0x2a, 0xe4, 0x75, 0xc5, 0x55, 0x47,
	0x21, 0xe7, 0xb9, 0xb5, 0x1b, 0xe0, 0x4c, 0xa3, 0x05, 0xab, 0xeb, 0xb6, 0x72, 0xe0, 0xfc, 0xd1,
	0xf7, 0x46, 0xe9, 0xfb, 0x1a, 0x58, 0x74, 0x16, 0xb1, 0x69, 0x61, 0xa2, 0x4e, 0x70, 0x07, 0xec,
	0x64, 0xc2, 0x63, 0x2a, 0x3c, 0xab, 0xdf, 0x08, 0xda, 0x83, 0xd5, 0x62, 0x1c, 0xe1, 0x65, 0x8e,
	0x9e, 0x67, 0xb3, 0x1b, 0x9a, 0x92, 0x42, 0xd2, 0x3b, 0x83, 0x76, 0x05, 0xc6, 0x17, 0x60, 0xc7,
	0x2a, 0x2a, 0x5e, 0x56, 0x64, 0xf8, 0x0a, 0x9a, 0x72, 0x9e, 0xe8, 0x17, 0x75, 0x07, 0x2b, 0xb5,
	0x8a, 0x57, 0xf3, 0x84, 0x12, 0x45, 0xfb, 0x3b, 0xe0, 0x94, 0x10, 0x02, 0xd8, 0xc3, 0x8b, 0xfd,
	0xd3, 0xb3, 0x23, 0x77, 0x09, 0x5b, 0xd0, 0x3c, 0xb9, 0x18, 0x1e, 0xb9, 0x46, 0x1e, 0x7d, 0xbe,
	0x20, 0xef, 0x5d, 0xd3, 0xff, 0x00, 0x1b, 0x77, 0x54, 0x2a, 0xf7, 0xab, 0xd3, 0x24, 0xf4, 0x3e,
	0xa3, 0x42, 0xe2, 0x5b, 0xe8, 0xa8, 0xe5, 0x2c, 0xe7, 0x65, 0x3e, 0x39, 0xaf, 0xba, 0xc8, 0xff,
	0x65, 0xc0, 0xaa, 0x12, 0x08, 0x3e, 0x26, 0xf4, 0x7e, 0x21, 0xc3, 0x00, 0xda, 0x7b, 0x0f, 0x51,
	0xb9, 0x45, 0x86, 0x5a, 0x12, 0x5b, 0xd7, 0x22, 0x55, 0x0a, 0x5f, 0x56, 0x97, 0xc9, 0xac, 0xe9,
	0x2a, 0x8b, 0xd4, 0x2b, 0xce, 0xa2, 0x51, 0x13, 0xe8, 0xf3, 0xd8, 0x5c, 0x9c, 0x43, 0xb3, 0x46,
	0x6a, 0xd0, 0xff, 0x0a, 0xeb, 0x4f, 0xf4, 0x2c, 0x12, 0x1e, 0x0b, 0x8a, 0xef, 0xe0, 0xb9, 0xbe,
	0xc8, 0x54, 0xbb, 0x40, 0x6f, 0x3d, 0x43, 0x0d, 0x71, 0x2d, 0x7c, 0xa2, 0x2b, 0xf2, 0xa7, 0x78,
	0x7b, 0x08, 0x9d, 0x9a, 0x3d, 0xd8, 0x01, 0xe7, 0x13, 0x9d, 0xb0, 0x51, 0x36, 0x8d, 0x52, 0x77,
	0x09, 0x57, 0xa0, 0x73, 0x99, 0xf2, 0xf1, 0x98, 0x0a, 0xa1, 0x8e, 0xcd, 0x35, 0xb0, 0x0d, 0xcb,
	0x27, 0x3c, 0x13, 0x2c, 0xbe, 0x73, 0xcd, 0x5c, 0x7e, 0xc0, 0x67, 0x49, 0x26, 0xf3, 0xb4, 0xb1,
	0x1d, 0x40, 0x57, 0x4f, 0xba, 0xac, 0x07, 0x60, 0x9f, 0x51, 0x21, 0x28, 0x75, 0x97, 0x16, 0x31,
	0x4f, 0x5d, 0x63, 0x30, 0x02, 0xd0, 0x7f, 0x3c, 0xa1, 0xa3, 0x6f, 0xf8, 0x11, 0xd6, 0x55, 0x70,
	0xcc, 0xd3, 0xbf, 0x7a, 0xc5, 0xcd, 0xf0, 0x3f, 0x63, 0xef, 0xf5, 0xc2, 0x7f, 0xba, 0xb3, 0x0f,
	0x5f, 0x5a, 0xe1, 0xae, 0xfe, 0x70, 0xdd, 0xd8, 0xea, 0xf7, 0xcd, 0xef, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x10, 0xfa, 0xed, 0x56, 0xc9, 0x04, 0x00, 0x00,
}
