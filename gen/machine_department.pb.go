// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: machine_department.proto

package ssov1

import (
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_machine_department_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_machine_department_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_machine_department_proto_rawDescGZIP(), []int{0}
}

type UserName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
}

func (x *UserName) Reset() {
	*x = UserName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_machine_department_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserName) ProtoMessage() {}

func (x *UserName) ProtoReflect() protoreflect.Message {
	mi := &file_machine_department_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserName.ProtoReflect.Descriptor instead.
func (*UserName) Descriptor() ([]byte, []int) {
	return file_machine_department_proto_rawDescGZIP(), []int{1}
}

func (x *UserName) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type InfoMachineDep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Engine1        *Engine        `protobuf:"bytes,1,opt,name=Engine1,proto3" json:"Engine1,omitempty"`
	Engine2        *Engine        `protobuf:"bytes,2,opt,name=Engine2,proto3" json:"Engine2,omitempty"`
	CoolingSystem1 *CoolingSystem `protobuf:"bytes,3,opt,name=CoolingSystem1,proto3" json:"CoolingSystem1,omitempty"`
	CoolingSystem2 *CoolingSystem `protobuf:"bytes,4,opt,name=CoolingSystem2,proto3" json:"CoolingSystem2,omitempty"`
	Generator1     *Generator     `protobuf:"bytes,5,opt,name=Generator1,proto3" json:"Generator1,omitempty"`
	Generator2     *Generator     `protobuf:"bytes,6,opt,name=Generator2,proto3" json:"Generator2,omitempty"`
	FuelSystem1    *FuelSystem    `protobuf:"bytes,7,opt,name=FuelSystem1,proto3" json:"FuelSystem1,omitempty"`
	FuelSystem2    *FuelSystem    `protobuf:"bytes,8,opt,name=FuelSystem2,proto3" json:"FuelSystem2,omitempty"`
}

func (x *InfoMachineDep) Reset() {
	*x = InfoMachineDep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_machine_department_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoMachineDep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoMachineDep) ProtoMessage() {}

func (x *InfoMachineDep) ProtoReflect() protoreflect.Message {
	mi := &file_machine_department_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoMachineDep.ProtoReflect.Descriptor instead.
func (*InfoMachineDep) Descriptor() ([]byte, []int) {
	return file_machine_department_proto_rawDescGZIP(), []int{2}
}

func (x *InfoMachineDep) GetEngine1() *Engine {
	if x != nil {
		return x.Engine1
	}
	return nil
}

func (x *InfoMachineDep) GetEngine2() *Engine {
	if x != nil {
		return x.Engine2
	}
	return nil
}

func (x *InfoMachineDep) GetCoolingSystem1() *CoolingSystem {
	if x != nil {
		return x.CoolingSystem1
	}
	return nil
}

func (x *InfoMachineDep) GetCoolingSystem2() *CoolingSystem {
	if x != nil {
		return x.CoolingSystem2
	}
	return nil
}

func (x *InfoMachineDep) GetGenerator1() *Generator {
	if x != nil {
		return x.Generator1
	}
	return nil
}

func (x *InfoMachineDep) GetGenerator2() *Generator {
	if x != nil {
		return x.Generator2
	}
	return nil
}

func (x *InfoMachineDep) GetFuelSystem1() *FuelSystem {
	if x != nil {
		return x.FuelSystem1
	}
	return nil
}

func (x *InfoMachineDep) GetFuelSystem2() *FuelSystem {
	if x != nil {
		return x.FuelSystem2
	}
	return nil
}

type Engine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status      string  `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Rpm         int64   `protobuf:"varint,3,opt,name=rpm,proto3" json:"rpm,omitempty"`
	Temperature float32 `protobuf:"fixed32,4,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Voltage     float32 `protobuf:"fixed32,5,opt,name=voltage,proto3" json:"voltage,omitempty"`
}

func (x *Engine) Reset() {
	*x = Engine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_machine_department_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Engine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Engine) ProtoMessage() {}

