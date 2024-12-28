// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: pkg/apiclient/course/v1/catalog.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CourseId    string                 `protobuf:"bytes,2,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	DisplayName string                 `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Instructors []*Instructor          `protobuf:"bytes,5,rep,name=instructors,proto3" json:"instructors,omitempty"`
	PublishedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=published_at,json=publishedAt,proto3" json:"published_at,omitempty"`
	Batches     []*Batch               `protobuf:"bytes,7,rep,name=batches,proto3" json:"batches,omitempty"`
	Price       *Price                 `protobuf:"bytes,8,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apiclient_course_v1_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apiclient_course_v1_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_pkg_apiclient_course_v1_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *Course) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Course) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *Course) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Course) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Course) GetInstructors() []*Instructor {
	if x != nil {
		return x.Instructors
	}
	return nil
}

func (x *Course) GetPublishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PublishedAt
	}
	return nil
}

func (x *Course) GetBatches() []*Batch {
	if x != nil {
		return x.Batches
	}
	return nil
}

func (x *Course) GetPrice() *Price {
	if x != nil {
		return x.Price
	}
	return nil
}

type Batch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BatchId        string                 `protobuf:"bytes,2,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
	DisplayName    string                 `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Course         string                 `protobuf:"bytes,4,opt,name=course,proto3" json:"course,omitempty"`
	StartDate      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate        *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	MaxSeats       int32                  `protobuf:"varint,7,opt,name=max_seats,json=maxSeats,proto3" json:"max_seats,omitempty"`
	AvailableSeats int32                  `protobuf:"varint,8,opt,name=available_seats,json=availableSeats,proto3" json:"available_seats,omitempty"`
	Price          *Price                 `protobuf:"bytes,9,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Batch) Reset() {
	*x = Batch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apiclient_course_v1_catalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Batch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Batch) ProtoMessage() {}

func (x *Batch) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apiclient_course_v1_catalog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Batch.ProtoReflect.Descriptor instead.
func (*Batch) Descriptor() ([]byte, []int) {
	return file_pkg_apiclient_course_v1_catalog_proto_rawDescGZIP(), []int{1}
}

func (x *Batch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Batch) GetBatchId() string {
	if x != nil {
		return x.BatchId
	}
	return ""
}

func (x *Batch) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Batch) GetCourse() string {
	if x != nil {
		return x.Course
	}
	return ""
}

func (x *Batch) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *Batch) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *Batch) GetMaxSeats() int32 {
	if x != nil {
		return x.MaxSeats
	}
	return 0
}

func (x *Batch) GetAvailableSeats() int32 {
	if x != nil {
		return x.AvailableSeats
	}
	return 0
}

func (x *Batch) GetPrice() *Price {
	if x != nil {
		return x.Price
	}
	return nil
}

type Instructor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl string   `protobuf:"bytes,2,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Roles    []string `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *Instructor) Reset() {
	*x = Instructor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apiclient_course_v1_catalog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instructor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instructor) ProtoMessage() {}

func (x *Instructor) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apiclient_course_v1_catalog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instructor.ProtoReflect.Descriptor instead.
func (*Instructor) Descriptor() ([]byte, []int) {
	return file_pkg_apiclient_course_v1_catalog_proto_rawDescGZIP(), []int{2}
}

func (x *Instructor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Instructor) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Instructor) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value    float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	Currency string  `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apiclient_course_v1_catalog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apiclient_course_v1_catalog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_pkg_apiclient_course_v1_catalog_proto_rawDescGZIP(), []int{3}
}

func (x *Price) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Price) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type ListCoursesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize  uint64                 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string                 `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	OrderBy   string                 `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	ListMask  *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=list_mask,json=listMask,proto3" json:"list_mask,omitempty"`
}

func (x *ListCoursesRequest) Reset() {
	*x = ListCoursesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apiclient_course_v1_catalog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCoursesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCoursesRequest) ProtoMessage() {}

func (x *ListCoursesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apiclient_course_v1_catalog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCoursesRequest.ProtoReflect.Descriptor instead.
func (*ListCoursesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apiclient_course_v1_catalog_proto_rawDescGZIP(), []int{4}
}

func (x *ListCoursesRequest) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCoursesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListCoursesRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ListCoursesRequest) GetListMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.ListMask
	}
	return nil
}

type ListCoursesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Courses       []*Course `protobuf:"bytes,1,rep,name=courses,proto3" json:"courses,omitempty"`
	NextPageToken string    `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListCoursesResponse) Reset() {
	*x = ListCoursesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apiclient_course_v1_catalog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCoursesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCoursesResponse) ProtoMessage() {}

func (x *ListCoursesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apiclient_course_v1_catalog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCoursesResponse.ProtoReflect.Descriptor instead.
func (*ListCoursesResponse) Descriptor() ([]byte, []int) {
	return file_pkg_apiclient_course_v1_catalog_proto_rawDescGZIP(), []int{5}
}

func (x *ListCoursesResponse) GetCourses() []*Course {
	if x != nil {
		return x.Courses
	}
	return nil
}

func (x *ListCoursesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The course identifier to retrieve
	Course string `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
}

