﻿// Code generated by protoc-gen-go. DO NOT EDIT.
// source: exchange.proto

/*
Package exchange is a generated protocol buffer package.

It is generated from these files:
	exchange.proto

It has these top-level messages:
	PushRequest
	BuyAssetResponse
	QueryRequest
	GetTxListResponse
	QueryTxData
	TxRow
	TransferRequest
	TransferResponse
	QueryTransferRequest
	QueryTransferResponse
	IsBuyAssetRequest
	IsBuyAssetResponse
*/
package exchange

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

type PushRequest struct {
	Version     uint32 `protobuf:"varint,1,opt,name=version" json:"version"`
	CursorNum   uint32 `protobuf:"varint,2,opt,name=cursor_num,json=cursorNum" json:"cursor_num"`
	CursorLabel uint32 `protobuf:"varint,3,opt,name=cursor_label,json=cursorLabel" json:"cursor_label"`
	Lifetime    uint64 `protobuf:"varint,4,opt,name=lifetime" json:"lifetime"`
	Sender      string `protobuf:"bytes,5,opt,name=sender" json:"sender"`
	Contract    string `protobuf:"bytes,6,opt,name=contract" json:"contract"`
	Method      string `protobuf:"bytes,7,opt,name=method" json:"method"`
	Param       string `protobuf:"bytes,8,opt,name=param" json:"param"`
	SigAlg      uint32 `protobuf:"varint,9,opt,name=sig_alg,json=sigAlg" json:"sig_alg"`
	Signature   string `protobuf:"bytes,10,opt,name=signature" json:"signature"`
}

func (m *PushRequest) Reset()                    { *m = PushRequest{} }
func (m *PushRequest) String() string            { return proto.CompactTextString(m) }
func (*PushRequest) ProtoMessage()               {}
func (*PushRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PushRequest) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PushRequest) GetCursorNum() uint32 {
	if m != nil {
		return m.CursorNum
	}
	return 0
}

func (m *PushRequest) GetCursorLabel() uint32 {
	if m != nil {
		return m.CursorLabel
	}
	return 0
}

func (m *PushRequest) GetLifetime() uint64 {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *PushRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *PushRequest) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *PushRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *PushRequest) GetParam() string {
	if m != nil {
		return m.Param
	}
	return ""
}

func (m *PushRequest) GetSigAlg() uint32 {
	if m != nil {
		return m.SigAlg
	}
	return 0
}

func (m *PushRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type BuyAssetResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Data string `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *BuyAssetResponse) Reset()                    { *m = BuyAssetResponse{} }
func (m *BuyAssetResponse) String() string            { return proto.CompactTextString(m) }
func (*BuyAssetResponse) ProtoMessage()               {}
func (*BuyAssetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BuyAssetResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BuyAssetResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *BuyAssetResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type QueryRequest struct {
	PageSize  uint32 `protobuf:"varint,1,opt,name=page_size,json=pageSize" json:"page_size"`
	PageNum   uint32 `protobuf:"varint,2,opt,name=page_num,json=pageNum" json:"page_num"`
	Username  string `protobuf:"bytes,3,opt,name=username" json:"username"`
	Random    string `protobuf:"bytes,4,opt,name=random" json:"random"`
	Signature string `protobuf:"bytes,5,opt,name=signature" json:"signature"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *QueryRequest) GetPageSize() uint32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *QueryRequest) GetPageNum() uint32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *QueryRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *QueryRequest) GetRandom() string {
	if m != nil {
		return m.Random
	}
	return ""
}