func (x *Engine) ProtoReflect() protoreflect.Message {
	mi := &file_machine_department_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Engine.ProtoReflect.Descriptor instead.
func (*Engine) Descriptor() ([]byte, []int) {
	return file_machine_department_proto_rawDescGZIP(), []int{3}
}

func (x *Engine) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Engine) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Engine) GetRpm() int64 {
	if x != nil {
		return x.Rpm
	}
	return 0
}

func (x *Engine) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *Engine) GetVoltage() float32 {
	if x != nil {
		return x.Voltage
	}
	return 0
}

type CoolingSystem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status             string  `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	CoolantTemperature float32 `protobuf:"fixed32,3,opt,name=coolant_temperature,json=coolantTemperature,proto3" json:"coolant_temperature,omitempty"`
	SystemPressure     float32 `protobuf:"fixed32,4,opt,name=system_pressure,json=systemPressure,proto3" json:"system_pressure,omitempty"`
}

func (x *CoolingSystem) Reset() {
	*x = CoolingSystem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_machine_department_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoolingSystem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoolingSystem) ProtoMessage() {}

func (x *CoolingSystem) ProtoReflect() protoreflect.Message {
	mi := &file_machine_department_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoolingSystem.ProtoReflect.Descriptor instead.
func (*CoolingSystem) Descriptor() ([]byte, []int) {
	return file_machine_department_proto_rawDescGZIP(), []int{4}
}

func (x *CoolingSystem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CoolingSystem) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CoolingSystem) GetCoolantTemperature() float32 {
	if x != nil {
		return x.CoolantTemperature
	}
	return 0
}

func (x *CoolingSystem) GetSystemPressure() float32 {
	if x != nil {
		return x.SystemPressure
	}
	return 0
}

type Generator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status  string  `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Power   float32 `protobuf:"fixed32,3,opt,name=power,proto3" json:"power,omitempty"`
	Fuel    string  `protobuf:"bytes,4,opt,name=fuel,proto3" json:"fuel,omitempty"`
	Voltage float32 `protobuf:"fixed32,5,opt,name=voltage,proto3" json:"voltage,omitempty"`
}

func (x *Generator) Reset() {
	*x = Generator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_machine_department_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Generator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Generator) ProtoMessage() {}

func (x *Generator) ProtoReflect() protoreflect.Message {
	mi := &file_machine_department_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Generator.ProtoReflect.Descriptor instead.
func (*Generator) Descriptor() ([]byte, []int) {
	return file_machine_department_proto_rawDescGZIP(), []int{5}
}

func (x *Generator) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Generator) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Generator) GetPower() float32 {
	if x != nil {
		return x.Power
	}
	return 0
}

func (x *Generator) GetFuel() string {
	if x != nil {
		return x.Fuel
	}
	return ""
}

func (x *Generator) GetVoltage() float32 {
	if x != nil {
		return x.Voltage
	}
	return 0
}