func (x *GetCourseRequest) Reset() {
	*x = GetCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apiclient_course_v1_catalog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCourseRequest) ProtoMessage() {}

func (x *GetCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apiclient_course_v1_catalog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCourseRequest.ProtoReflect.Descriptor instead.
func (*GetCourseRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apiclient_course_v1_catalog_proto_rawDescGZIP(), []int{6}
}

func (x *GetCourseRequest) GetCourse() string {
	if x != nil {
		return x.Course
	}
	return ""
}

var File_pkg_apiclient_course_v1_catalog_proto protoreflect.FileDescriptor

var file_pkg_apiclient_course_v1_catalog_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x69, 0x6d, 0x72, 0x65, 0x6e, 0x61, 0x67,
	0x69, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x03, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x04, 0xe2, 0x41, 0x01, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x09,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x6f, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x69, 0x6d, 0x72, 0x65,
	0x6e, 0x61, 0x67, 0x69, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x61, 0x70, 0x70, 0x2e,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72,
	0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x3e, 0x0a, 0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x69, 0x6d, 0x72, 0x65, 0x6e, 0x61, 0x67, 0x69, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x65, 0x6d, 0x6f, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x12, 0x3a, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x69, 0x6d, 0x72, 0x65, 0x6e, 0x61, 0x67, 0x69, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x65,
	0x6d, 0x6f, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x3a, 0x49, 0xea, 0x41,
	0x46, 0x0a, 0x21, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x61, 0x70,
	0x70, 0x2e, 0x69, 0x6d, 0x72, 0x65, 0x6e, 0x61, 0x67, 0x69, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x12, 0x10, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x7d, 0x2a, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x32,
	0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0xe8, 0x03, 0x0a, 0x05, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x04, 0xe2, 0x41, 0x01, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2,
	0x41, 0x01, 0x03, 0x52, 0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x3e, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x26, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x64, 0x65, 0x6d,
	0x6f, 0x61, 0x70, 0x70, 0x2e, 0x69, 0x6d, 0x72, 0x65, 0x6e, 0x61, 0x67, 0x69, 0x63, 0x6f, 0x6d,
	0x2f, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x65, 0x61, 0x74, 0x73, 0x12, 0x27,
	0x0a, 0x0f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x74,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x53, 0x65, 0x61, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x6d, 0x72, 0x65, 0x6e, 0x61, 0x67,
	0x69, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x3a, 0x4d, 0xea, 0x41, 0x4a, 0x0a, 0x26, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x61, 0x70, 0x70, 0x2e, 0x69, 0x6d, 0x72, 0x65, 0x6e, 0x61, 0x67,
	0x69, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x20, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x7d, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x2f, 0x7b, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x7d, 0x22, 0x53, 0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x39, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x22, 0xa4, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x12, 0x37, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52,
	0x08, 0x6c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x7e, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3f, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6d, 0x72, 0x65, 0x6e, 0x61, 0x67, 0x69, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x65, 0x6d, 0x6f, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x56, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a,
	0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xe2,
	0x41, 0x01, 0x02, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x64,
	0x65, 0x6d, 0x6f, 0x61, 0x70, 0x70, 0x2e, 0x69, 0x6d, 0x72, 0x65, 0x6e, 0x61, 0x67, 0x69, 0x63,
	0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x32, 0xe0, 0x02, 0x0a, 0x0e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0xa6, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x73, 0x12, 0x31, 0x2e, 0x69, 0x6d, 0x72, 0x65, 0x6e, 0x61, 0x67, 0x69, 0x63,
	0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x69, 0x6d, 0x72, 0x65, 0x6e, 0x61,
	0x67, 0x69, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x92, 0x41, 0x0f,
	0x12, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x72, 0x74, 0x73, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x12, 0xa4, 0x01,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x2f, 0x2e, 0x69, 0x6d,
	0x72, 0x65, 0x6e, 0x61, 0x67, 0x69, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x61, 0x70,
	0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x69,
	0x6d, 0x72, 0x65, 0x6e, 0x61, 0x67, 0x69, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x61,
	0x70, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x22, 0x3f, 0x92, 0x41, 0x0c, 0x12, 0x0a, 0x47, 0x65, 0x74, 0x20, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0xda, 0x41, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x7d, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6d, 0x72, 0x65, 0x6e, 0x61, 0x67, 0x69, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x65, 0x6d, 0x6f, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_apiclient_course_v1_catalog_proto_rawDescOnce sync.Once
	file_pkg_apiclient_course_v1_catalog_proto_rawDescData = file_pkg_apiclient_course_v1_catalog_proto_rawDesc
)