func (m *QueryRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type GetTxListResponse struct {
	Code uint32       `protobuf:"varint,1,opt,name=code" json:"code"`
	Data *QueryTxData `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string       `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *GetTxListResponse) Reset()                    { *m = GetTxListResponse{} }
func (m *GetTxListResponse) String() string            { return proto.CompactTextString(m) }
func (*GetTxListResponse) ProtoMessage()               {}
func (*GetTxListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetTxListResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetTxListResponse) GetData() *QueryTxData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetTxListResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type QueryTxData struct {
	PageNum  uint64   `protobuf:"varint,1,opt,name=pageNum" json:"pageNum"`
	RowCount uint64   `protobuf:"varint,2,opt,name=rowCount" json:"rowCount"`
	Row      []*TxRow `protobuf:"bytes,3,rep,name=row" json:"row"`
}

func (m *QueryTxData) Reset()                    { *m = QueryTxData{} }
func (m *QueryTxData) String() string            { return proto.CompactTextString(m) }
func (*QueryTxData) ProtoMessage()               {}
func (*QueryTxData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *QueryTxData) GetPageNum() uint64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *QueryTxData) GetRowCount() uint64 {
	if m != nil {
		return m.RowCount
	}
	return 0
}

func (m *QueryTxData) GetRow() []*TxRow {
	if m != nil {
		return m.Row
	}
	return nil
}

type TxRow struct {
	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId" json:"transaction_id"`
	From          string `protobuf:"bytes,2,opt,name=from" json:"from"`
	To            string `protobuf:"bytes,3,opt,name=to" json:"to"`
	Price         uint64 `protobuf:"varint,4,opt,name=price" json:"price"`
	Type          string `protobuf:"bytes,5,opt,name=type" json:"type"`
	Date          string `protobuf:"bytes,6,opt,name=date" json:"date"`
	BlockId       uint64 `protobuf:"varint,7,opt,name=block_id,json=blockId" json:"block_id"`
}

func (m *TxRow) Reset()                    { *m = TxRow{} }
func (m *TxRow) String() string            { return proto.CompactTextString(m) }
func (*TxRow) ProtoMessage()               {}
func (*TxRow) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TxRow) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *TxRow) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *TxRow) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *TxRow) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *TxRow) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TxRow) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *TxRow) GetBlockId() uint64 {
	if m != nil {
		return m.BlockId
	}
	return 0
}

type TransferRequest struct {
	PostBody string `protobuf:"bytes,1,opt,name=postBody" json:"postBody"`
}

func (m *TransferRequest) Reset()                    { *m = TransferRequest{} }
func (m *TransferRequest) String() string            { return proto.CompactTextString(m) }
func (*TransferRequest) ProtoMessage()               {}
func (*TransferRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TransferRequest) GetPostBody() string {
	if m != nil {
		return m.PostBody
	}
	return ""
}

type TransferResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Data string `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *TransferResponse) Reset()                    { *m = TransferResponse{} }
func (m *TransferResponse) String() string            { return proto.CompactTextString(m) }
func (*TransferResponse) ProtoMessage()               {}
func (*TransferResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TransferResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TransferResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *TransferResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type QueryTransferRequest struct {
	Username  string `protobuf:"bytes,1,opt,name=username" json:"username"`
	SessionId string `protobuf:"bytes,2,opt,name=session_id,json=sessionId" json:"session_id"`
	Random    string `protobuf:"bytes,3,opt,name=random" json:"random"`
	Signature string `protobuf:"bytes,4,opt,name=signature" json:"signature"`
}

func (m *QueryTransferRequest) Reset()                    { *m = QueryTransferRequest{} }
func (m *QueryTransferRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryTransferRequest) ProtoMessage()               {}
func (*QueryTransferRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *QueryTransferRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *QueryTransferRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *QueryTransferRequest) GetRandom() string {
	if m != nil {
		return m.Random
	}
	return ""
}

func (m *QueryTransferRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type QueryTransferResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Data string `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *QueryTransferResponse) Reset()                    { *m = QueryTransferResponse{} }
func (m *QueryTransferResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryTransferResponse) ProtoMessage()               {}
func (*QueryTransferResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *QueryTransferResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *QueryTransferResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *QueryTransferResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type IsBuyAssetRequest struct {
	PageSize  uint32 `protobuf:"varint,1,opt,name=page_size,json=pageSize" json:"page_size"`
	PageNum   uint32 `protobuf:"varint,2,opt,name=page_num,json=pageNum" json:"page_num"`
	Username  string `protobuf:"bytes,3,opt,name=username" json:"username"`
	Random    string `protobuf:"bytes,4,opt,name=random" json:"random"`
	Signature string `protobuf:"bytes,5,opt,name=signature" json:"signature"`
	AssetId   string `protobuf:"bytes,6,opt,name=asset_id,json=assetId" json:"asset_id"`
}

func (m *IsBuyAssetRequest) Reset()                    { *m = IsBuyAssetRequest{} }
func (m *IsBuyAssetRequest) String() string            { return proto.CompactTextString(m) }
func (*IsBuyAssetRequest) ProtoMessage()               {}
func (*IsBuyAssetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *IsBuyAssetRequest) GetPageSize() uint32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *IsBuyAssetRequest) GetPageNum() uint32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *IsBuyAssetRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *IsBuyAssetRequest) GetRandom() string {
	if m != nil {
		return m.Random
	}
	return ""
}

