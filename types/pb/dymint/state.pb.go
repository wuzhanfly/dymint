// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/dymint/state.proto

package dymint

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/tendermint/tendermint/abci/types"
	state "github.com/tendermint/tendermint/proto/tendermint/state"
	types "github.com/tendermint/tendermint/proto/tendermint/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type State struct {
	Version                          *state.Version        `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	ChainId                          string                `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	InitialHeight                    int64                 `protobuf:"varint,3,opt,name=initial_height,json=initialHeight,proto3" json:"initial_height,omitempty"`
	LastBlockHeight                  int64                 `protobuf:"varint,4,opt,name=last_block_height,json=lastBlockHeight,proto3" json:"last_block_height,omitempty"`
	LastBlockID                      types.BlockID         `protobuf:"bytes,5,opt,name=last_block_id,json=lastBlockId,proto3" json:"last_block_id"`
	LastBlockTime                    time.Time             `protobuf:"bytes,6,opt,name=last_block_time,json=lastBlockTime,proto3,stdtime" json:"last_block_time"`
	Validators                       *types.ValidatorSet   `protobuf:"bytes,9,opt,name=validators,proto3" json:"validators,omitempty"` // Deprecated: Do not use.
	LastHeightValidatorsChanged      int64                 `protobuf:"varint,11,opt,name=last_height_validators_changed,json=lastHeightValidatorsChanged,proto3" json:"last_height_validators_changed,omitempty"`
	ConsensusParams                  types.ConsensusParams `protobuf:"bytes,12,opt,name=consensus_params,json=consensusParams,proto3" json:"consensus_params"`
	LastHeightConsensusParamsChanged int64                 `protobuf:"varint,13,opt,name=last_height_consensus_params_changed,json=lastHeightConsensusParamsChanged,proto3" json:"last_height_consensus_params_changed,omitempty"`
	LastResultsHash                  []byte                `protobuf:"bytes,14,opt,name=last_results_hash,json=lastResultsHash,proto3" json:"last_results_hash,omitempty"`
	AppHash                          []byte                `protobuf:"bytes,15,opt,name=app_hash,json=appHash,proto3" json:"app_hash,omitempty"`
	LastStoreHeight                  uint64                `protobuf:"varint,16,opt,name=last_store_height,json=lastStoreHeight,proto3" json:"last_store_height,omitempty"`
	BaseHeight                       uint64                `protobuf:"varint,17,opt,name=base_height,json=baseHeight,proto3" json:"base_height,omitempty"`
	SequencerSet                     SequencerSet          `protobuf:"bytes,18,opt,name=sequencerSet,proto3" json:"sequencerSet"`
	RollappParams                    RollappParams         `protobuf:"bytes,19,opt,name=rollapp_params,json=rollappParams,proto3" json:"rollapp_params"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b679420add07272, []int{0}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_State.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return m.Size()
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetVersion() *state.Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *State) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *State) GetInitialHeight() int64 {
	if m != nil {
		return m.InitialHeight
	}
	return 0
}

func (m *State) GetLastBlockHeight() int64 {
	if m != nil {
		return m.LastBlockHeight
	}
	return 0
}

func (m *State) GetLastBlockID() types.BlockID {
	if m != nil {
		return m.LastBlockID
	}
	return types.BlockID{}
}

func (m *State) GetLastBlockTime() time.Time {
	if m != nil {
		return m.LastBlockTime
	}
	return time.Time{}
}

// Deprecated: Do not use.
func (m *State) GetValidators() *types.ValidatorSet {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *State) GetLastHeightValidatorsChanged() int64 {
	if m != nil {
		return m.LastHeightValidatorsChanged
	}
	return 0
}

func (m *State) GetConsensusParams() types.ConsensusParams {
	if m != nil {
		return m.ConsensusParams
	}
	return types.ConsensusParams{}
}

func (m *State) GetLastHeightConsensusParamsChanged() int64 {
	if m != nil {
		return m.LastHeightConsensusParamsChanged
	}
	return 0
}

func (m *State) GetLastResultsHash() []byte {
	if m != nil {
		return m.LastResultsHash
	}
	return nil
}

func (m *State) GetAppHash() []byte {
	if m != nil {
		return m.AppHash
	}
	return nil
}

func (m *State) GetLastStoreHeight() uint64 {
	if m != nil {
		return m.LastStoreHeight
	}
	return 0
}

func (m *State) GetBaseHeight() uint64 {
	if m != nil {
		return m.BaseHeight
	}
	return 0
}

func (m *State) GetSequencerSet() SequencerSet {
	if m != nil {
		return m.SequencerSet
	}
	return SequencerSet{}
}

func (m *State) GetRollappParams() RollappParams {
	if m != nil {
		return m.RollappParams
	}
	return RollappParams{}
}

//rollapp params defined in genesis and updated via gov proposal
type RollappParams struct {
	//data availability type (e.g. celestia) used in the rollapp
	Da string `protobuf:"bytes,1,opt,name=da,proto3" json:"da,omitempty"`
	//commit used for the rollapp executable
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *RollappParams) Reset()         { *m = RollappParams{} }
func (m *RollappParams) String() string { return proto.CompactTextString(m) }
func (*RollappParams) ProtoMessage()    {}
func (*RollappParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b679420add07272, []int{1}
}
func (m *RollappParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RollappParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RollappParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RollappParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollappParams.Merge(m, src)
}
func (m *RollappParams) XXX_Size() int {
	return m.Size()
}
func (m *RollappParams) XXX_DiscardUnknown() {
	xxx_messageInfo_RollappParams.DiscardUnknown(m)
}

var xxx_messageInfo_RollappParams proto.InternalMessageInfo

func (m *RollappParams) GetDa() string {
	if m != nil {
		return m.Da
	}
	return ""
}

func (m *RollappParams) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*State)(nil), "dymint.State")
	proto.RegisterType((*RollappParams)(nil), "dymint.RollappParams")
}

func init() { proto.RegisterFile("types/dymint/state.proto", fileDescriptor_4b679420add07272) }

var fileDescriptor_4b679420add07272 = []byte{
	// 685 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xcb, 0x4e, 0xdb, 0x4c,
	0x14, 0xc7, 0xe3, 0x10, 0x88, 0x33, 0x21, 0x17, 0x06, 0x3e, 0xc9, 0xf0, 0x49, 0x4e, 0xe0, 0xbb,
	0x28, 0xea, 0xc2, 0x91, 0xca, 0xaa, 0x9b, 0x56, 0x32, 0x2c, 0x08, 0x42, 0x55, 0xe5, 0x54, 0x2c,
	0xba, 0xb1, 0xc6, 0xf6, 0xd4, 0x1e, 0xd5, 0xf1, 0xb8, 0x9e, 0x09, 0x2a, 0x7d, 0x0a, 0x1e, 0xa7,
	0x8f, 0xc0, 0x92, 0x65, 0x57, 0xb4, 0x0a, 0x2f, 0x52, 0xcd, 0xc5, 0x89, 0x43, 0xc4, 0x2a, 0x99,
	0xff, 0xf9, 0xcd, 0xdf, 0x67, 0xce, 0x39, 0x33, 0xc0, 0xe2, 0xb7, 0x39, 0x66, 0xe3, 0xe8, 0x76,
	0x46, 0x32, 0x3e, 0x66, 0x1c, 0x71, 0xec, 0xe4, 0x05, 0xe5, 0x14, 0xee, 0x28, 0xed, 0xe8, 0x20,
	0xa6, 0x31, 0x95, 0xd2, 0x58, 0xfc, 0x53, 0xd1, 0xa3, 0x41, 0x4c, 0x69, 0x9c, 0xe2, 0xb1, 0x5c,
	0x05, 0xf3, 0xcf, 0x63, 0x4e, 0x66, 0x98, 0x71, 0x34, 0xcb, 0x35, 0x70, 0xac, 0x8c, 0x39, 0xce,
	0x22, 0x5c, 0x48, 0x73, 0x14, 0x84, 0x64, 0x2c, 0x55, 0x8d, 0x9c, 0x6c, 0x20, 0x5a, 0xa8, 0x30,
	0xff, 0xbf, 0xc0, 0xdc, 0xa0, 0x94, 0x44, 0x88, 0xd3, 0x42, 0x73, 0xff, 0xbc, 0xc0, 0xe5, 0xa8,
	0x40, 0xb3, 0x97, 0x3f, 0x28, 0x0f, 0xbc, 0xf6, 0xc1, 0xc3, 0xb5, 0x82, 0xa8, 0x1f, 0x15, 0x3a,
	0xf9, 0xd1, 0x04, 0xdb, 0x53, 0xb1, 0x01, 0x9e, 0x82, 0xe6, 0x0d, 0x2e, 0x18, 0xa1, 0x99, 0x65,
	0x0c, 0x8d, 0x51, 0xfb, 0xf5, 0xa1, 0xb3, 0x32, 0x75, 0x54, 0x15, 0xaf, 0x15, 0xe0, 0x95, 0x24,
	0x3c, 0x04, 0x66, 0x98, 0x20, 0x92, 0xf9, 0x24, 0xb2, 0xea, 0x43, 0x63, 0xd4, 0xf2, 0x9a, 0x72,
	0x3d, 0x89, 0xe0, 0x7f, 0xa0, 0x4b, 0x32, 0xc2, 0x09, 0x4a, 0xfd, 0x04, 0x93, 0x38, 0xe1, 0xd6,
	0xd6, 0xd0, 0x18, 0x6d, 0x79, 0x1d, 0xad, 0x5e, 0x48, 0x11, 0xbe, 0x02, 0x7b, 0x29, 0x62, 0xdc,
	0x0f, 0x52, 0x1a, 0x7e, 0x29, 0xc9, 0x86, 0x24, 0x7b, 0x22, 0xe0, 0x0a, 0x5d, 0xb3, 0x1e, 0xe8,
	0x54, 0x58, 0x12, 0x59, 0xdb, 0x9b, 0x89, 0xaa, 0x73, 0xcb, 0x5d, 0x93, 0x73, 0x77, 0xff, 0xfe,
	0x71, 0x50, 0x5b, 0x3c, 0x0e, 0xda, 0x57, 0xa5, 0xd5, 0xe4, 0xdc, 0x6b, 0x2f, 0x7d, 0x27, 0x11,
	0xbc, 0x02, 0xbd, 0x8a, 0xa7, 0xe8, 0xb8, 0xb5, 0x23, 0x5d, 0x8f, 0x1c, 0x35, 0x0e, 0x4e, 0x39,
	0x0e, 0xce, 0xc7, 0x72, 0x1c, 0x5c, 0x53, 0xd8, 0xde, 0xfd, 0x1a, 0x18, 0x5e, 0x67, 0xe9, 0x25,
	0xa2, 0xd0, 0x05, 0x60, 0xd9, 0x45, 0x66, 0xb5, 0xa4, 0x91, 0xbd, 0x99, 0xde, 0x75, 0xc9, 0x4c,
	0x31, 0x77, 0xeb, 0x96, 0xe1, 0x55, 0x76, 0xc1, 0x33, 0x60, 0xcb, 0x8c, 0x54, 0x2d, 0xfc, 0x55,
	0xc4, 0x0f, 0x13, 0x94, 0xc5, 0x38, 0xb2, 0xda, 0xb2, 0x3c, 0x7f, 0x0b, 0x4a, 0x55, 0x66, 0xe9,
	0xc7, 0xce, 0x14, 0x02, 0x3d, 0xd0, 0x0f, 0x69, 0xc6, 0x70, 0xc6, 0xe6, 0xcc, 0x57, 0x03, 0x63,
	0xed, 0xca, 0x74, 0x8e, 0x37, 0xd3, 0x39, 0x2b, 0xc9, 0x0f, 0x12, 0x74, 0x1b, 0xe2, 0x78, 0x5e,
	0x2f, 0x5c, 0x97, 0xe1, 0x7b, 0xf0, 0x6f, 0x35, 0xb1, 0xe7, 0xfe, 0xcb, 0xf4, 0x3a, 0x32, 0xbd,
	0xe1, 0x2a, 0xbd, 0x67, 0xfe, 0x65, 0x8e, 0x65, 0xeb, 0x0b, 0xcc, 0xe6, 0x29, 0x67, 0x7e, 0x82,
	0x58, 0x62, 0x75, 0x87, 0xc6, 0x68, 0x57, 0xb5, 0xde, 0x53, 0xfa, 0x05, 0x62, 0x89, 0x18, 0x34,
	0x94, 0xe7, 0x0a, 0xe9, 0x49, 0xa4, 0x89, 0xf2, 0x5c, 0x86, 0xde, 0x69, 0x1b, 0xc6, 0x69, 0x81,
	0xcb, 0x09, 0xea, 0x0f, 0x8d, 0x51, 0xc3, 0xdd, 0x5f, 0x3c, 0x0e, 0x7a, 0xa2, 0xf5, 0x53, 0x11,
	0x53, 0xc9, 0x28, 0xef, 0x8a, 0x00, 0x07, 0xa0, 0x1d, 0x20, 0xb6, 0xdc, 0xba, 0x27, 0xb6, 0x7a,
	0x40, 0x48, 0x1a, 0x78, 0x0b, 0x76, 0x19, 0xfe, 0x3a, 0xc7, 0x59, 0x88, 0x45, 0xc7, 0x2c, 0x28,
	0x0b, 0x79, 0xe0, 0xe8, 0x9b, 0x34, 0xad, 0xc4, 0x74, 0xed, 0xd6, 0x78, 0xe8, 0x82, 0x6e, 0x41,
	0xd3, 0x54, 0x1c, 0x40, 0xb7, 0x62, 0x5f, 0x3a, 0xfc, 0x55, 0x3a, 0x78, 0x2a, 0xba, 0x56, 0xfe,
	0x4e, 0x51, 0x15, 0x2f, 0x1b, 0x66, 0xb3, 0x6f, 0x5e, 0x36, 0x4c, 0xb3, 0xdf, 0xba, 0x6c, 0x98,
	0xa0, 0xdf, 0x3e, 0x79, 0x03, 0x3a, 0x6b, 0xfb, 0x60, 0x17, 0xd4, 0x23, 0x24, 0x2f, 0x6f, 0xcb,
	0xab, 0x47, 0x08, 0x5a, 0xab, 0x1b, 0xad, 0xef, 0xa6, 0x5e, 0xba, 0x17, 0xf7, 0x0b, 0xdb, 0x78,
	0x58, 0xd8, 0xc6, 0xef, 0x85, 0x6d, 0xdc, 0x3d, 0xd9, 0xb5, 0x87, 0x27, 0xbb, 0xf6, 0xf3, 0xc9,
	0xae, 0x7d, 0x72, 0x62, 0xc2, 0x93, 0x79, 0xe0, 0x84, 0x74, 0x26, 0x1e, 0x0a, 0x9c, 0x09, 0xfe,
	0xdb, 0xed, 0xf7, 0xf2, 0xf1, 0xd0, 0x2f, 0x50, 0xa0, 0xd7, 0xc1, 0x8e, 0xbc, 0x1d, 0xa7, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x4d, 0xfd, 0x88, 0x74, 0x05, 0x00, 0x00,
}

func (m *State) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *State) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *State) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.RollappParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x9a
	{
		size, err := m.SequencerSet.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x92
	if m.BaseHeight != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.BaseHeight))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if m.LastStoreHeight != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.LastStoreHeight))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if len(m.AppHash) > 0 {
		i -= len(m.AppHash)
		copy(dAtA[i:], m.AppHash)
		i = encodeVarintState(dAtA, i, uint64(len(m.AppHash)))
		i--
		dAtA[i] = 0x7a
	}
	if len(m.LastResultsHash) > 0 {
		i -= len(m.LastResultsHash)
		copy(dAtA[i:], m.LastResultsHash)
		i = encodeVarintState(dAtA, i, uint64(len(m.LastResultsHash)))
		i--
		dAtA[i] = 0x72
	}
	if m.LastHeightConsensusParamsChanged != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.LastHeightConsensusParamsChanged))
		i--
		dAtA[i] = 0x68
	}
	{
		size, err := m.ConsensusParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if m.LastHeightValidatorsChanged != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.LastHeightValidatorsChanged))
		i--
		dAtA[i] = 0x58
	}
	if m.Validators != nil {
		{
			size, err := m.Validators.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintState(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	n5, err5 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastBlockTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastBlockTime):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintState(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0x32
	{
		size, err := m.LastBlockID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.LastBlockHeight != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.LastBlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if m.InitialHeight != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.InitialHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintState(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Version != nil {
		{
			size, err := m.Version.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintState(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RollappParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RollappParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RollappParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintState(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Da) > 0 {
		i -= len(m.Da)
		copy(dAtA[i:], m.Da)
		i = encodeVarintState(dAtA, i, uint64(len(m.Da)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	offset -= sovState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *State) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != nil {
		l = m.Version.Size()
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.InitialHeight != 0 {
		n += 1 + sovState(uint64(m.InitialHeight))
	}
	if m.LastBlockHeight != 0 {
		n += 1 + sovState(uint64(m.LastBlockHeight))
	}
	l = m.LastBlockID.Size()
	n += 1 + l + sovState(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastBlockTime)
	n += 1 + l + sovState(uint64(l))
	if m.Validators != nil {
		l = m.Validators.Size()
		n += 1 + l + sovState(uint64(l))
	}
	if m.LastHeightValidatorsChanged != 0 {
		n += 1 + sovState(uint64(m.LastHeightValidatorsChanged))
	}
	l = m.ConsensusParams.Size()
	n += 1 + l + sovState(uint64(l))
	if m.LastHeightConsensusParamsChanged != 0 {
		n += 1 + sovState(uint64(m.LastHeightConsensusParamsChanged))
	}
	l = len(m.LastResultsHash)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.AppHash)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.LastStoreHeight != 0 {
		n += 2 + sovState(uint64(m.LastStoreHeight))
	}
	if m.BaseHeight != 0 {
		n += 2 + sovState(uint64(m.BaseHeight))
	}
	l = m.SequencerSet.Size()
	n += 2 + l + sovState(uint64(l))
	l = m.RollappParams.Size()
	n += 2 + l + sovState(uint64(l))
	return n
}

func (m *RollappParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Da)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func sovState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *State) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: State: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: State: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Version == nil {
				m.Version = &state.Version{}
			}
			if err := m.Version.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialHeight", wireType)
			}
			m.InitialHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockHeight", wireType)
			}
			m.LastBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastBlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastBlockID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastBlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Validators == nil {
				m.Validators = &types.ValidatorSet{}
			}
			if err := m.Validators.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastHeightValidatorsChanged", wireType)
			}
			m.LastHeightValidatorsChanged = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastHeightValidatorsChanged |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ConsensusParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastHeightConsensusParamsChanged", wireType)
			}
			m.LastHeightConsensusParamsChanged = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastHeightConsensusParamsChanged |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastResultsHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastResultsHash = append(m.LastResultsHash[:0], dAtA[iNdEx:postIndex]...)
			if m.LastResultsHash == nil {
				m.LastResultsHash = []byte{}
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppHash = append(m.AppHash[:0], dAtA[iNdEx:postIndex]...)
			if m.AppHash == nil {
				m.AppHash = []byte{}
			}
			iNdEx = postIndex
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastStoreHeight", wireType)
			}
			m.LastStoreHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastStoreHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseHeight", wireType)
			}
			m.BaseHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BaseHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequencerSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SequencerSet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 19:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RollappParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func (m *RollappParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: RollappParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RollappParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Da", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Da = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupState = fmt.Errorf("proto: unexpected end of group")
)