func file_pkg_apiclient_course_v1_catalog_proto_rawDescGZIP() []byte {
	file_pkg_apiclient_course_v1_catalog_proto_rawDescOnce.Do(func() {
		file_pkg_apiclient_course_v1_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_apiclient_course_v1_catalog_proto_rawDescData)
	})
	return file_pkg_apiclient_course_v1_catalog_proto_rawDescData
}

var file_pkg_apiclient_course_v1_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_apiclient_course_v1_catalog_proto_goTypes = []any{
	(*Course)(nil),                // 0: imrenagicom.demoapp.course.v1.Course
	(*Batch)(nil),                 // 1: imrenagicom.demoapp.course.v1.Batch
	(*Instructor)(nil),            // 2: imrenagicom.demoapp.course.v1.Instructor
	(*Price)(nil),                 // 3: imrenagicom.demoapp.course.v1.Price
	(*ListCoursesRequest)(nil),    // 4: imrenagicom.demoapp.course.v1.ListCoursesRequest
	(*ListCoursesResponse)(nil),   // 5: imrenagicom.demoapp.course.v1.ListCoursesResponse
	(*GetCourseRequest)(nil),      // 6: imrenagicom.demoapp.course.v1.GetCourseRequest
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil), // 8: google.protobuf.FieldMask
}
var file_pkg_apiclient_course_v1_catalog_proto_depIdxs = []int32{
	2,  // 0: imrenagicom.demoapp.course.v1.Course.instructors:type_name -> imrenagicom.demoapp.course.v1.Instructor
	7,  // 1: imrenagicom.demoapp.course.v1.Course.published_at:type_name -> google.protobuf.Timestamp
	1,  // 2: imrenagicom.demoapp.course.v1.Course.batches:type_name -> imrenagicom.demoapp.course.v1.Batch
	3,  // 3: imrenagicom.demoapp.course.v1.Course.price:type_name -> imrenagicom.demoapp.course.v1.Price
	7,  // 4: imrenagicom.demoapp.course.v1.Batch.start_date:type_name -> google.protobuf.Timestamp
	7,  // 5: imrenagicom.demoapp.course.v1.Batch.end_date:type_name -> google.protobuf.Timestamp
	3,  // 6: imrenagicom.demoapp.course.v1.Batch.price:type_name -> imrenagicom.demoapp.course.v1.Price
	8,  // 7: imrenagicom.demoapp.course.v1.ListCoursesRequest.list_mask:type_name -> google.protobuf.FieldMask
	0,  // 8: imrenagicom.demoapp.course.v1.ListCoursesResponse.courses:type_name -> imrenagicom.demoapp.course.v1.Course
	4,  // 9: imrenagicom.demoapp.course.v1.CatalogService.ListCourses:input_type -> imrenagicom.demoapp.course.v1.ListCoursesRequest
	6,  // 10: imrenagicom.demoapp.course.v1.CatalogService.GetCourse:input_type -> imrenagicom.demoapp.course.v1.GetCourseRequest
	5,  // 11: imrenagicom.demoapp.course.v1.CatalogService.ListCourses:output_type -> imrenagicom.demoapp.course.v1.ListCoursesResponse
	0,  // 12: imrenagicom.demoapp.course.v1.CatalogService.GetCourse:output_type -> imrenagicom.demoapp.course.v1.Course
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_pkg_apiclient_course_v1_catalog_proto_init() }
func file_pkg_apiclient_course_v1_catalog_proto_init() {
	if File_pkg_apiclient_course_v1_catalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_apiclient_course_v1_catalog_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Course); i {
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
		file_pkg_apiclient_course_v1_catalog_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Batch); i {
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
		file_pkg_apiclient_course_v1_catalog_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Instructor); i {
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
		file_pkg_apiclient_course_v1_catalog_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Price); i {
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
		file_pkg_apiclient_course_v1_catalog_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListCoursesRequest); i {
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
		file_pkg_apiclient_course_v1_catalog_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ListCoursesResponse); i {
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
		file_pkg_apiclient_course_v1_catalog_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetCourseRequest); i {
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
			RawDescriptor: file_pkg_apiclient_course_v1_catalog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_apiclient_course_v1_catalog_proto_goTypes,
		DependencyIndexes: file_pkg_apiclient_course_v1_catalog_proto_depIdxs,
		MessageInfos:      file_pkg_apiclient_course_v1_catalog_proto_msgTypes,
	}.Build()
	File_pkg_apiclient_course_v1_catalog_proto = out.File
	file_pkg_apiclient_course_v1_catalog_proto_rawDesc = nil
	file_pkg_apiclient_course_v1_catalog_proto_goTypes = nil
	file_pkg_apiclient_course_v1_catalog_proto_depIdxs = nil
}