func (m *IsBuyAssetRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *IsBuyAssetRequest) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

type IsBuyAssetResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code"`
	Data string `protobuf:"bytes,2,opt,name=data" json:"data"`
	Msg  string `protobuf:"bytes,3,opt,name=msg" json:"msg"`
}

func (m *IsBuyAssetResponse) Reset()                    { *m = IsBuyAssetResponse{} }
func (m *IsBuyAssetResponse) String() string            { return proto.CompactTextString(m) }
func (*IsBuyAssetResponse) ProtoMessage()               {}
func (*IsBuyAssetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *IsBuyAssetResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *IsBuyAssetResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *IsBuyAssetResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*PushRequest)(nil), "PushRequest")
	proto.RegisterType((*BuyAssetResponse)(nil), "BuyAssetResponse")
	proto.RegisterType((*QueryRequest)(nil), "QueryRequest")
	proto.RegisterType((*GetTxListResponse)(nil), "GetTxListResponse")
	proto.RegisterType((*QueryTxData)(nil), "QueryTxData")
	proto.RegisterType((*TxRow)(nil), "TxRow")
	proto.RegisterType((*TransferRequest)(nil), "TransferRequest")
	proto.RegisterType((*TransferResponse)(nil), "TransferResponse")
	proto.RegisterType((*QueryTransferRequest)(nil), "QueryTransferRequest")
	proto.RegisterType((*QueryTransferResponse)(nil), "QueryTransferResponse")
	proto.RegisterType((*IsBuyAssetRequest)(nil), "IsBuyAssetRequest")
	proto.RegisterType((*IsBuyAssetResponse)(nil), "IsBuyAssetResponse")
}