type FuelSystem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status            string  `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	FuelLevel         float32 `protobuf:"fixed32,3,opt,name=fuel_level,json=fuelLevel,proto3" json:"fuel_level,omitempty"`
	ContaminantsLevel float32 `protobuf:"fixed32,4,opt,name=contaminants_level,json=contaminantsLevel,proto3" json:"contaminants_level,omitempty"`
	FuelFilterStatus  string  `protobuf:"bytes,5,opt,name=fuel_filter_status,json=fuelFilterStatus,proto3" json:"fuel_filter_status,omitempty"`
	FlowRate          float32 `protobuf:"fixed32,6,opt,name=flow_rate,json=flowRate,proto3" json:"flow_rate,omitempty"`
	LeakDetection     bool    `protobuf:"varint,7,opt,name=leak_detection,json=leakDetection,proto3" json:"leak_detection,omitempty"`
	FuelPumpStatus    string  `protobuf:"bytes,8,opt,name=fuel_pump_status,json=fuelPumpStatus,proto3" json:"fuel_pump_status,omitempty"`
}

func (x *FuelSystem) Reset() {
	*x = FuelSystem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_machine_department_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FuelSystem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FuelSystem) ProtoMessage() {}

func (x *FuelSystem) ProtoReflect() protoreflect.Message {
	mi := &file_machine_department_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FuelSystem.ProtoReflect.Descriptor instead.
func (*FuelSystem) Descriptor() ([]byte, []int) {
	return file_machine_department_proto_rawDescGZIP(), []int{6}
}

func (x *FuelSystem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FuelSystem) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *FuelSystem) GetFuelLevel() float32 {
	if x != nil {
		return x.FuelLevel
	}
	return 0
}

func (x *FuelSystem) GetContaminantsLevel() float32 {
	if x != nil {
		return x.ContaminantsLevel
	}
	return 0
}

func (x *FuelSystem) GetFuelFilterStatus() string {
	if x != nil {
		return x.FuelFilterStatus
	}
	return ""
}

func (x *FuelSystem) GetFlowRate() float32 {
	if x != nil {
		return x.FlowRate
	}
	return 0
}

func (x *FuelSystem) GetLeakDetection() bool {
	if x != nil {
		return x.LeakDetection
	}
	return false
}

func (x *FuelSystem) GetFuelPumpStatus() string {
	if x != nil {
		return x.FuelPumpStatus
	}
	return ""
}

var File_machine_department_proto protoreflect.FileDescriptor

var file_machine_department_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x79, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x26,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xcc, 0x03, 0x0a, 0x0e, 0x49, 0x6e, 0x66, 0x6f, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x44, 0x65, 0x70, 0x12, 0x2b, 0x0a, 0x07, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x79, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x07, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x31, 0x12, 0x2b, 0x0a, 0x07, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x07, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x32, 0x12, 0x40, 0x0a, 0x0e, 0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x79,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x0e, 0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x31, 0x12, 0x40, 0x0a, 0x0e, 0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e,
	0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x0e, 0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x32, 0x12, 0x34, 0x0a, 0x0a, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x31, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x79,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x0a, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x31, 0x12, 0x34, 0x0a,
	0x0a, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x32, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x32, 0x12, 0x37, 0x0a, 0x0b, 0x46, 0x75, 0x65, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x31, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x75, 0x65, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52,
	0x0b, 0x46, 0x75, 0x65, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x31, 0x12, 0x37, 0x0a, 0x0b,
	0x46, 0x75, 0x65, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x32, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x75,
	0x65, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x46, 0x75, 0x65, 0x6c, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x32, 0x22, 0x7e, 0x0a, 0x06, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x70, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x72, 0x70, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x76, 0x6f,
	0x6c, 0x74, 0x61, 0x67, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e,
	0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x2f, 0x0a, 0x13, 0x63, 0x6f, 0x6f, 0x6c, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x63, 0x6f,
	0x6f, 0x6c, 0x61, 0x6e, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x27, 0x0a, 0x0f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x22, 0x77, 0x0a, 0x09, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70,
	0x6f, 0x77, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x75, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x75, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x6f, 0x6c, 0x74,
	0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x76, 0x6f, 0x6c, 0x74, 0x61,
	0x67, 0x65, 0x22, 0x9e, 0x02, 0x0a, 0x0a, 0x46, 0x75, 0x65, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x75, 0x65,
	0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x66,
	0x75, 0x65, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x6d, 0x69, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x6e,
	0x74, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x66, 0x75, 0x65, 0x6c, 0x5f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x75, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x61,
	0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x61,
	0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x65, 0x61, 0x6b, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6c, 0x65, 0x61, 0x6b,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x75, 0x65,
	0x6c, 0x5f, 0x70, 0x75, 0x6d, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x75, 0x65, 0x6c, 0x50, 0x75, 0x6d, 0x70, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x32, 0xda, 0x02, 0x0a, 0x11, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x44,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x44, 0x65, 0x70, 0x12, 0x13,
	0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x1a, 0x19, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x44, 0x65, 0x70, 0x30, 0x01,
	0x12, 0x37, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x12, 0x11, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x1a, 0x10, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x17, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x12, 0x18, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x10,
	0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x3d, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x1a, 0x10, 0x2e,
	0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x3f, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x75, 0x65,
	0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x15, 0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x46, 0x75, 0x65, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x10,
	0x2e, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x3b, 0x73, 0x73, 0x6f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_machine_department_proto_rawDescOnce sync.Once
	file_machine_department_proto_rawDescData = file_machine_department_proto_rawDesc
)

func file_machine_department_proto_rawDescGZIP() []byte {
	file_machine_department_proto_rawDescOnce.Do(func() {
		file_machine_department_proto_rawDescData = protoimpl.X.CompressGZIP(file_machine_department_proto_rawDescData)
	})
	return file_machine_department_proto_rawDescData
}

var file_machine_department_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_machine_department_proto_goTypes = []interface{}{
	(*Empty)(nil),          // 0: myservice.Empty
	(*UserName)(nil),       // 1: myservice.userName
	(*InfoMachineDep)(nil), // 2: myservice.InfoMachineDep
	(*Engine)(nil),         // 3: myservice.Engine
	(*CoolingSystem)(nil),  // 4: myservice.CoolingSystem
	(*Generator)(nil),      // 5: myservice.Generator
	(*FuelSystem)(nil),     // 6: myservice.FuelSystem
}
var file_machine_department_proto_depIdxs = []int32{
	3,  // 0: myservice.InfoMachineDep.Engine1:type_name -> myservice.Engine
	3,  // 1: myservice.InfoMachineDep.Engine2:type_name -> myservice.Engine
	4,  // 2: myservice.InfoMachineDep.CoolingSystem1:type_name -> myservice.CoolingSystem
	4,  // 3: myservice.InfoMachineDep.CoolingSystem2:type_name -> myservice.CoolingSystem
	5,  // 4: myservice.InfoMachineDep.Generator1:type_name -> myservice.Generator
	5,  // 5: myservice.InfoMachineDep.Generator2:type_name -> myservice.Generator
	6,  // 6: myservice.InfoMachineDep.FuelSystem1:type_name -> myservice.FuelSystem
	6,  // 7: myservice.InfoMachineDep.FuelSystem2:type_name -> myservice.FuelSystem
	1,  // 8: myservice.MachineDepartment.GetInfoMachineDep:input_type -> myservice.userName
	3,  // 9: myservice.MachineDepartment.ChangeInfoEngine:input_type -> myservice.Engine
	4,  // 10: myservice.MachineDepartment.ChangeInfoCoolingSystem:input_type -> myservice.CoolingSystem
	5,  // 11: myservice.MachineDepartment.ChangeInfoGenerator:input_type -> myservice.Generator
	6,  // 12: myservice.MachineDepartment.ChangeInfoFuelSystem:input_type -> myservice.FuelSystem
	2,  // 13: myservice.MachineDepartment.GetInfoMachineDep:output_type -> myservice.InfoMachineDep
	0,  // 14: myservice.MachineDepartment.ChangeInfoEngine:output_type -> myservice.Empty
	0,  // 15: myservice.MachineDepartment.ChangeInfoCoolingSystem:output_type -> myservice.Empty
	0,  // 16: myservice.MachineDepartment.ChangeInfoGenerator:output_type -> myservice.Empty
	0,  // 17: myservice.MachineDepartment.ChangeInfoFuelSystem:output_type -> myservice.Empty
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_machine_department_proto_init() }
func file_machine_department_proto_init() {
	if File_machine_department_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_machine_department_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_machine_department_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserName); i {
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
		file_machine_department_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoMachineDep); i {
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
		file_machine_department_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Engine); i {
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
		file_machine_department_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoolingSystem); i {
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
		file_machine_department_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Generator); i {
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
		file_machine_department_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FuelSystem); i {
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
			RawDescriptor: file_machine_department_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_machine_department_proto_goTypes,
		DependencyIndexes: file_machine_department_proto_depIdxs,
		MessageInfos:      file_machine_department_proto_msgTypes,
	}.Build()
	File_machine_department_proto = out.File
	file_machine_department_proto_rawDesc = nil
	file_machine_department_proto_goTypes = nil
	file_machine_department_proto_depIdxs = nil
}
