// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: proto/satellite/v1/satellite.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Satellite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Lat      float64 `protobuf:"fixed64,2,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon      float64 `protobuf:"fixed64,3,opt,name=lon,proto3" json:"lon,omitempty"`
	Altitude float64 `protobuf:"fixed64,4,opt,name=altitude,proto3" json:"altitude,omitempty"`
	Velocity float64 `protobuf:"fixed64,5,opt,name=velocity,proto3" json:"velocity,omitempty"`
}

func (x *Satellite) Reset() {
	*x = Satellite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_satellite_v1_satellite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Satellite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Satellite) ProtoMessage() {}

func (x *Satellite) ProtoReflect() protoreflect.Message {
	mi := &file_proto_satellite_v1_satellite_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Satellite.ProtoReflect.Descriptor instead.
func (*Satellite) Descriptor() ([]byte, []int) {
	return file_proto_satellite_v1_satellite_proto_rawDescGZIP(), []int{0}
}

func (x *Satellite) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Satellite) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Satellite) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *Satellite) GetAltitude() float64 {
	if x != nil {
		return x.Altitude
	}
	return 0
}

func (x *Satellite) GetVelocity() float64 {
	if x != nil {
		return x.Velocity
	}
	return 0
}

type SatelliteDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CcsdsOmmVers       string `protobuf:"bytes,1,opt,name=ccsds_omm_vers,json=ccsdsOmmVers,proto3" json:"ccsds_omm_vers,omitempty"`
	Comment            string `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	CreationDate       string `protobuf:"bytes,3,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	Originator         string `protobuf:"bytes,4,opt,name=originator,proto3" json:"originator,omitempty"`
	ObjectName         string `protobuf:"bytes,5,opt,name=object_name,json=objectName,proto3" json:"object_name,omitempty"`
	ObjectId           string `protobuf:"bytes,6,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	CenterName         string `protobuf:"bytes,7,opt,name=center_name,json=centerName,proto3" json:"center_name,omitempty"`
	RefFrame           string `protobuf:"bytes,8,opt,name=ref_frame,json=refFrame,proto3" json:"ref_frame,omitempty"`
	TimeSystem         string `protobuf:"bytes,9,opt,name=time_system,json=timeSystem,proto3" json:"time_system,omitempty"`
	MeanElementTheory  string `protobuf:"bytes,10,opt,name=mean_element_theory,json=meanElementTheory,proto3" json:"mean_element_theory,omitempty"`
	Epoch              string `protobuf:"bytes,11,opt,name=epoch,proto3" json:"epoch,omitempty"`
	MeanMotion         string `protobuf:"bytes,12,opt,name=mean_motion,json=meanMotion,proto3" json:"mean_motion,omitempty"`
	Eccentricity       string `protobuf:"bytes,13,opt,name=eccentricity,proto3" json:"eccentricity,omitempty"`
	Inclination        string `protobuf:"bytes,14,opt,name=inclination,proto3" json:"inclination,omitempty"`
	RaOfAscNode        string `protobuf:"bytes,15,opt,name=ra_of_asc_node,json=raOfAscNode,proto3" json:"ra_of_asc_node,omitempty"`
	ArgOfPericenter    string `protobuf:"bytes,16,opt,name=arg_of_pericenter,json=argOfPericenter,proto3" json:"arg_of_pericenter,omitempty"`
	MeanAnomaly        string `protobuf:"bytes,17,opt,name=mean_anomaly,json=meanAnomaly,proto3" json:"mean_anomaly,omitempty"`
	EphemerisType      string `protobuf:"bytes,18,opt,name=ephemeris_type,json=ephemerisType,proto3" json:"ephemeris_type,omitempty"`
	ClassificationType string `protobuf:"bytes,19,opt,name=classification_type,json=classificationType,proto3" json:"classification_type,omitempty"`
	NoradCatId         string `protobuf:"bytes,20,opt,name=norad_cat_id,json=noradCatId,proto3" json:"norad_cat_id,omitempty"`
	ElementSetNo       string `protobuf:"bytes,21,opt,name=element_set_no,json=elementSetNo,proto3" json:"element_set_no,omitempty"`
	RevAtEpoch         string `protobuf:"bytes,22,opt,name=rev_at_epoch,json=revAtEpoch,proto3" json:"rev_at_epoch,omitempty"`
	Bstar              string `protobuf:"bytes,23,opt,name=bstar,proto3" json:"bstar,omitempty"`
	MeanMotionDot      string `protobuf:"bytes,24,opt,name=mean_motion_dot,json=meanMotionDot,proto3" json:"mean_motion_dot,omitempty"`
	MeanMotionDdot     string `protobuf:"bytes,25,opt,name=mean_motion_ddot,json=meanMotionDdot,proto3" json:"mean_motion_ddot,omitempty"`
	SemimajorAxis      string `protobuf:"bytes,26,opt,name=semimajor_axis,json=semimajorAxis,proto3" json:"semimajor_axis,omitempty"`
	Period             string `protobuf:"bytes,27,opt,name=period,proto3" json:"period,omitempty"`
	Apoapsis           string `protobuf:"bytes,28,opt,name=apoapsis,proto3" json:"apoapsis,omitempty"`
	Periapsis          string `protobuf:"bytes,29,opt,name=periapsis,proto3" json:"periapsis,omitempty"`
	ObjectType         string `protobuf:"bytes,30,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`
	RcsSize            string `protobuf:"bytes,31,opt,name=rcs_size,json=rcsSize,proto3" json:"rcs_size,omitempty"`
	CountryCode        string `protobuf:"bytes,32,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	LaunchDate         string `protobuf:"bytes,33,opt,name=launch_date,json=launchDate,proto3" json:"launch_date,omitempty"`
	Site               string `protobuf:"bytes,34,opt,name=site,proto3" json:"site,omitempty"`
	DecayDate          string `protobuf:"bytes,35,opt,name=decay_date,json=decayDate,proto3" json:"decay_date,omitempty"`
	File               string `protobuf:"bytes,36,opt,name=file,proto3" json:"file,omitempty"`
	GpId               string `protobuf:"bytes,37,opt,name=gp_id,json=gpId,proto3" json:"gp_id,omitempty"`
	TleLine0           string `protobuf:"bytes,38,opt,name=tle_line0,json=tleLine0,proto3" json:"tle_line0,omitempty"`
	TleLine1           string `protobuf:"bytes,39,opt,name=tle_line1,json=tleLine1,proto3" json:"tle_line1,omitempty"`
	TleLine2           string `protobuf:"bytes,40,opt,name=tle_line2,json=tleLine2,proto3" json:"tle_line2,omitempty"`
}

func (x *SatelliteDetail) Reset() {
	*x = SatelliteDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_satellite_v1_satellite_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SatelliteDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SatelliteDetail) ProtoMessage() {}

func (x *SatelliteDetail) ProtoReflect() protoreflect.Message {
	mi := &file_proto_satellite_v1_satellite_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SatelliteDetail.ProtoReflect.Descriptor instead.
func (*SatelliteDetail) Descriptor() ([]byte, []int) {
	return file_proto_satellite_v1_satellite_proto_rawDescGZIP(), []int{1}
}

func (x *SatelliteDetail) GetCcsdsOmmVers() string {
	if x != nil {
		return x.CcsdsOmmVers
	}
	return ""
}

func (x *SatelliteDetail) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *SatelliteDetail) GetCreationDate() string {
	if x != nil {
		return x.CreationDate
	}
	return ""
}

func (x *SatelliteDetail) GetOriginator() string {
	if x != nil {
		return x.Originator
	}
	return ""
}

func (x *SatelliteDetail) GetObjectName() string {
	if x != nil {
		return x.ObjectName
	}
	return ""
}

func (x *SatelliteDetail) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *SatelliteDetail) GetCenterName() string {
	if x != nil {
		return x.CenterName
	}
	return ""
}

func (x *SatelliteDetail) GetRefFrame() string {
	if x != nil {
		return x.RefFrame
	}
	return ""
}

func (x *SatelliteDetail) GetTimeSystem() string {
	if x != nil {
		return x.TimeSystem
	}
	return ""
}

func (x *SatelliteDetail) GetMeanElementTheory() string {
	if x != nil {
		return x.MeanElementTheory
	}
	return ""
}

func (x *SatelliteDetail) GetEpoch() string {
	if x != nil {
		return x.Epoch
	}
	return ""
}

func (x *SatelliteDetail) GetMeanMotion() string {
	if x != nil {
		return x.MeanMotion
	}
	return ""
}

func (x *SatelliteDetail) GetEccentricity() string {
	if x != nil {
		return x.Eccentricity
	}
	return ""
}

func (x *SatelliteDetail) GetInclination() string {
	if x != nil {
		return x.Inclination
	}
	return ""
}

func (x *SatelliteDetail) GetRaOfAscNode() string {
	if x != nil {
		return x.RaOfAscNode
	}
	return ""
}

func (x *SatelliteDetail) GetArgOfPericenter() string {
	if x != nil {
		return x.ArgOfPericenter
	}
	return ""
}

func (x *SatelliteDetail) GetMeanAnomaly() string {
	if x != nil {
		return x.MeanAnomaly
	}
	return ""
}

func (x *SatelliteDetail) GetEphemerisType() string {
	if x != nil {
		return x.EphemerisType
	}
	return ""
}

func (x *SatelliteDetail) GetClassificationType() string {
	if x != nil {
		return x.ClassificationType
	}
	return ""
}

func (x *SatelliteDetail) GetNoradCatId() string {
	if x != nil {
		return x.NoradCatId
	}
	return ""
}

func (x *SatelliteDetail) GetElementSetNo() string {
	if x != nil {
		return x.ElementSetNo
	}
	return ""
}

func (x *SatelliteDetail) GetRevAtEpoch() string {
	if x != nil {
		return x.RevAtEpoch
	}
	return ""
}

func (x *SatelliteDetail) GetBstar() string {
	if x != nil {
		return x.Bstar
	}
	return ""
}

func (x *SatelliteDetail) GetMeanMotionDot() string {
	if x != nil {
		return x.MeanMotionDot
	}
	return ""
}

func (x *SatelliteDetail) GetMeanMotionDdot() string {
	if x != nil {
		return x.MeanMotionDdot
	}
	return ""
}

func (x *SatelliteDetail) GetSemimajorAxis() string {
	if x != nil {
		return x.SemimajorAxis
	}
	return ""
}

func (x *SatelliteDetail) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *SatelliteDetail) GetApoapsis() string {
	if x != nil {
		return x.Apoapsis
	}
	return ""
}

func (x *SatelliteDetail) GetPeriapsis() string {
	if x != nil {
		return x.Periapsis
	}
	return ""
}

func (x *SatelliteDetail) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

func (x *SatelliteDetail) GetRcsSize() string {
	if x != nil {
		return x.RcsSize
	}
	return ""
}

func (x *SatelliteDetail) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *SatelliteDetail) GetLaunchDate() string {
	if x != nil {
		return x.LaunchDate
	}
	return ""
}

func (x *SatelliteDetail) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *SatelliteDetail) GetDecayDate() string {
	if x != nil {
		return x.DecayDate
	}
	return ""
}

func (x *SatelliteDetail) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *SatelliteDetail) GetGpId() string {
	if x != nil {
		return x.GpId
	}
	return ""
}

func (x *SatelliteDetail) GetTleLine0() string {
	if x != nil {
		return x.TleLine0
	}
	return ""
}

func (x *SatelliteDetail) GetTleLine1() string {
	if x != nil {
		return x.TleLine1
	}
	return ""
}

func (x *SatelliteDetail) GetTleLine2() string {
	if x != nil {
		return x.TleLine2
	}
	return ""
}

type GetSatellitePositionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *GetSatellitePositionsRequest) Reset() {
	*x = GetSatellitePositionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_satellite_v1_satellite_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSatellitePositionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSatellitePositionsRequest) ProtoMessage() {}

func (x *GetSatellitePositionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_satellite_v1_satellite_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSatellitePositionsRequest.ProtoReflect.Descriptor instead.
func (*GetSatellitePositionsRequest) Descriptor() ([]byte, []int) {
	return file_proto_satellite_v1_satellite_proto_rawDescGZIP(), []int{2}
}

func (x *GetSatellitePositionsRequest) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type GetSatellitePositionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Satellites []*Satellite `protobuf:"bytes,1,rep,name=satellites,proto3" json:"satellites,omitempty"`
}

func (x *GetSatellitePositionsResponse) Reset() {
	*x = GetSatellitePositionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_satellite_v1_satellite_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSatellitePositionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSatellitePositionsResponse) ProtoMessage() {}

func (x *GetSatellitePositionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_satellite_v1_satellite_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSatellitePositionsResponse.ProtoReflect.Descriptor instead.
func (*GetSatellitePositionsResponse) Descriptor() ([]byte, []int) {
	return file_proto_satellite_v1_satellite_proto_rawDescGZIP(), []int{3}
}

func (x *GetSatellitePositionsResponse) GetSatellites() []*Satellite {
	if x != nil {
		return x.Satellites
	}
	return nil
}

var File_proto_satellite_v1_satellite_proto protoreflect.FileDescriptor

var file_proto_satellite_v1_satellite_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x77, 0x0a, 0x09, 0x53, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x08, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x22, 0xaa, 0x0a, 0x0a, 0x0f, 0x53,
	0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x24,
	0x0a, 0x0e, 0x63, 0x63, 0x73, 0x64, 0x73, 0x5f, 0x6f, 0x6d, 0x6d, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x63, 0x73, 0x64, 0x73, 0x4f, 0x6d, 0x6d,
	0x56, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x66, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x65, 0x61, 0x6e, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x68, 0x65, 0x6f, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d,
	0x65, 0x61, 0x6e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x68, 0x65, 0x6f, 0x72, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x61, 0x6e, 0x5f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x61,
	0x6e, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x63, 0x63, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65,
	0x63, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x63, 0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x69,
	0x6e, 0x63, 0x6c, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x69, 0x6e, 0x63, 0x6c, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a,
	0x0e, 0x72, 0x61, 0x5f, 0x6f, 0x66, 0x5f, 0x61, 0x73, 0x63, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x61, 0x4f, 0x66, 0x41, 0x73, 0x63, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x72, 0x67, 0x5f, 0x6f, 0x66, 0x5f, 0x70, 0x65, 0x72,
	0x69, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61,
	0x72, 0x67, 0x4f, 0x66, 0x50, 0x65, 0x72, 0x69, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x65, 0x61, 0x6e, 0x5f, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x61, 0x6e, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c,
	0x79, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x69, 0x73, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x70, 0x68, 0x65, 0x6d,
	0x65, 0x72, 0x69, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x6e, 0x6f, 0x72,
	0x61, 0x64, 0x5f, 0x63, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6e, 0x6f, 0x72, 0x61, 0x64, 0x43, 0x61, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x6f, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x4e,
	0x6f, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x65, 0x76, 0x5f, 0x61, 0x74, 0x5f, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x76, 0x41, 0x74, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x73, 0x74, 0x61, 0x72, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x62, 0x73, 0x74, 0x61, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x65, 0x61,
	0x6e, 0x5f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x6f, 0x74, 0x18, 0x18, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x61, 0x6e, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f,
	0x74, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x65, 0x61, 0x6e, 0x5f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x64, 0x64, 0x6f, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x65, 0x61,
	0x6e, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x64, 0x6f, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73,
	0x65, 0x6d, 0x69, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x18, 0x1a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x6d, 0x69, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x41, 0x78,
	0x69, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x1b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70,
	0x6f, 0x61, 0x70, 0x73, 0x69, 0x73, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70,
	0x6f, 0x61, 0x70, 0x73, 0x69, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x69, 0x61, 0x70,
	0x73, 0x69, 0x73, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x65, 0x72, 0x69, 0x61,
	0x70, 0x73, 0x69, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x63, 0x73, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x63, 0x73, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x18, 0x22, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x63, 0x61,
	0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65,
	0x63, 0x61, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x24, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x67,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x25, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x70, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x30, 0x18, 0x26, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x30, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x31, 0x18, 0x27, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x31, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6c,
	0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x32, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x32, 0x22, 0x4e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x53, 0x61,
	0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x58, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x53, 0x61,
	0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x73, 0x61, 0x74, 0x65,
	0x6c, 0x6c, 0x69, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73,
	0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x74, 0x65,
	0x6c, 0x6c, 0x69, 0x74, 0x65, 0x52, 0x0a, 0x73, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65,
	0x73, 0x32, 0x97, 0x02, 0x0a, 0x10, 0x53, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x91, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x61,
	0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x2a, 0x2e, 0x73, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x73,
	0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65,
	0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x6f, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x53, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x17, 0x2e, 0x73, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x1a, 0x1d, 0x2e, 0x73, 0x61, 0x74, 0x65,
	0x6c, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69,
	0x74, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b,
	0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x74, 0x65, 0x2f,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x40, 0x5a, 0x3e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x73, 0x75, 0x6b, 0x6f, 0x79,
	0x61, 0x63, 0x68, 0x69, 0x2f, 0x72, 0x65, 0x61, 0x63, 0x74, 0x2d, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x2d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2d, 0x73, 0x61, 0x74, 0x65, 0x6c, 0x6c,
	0x69, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_satellite_v1_satellite_proto_rawDescOnce sync.Once
	file_proto_satellite_v1_satellite_proto_rawDescData = file_proto_satellite_v1_satellite_proto_rawDesc
)

func file_proto_satellite_v1_satellite_proto_rawDescGZIP() []byte {
	file_proto_satellite_v1_satellite_proto_rawDescOnce.Do(func() {
		file_proto_satellite_v1_satellite_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_satellite_v1_satellite_proto_rawDescData)
	})
	return file_proto_satellite_v1_satellite_proto_rawDescData
}

var file_proto_satellite_v1_satellite_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_satellite_v1_satellite_proto_goTypes = []interface{}{
	(*Satellite)(nil),                     // 0: satellite.v1.Satellite
	(*SatelliteDetail)(nil),               // 1: satellite.v1.SatelliteDetail
	(*GetSatellitePositionsRequest)(nil),  // 2: satellite.v1.GetSatellitePositionsRequest
	(*GetSatellitePositionsResponse)(nil), // 3: satellite.v1.GetSatellitePositionsResponse
	(*timestamppb.Timestamp)(nil),         // 4: google.protobuf.Timestamp
}
var file_proto_satellite_v1_satellite_proto_depIdxs = []int32{
	4, // 0: satellite.v1.GetSatellitePositionsRequest.time:type_name -> google.protobuf.Timestamp
	0, // 1: satellite.v1.GetSatellitePositionsResponse.satellites:type_name -> satellite.v1.Satellite
	2, // 2: satellite.v1.SatelliteService.GetSatellitePositions:input_type -> satellite.v1.GetSatellitePositionsRequest
	0, // 3: satellite.v1.SatelliteService.GetSatelliteDetail:input_type -> satellite.v1.Satellite
	3, // 4: satellite.v1.SatelliteService.GetSatellitePositions:output_type -> satellite.v1.GetSatellitePositionsResponse
	1, // 5: satellite.v1.SatelliteService.GetSatelliteDetail:output_type -> satellite.v1.SatelliteDetail
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_satellite_v1_satellite_proto_init() }
func file_proto_satellite_v1_satellite_proto_init() {
	if File_proto_satellite_v1_satellite_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_satellite_v1_satellite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Satellite); i {
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
		file_proto_satellite_v1_satellite_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SatelliteDetail); i {
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
		file_proto_satellite_v1_satellite_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSatellitePositionsRequest); i {
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
		file_proto_satellite_v1_satellite_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSatellitePositionsResponse); i {
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
			RawDescriptor: file_proto_satellite_v1_satellite_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_satellite_v1_satellite_proto_goTypes,
		DependencyIndexes: file_proto_satellite_v1_satellite_proto_depIdxs,
		MessageInfos:      file_proto_satellite_v1_satellite_proto_msgTypes,
	}.Build()
	File_proto_satellite_v1_satellite_proto = out.File
	file_proto_satellite_v1_satellite_proto_rawDesc = nil
	file_proto_satellite_v1_satellite_proto_goTypes = nil
	file_proto_satellite_v1_satellite_proto_depIdxs = nil
}