func init() { proto.RegisterFile("exchange.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 688 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x5e, 0xfa, 0x9f, 0xd3, 0x6e, 0x6c, 0x66, 0x40, 0x28, 0x4c, 0x2a, 0x91, 0x90, 0x7a, 0xb3,
	0x08, 0xc6, 0x05, 0xd7, 0xdb, 0x40, 0x53, 0xa5, 0x69, 0x62, 0xa1, 0x77, 0x08, 0x55, 0x5e, 0x72,
	0x96, 0x45, 0x24, 0x76, 0xb1, 0x1d, 0xd6, 0xee, 0x05, 0x78, 0x05, 0xae, 0x78, 0x01, 0x5e, 0x80,
	0x37, 0xe1, 0x75, 0x90, 0x9d, 0xa4, 0xcd, 0x3a, 0x0d, 0x4d, 0xda, 0x0d, 0x77, 0xe7, 0x3b, 0xf6,
	0xb1, 0xbf, 0xef, 0xf3, 0xd7, 0x14, 0x36, 0x70, 0x16, 0x5c, 0x50, 0x16, 0xa1, 0x37, 0x15, 0x5c,
	0x71, 0xf7, 0x67, 0x0d, 0xba, 0x1f, 0x32, 0x79, 0xe1, 0xe3, 0xd7, 0x0c, 0xa5, 0x22, 0x0e, 0xb4,
	0xbf, 0xa1, 0x90, 0x31, 0x67, 0x8e, 0x35, 0xb0, 0x86, 0xeb, 0x7e, 0x09, 0xc9, 0x0e, 0x40, 0x90,
	0x09, 0xc9, 0xc5, 0x84, 0x65, 0xa9, 0x53, 0x33, 0x8b, 0x76, 0xde, 0x39, 0xc9, 0x52, 0xf2, 0x02,
	0x7a, 0xc5, 0x72, 0x42, 0xcf, 0x30, 0x71, 0xea, 0x66, 0x43, 0x37, 0xef, 0x1d, 0xeb, 0x16, 0xe9,
	0x43, 0x27, 0x89, 0xcf, 0x51, 0xc5, 0x29, 0x3a, 0x8d, 0x81, 0x35, 0x6c, 0xf8, 0x0b, 0x4c, 0x1e,
	0x43, 0x4b, 0x22, 0x0b, 0x51, 0x38, 0xcd, 0x81, 0x35, 0xb4, 0xfd, 0x02, 0xe9, 0x99, 0x80, 0x33,
	0x25, 0x68, 0xa0, 0x9c, 0x96, 0x59, 0x59, 0x60, 0x3d, 0x93, 0xa2, 0xba, 0xe0, 0xa1, 0xd3, 0xce,
	0x67, 0x72, 0x44, 0xb6, 0xa1, 0x39, 0xa5, 0x82, 0xa6, 0x4e, 0xc7, 0xb4, 0x73, 0x40, 0x9e, 0x40,
	0x5b, 0xc6, 0xd1, 0x84, 0x26, 0x91, 0x63, 0x1b, 0x6e, 0x2d, 0x19, 0x47, 0xfb, 0x49, 0x44, 0x9e,
	0x83, 0x2d, 0xe3, 0x88, 0x51, 0x95, 0x09, 0x74, 0xc0, 0x8c, 0x2c, 0x1b, 0xee, 0x31, 0x6c, 0x1e,
	0x64, 0xf3, 0x7d, 0x29, 0x51, 0xf9, 0x28, 0xa7, 0x9c, 0x49, 0x24, 0x04, 0x1a, 0x01, 0x0f, 0xb1,
	0x70, 0xc8, 0xd4, 0xba, 0x17, 0x52, 0x45, 0x8d, 0x31, 0xb6, 0x6f, 0x6a, 0xb2, 0x09, 0xf5, 0x54,
	0x46, 0xc6, 0x0a, 0xdb, 0xd7, 0xa5, 0xfb, 0xc3, 0x82, 0xde, 0x69, 0x86, 0x62, 0x5e, 0xfa, 0xfd,
	0x0c, 0xec, 0x29, 0x8d, 0x70, 0x22, 0xe3, 0xab, 0xf2, 0xbc, 0x8e, 0x6e, 0x7c, 0x8c, 0xaf, 0x90,
	0x3c, 0x05, 0x53, 0x57, 0x0c, 0x6f, 0x6b, 0xac, 0xed, 0xee, 0x43, 0x27, 0x93, 0x28, 0x18, 0x4d,
	0xb1, 0x38, 0x7f, 0x81, 0xb5, 0x2f, 0x82, 0xb2, 0x90, 0xa7, 0xc6, 0x65, 0xdb, 0x2f, 0xd0, 0x75,
	0xa1, 0xcd, 0x55, 0xa1, 0x9f, 0x60, 0xeb, 0x08, 0xd5, 0x78, 0x76, 0x1c, 0xcb, 0x7f, 0x2b, 0x1d,
	0x54, 0x94, 0x76, 0xf7, 0x7a, 0x9e, 0xd1, 0x33, 0x9e, 0xbd, 0xa3, 0x8a, 0xde, 0xaa, 0xfb, 0x33,
	0x74, 0x2b, 0xdb, 0x74, 0xca, 0x0a, 0x21, 0xe6, 0xe4, 0xc6, 0x35, 0x5d, 0x82, 0x5f, 0x1e, 0xf2,
	0x8c, 0x29, 0x73, 0x41, 0xc3, 0x5f, 0x60, 0xe2, 0x40, 0x5d, 0xf0, 0x4b, 0xa7, 0x3e, 0xa8, 0x0f,
	0xbb, 0x7b, 0x2d, 0x6f, 0x3c, 0xf3, 0xf9, 0xa5, 0xaf, 0x5b, 0xee, 0x2f, 0x0b, 0x9a, 0x06, 0x92,
	0x97, 0xb0, 0xa1, 0x04, 0x65, 0x92, 0x06, 0x2a, 0xe6, 0x6c, 0x12, 0x87, 0xe6, 0x02, 0xdb, 0x5f,
	0xaf, 0x74, 0x47, 0xa1, 0xd6, 0x75, 0x2e, 0x78, 0x5a, 0xbe, 0x96, 0xae, 0xc9, 0x06, 0xd4, 0x14,
	0x2f, 0x48, 0xd7, 0x14, 0x37, 0x31, 0x12, 0x71, 0x50, 0x66, 0x35, 0x07, 0x7a, 0x52, 0xcd, 0xa7,
	0xa5, 0x7f, 0xa6, 0x2e, 0xde, 0x1e, 0x8b, 0x80, 0x9a, 0x5a, 0xbf, 0xdd, 0x59, 0xc2, 0x83, 0x2f,
	0x9a, 0x42, 0x3b, 0xd7, 0x68, 0xf0, 0x28, 0x74, 0x77, 0xe1, 0xc1, 0x58, 0xb3, 0x39, 0x47, 0x51,
	0xc6, 0xa0, 0x0f, 0x9d, 0x29, 0x97, 0xea, 0x80, 0x87, 0xf3, 0x82, 0xf0, 0x02, 0xeb, 0x04, 0x2e,
	0xb7, 0xdf, 0x3b, 0x81, 0xdf, 0x2d, 0xd8, 0xce, 0x9f, 0xe2, 0x26, 0x85, 0x45, 0xa2, 0xac, 0x95,
	0x44, 0xed, 0x00, 0x48, 0x94, 0xb2, 0x70, 0xb4, 0x56, 0x44, 0x27, 0xef, 0x8c, 0xc2, 0x4a, 0xe0,
	0xea, 0xb7, 0x07, 0xae, 0xb1, 0x1a, 0xb8, 0x53, 0x78, 0xb4, 0x42, 0xe4, 0xde, 0xe2, 0x7e, 0x5b,
	0xb0, 0x35, 0x92, 0xcb, 0xdf, 0xeb, 0x7f, 0xf6, 0x1b, 0xd3, 0x97, 0x51, 0xcd, 0x4c, 0xbb, 0x98,
	0x87, 0xa5, 0x6d, 0xf0, 0x28, 0x74, 0x4f, 0x80, 0x54, 0x99, 0xdf, 0xd7, 0x8a, 0xbd, 0x3f, 0x16,
	0x74, 0xde, 0x17, 0xdf, 0x7a, 0xb2, 0x0b, 0x9d, 0xf2, 0x68, 0xd2, 0xf3, 0x2a, 0xdf, 0xfb, 0xfe,
	0x96, 0xb7, 0x7a, 0xa7, 0xbb, 0x46, 0xde, 0x02, 0x2c, 0xb9, 0x10, 0xe2, 0xdd, 0xb0, 0xb4, 0xff,
	0xd0, 0xbb, 0x49, 0xd6, 0x5d, 0x23, 0xaf, 0xa0, 0x7b, 0x24, 0x28, 0x53, 0x87, 0x02, 0xc3, 0xf8,
	0x4e, 0x57, 0xbd, 0x86, 0xde, 0x21, 0x65, 0x01, 0x26, 0xc1, 0x5d, 0x47, 0xce, 0x5a, 0xe6, 0x9f,
	0xeb, 0xcd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xfb, 0x5b, 0xed, 0xcb, 0x06, 0x00, 0x00,
}
