// Code generated by MockGen. DO NOT EDIT.
// Source: sdk/client/expected_clients.go
//
// Generated by this command:
//
//	mockgen -source=sdk/client/expected_clients.go -package mocks -destination /Users/lyy/babylon/babylon-finality-gadget/testutil/mocks/expected_clients_mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	cwclient "github.com/babylonchain/babylon-finality-gadget/sdk/cwclient"
	chainhash "github.com/btcsuite/btcd/chaincfg/chainhash"
	wire "github.com/btcsuite/btcd/wire"
	gomock "go.uber.org/mock/gomock"
)

// MockIBabylonClient is a mock of IBabylonClient interface.
type MockIBabylonClient struct {
	ctrl     *gomock.Controller
	recorder *MockIBabylonClientMockRecorder
}

// MockIBabylonClientMockRecorder is the mock recorder for MockIBabylonClient.
type MockIBabylonClientMockRecorder struct {
	mock *MockIBabylonClient
}

// NewMockIBabylonClient creates a new mock instance.
func NewMockIBabylonClient(ctrl *gomock.Controller) *MockIBabylonClient {
	mock := &MockIBabylonClient{ctrl: ctrl}
	mock.recorder = &MockIBabylonClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBabylonClient) EXPECT() *MockIBabylonClientMockRecorder {
	return m.recorder
}

// QueryAllFpBtcPubKeys mocks base method.
func (m *MockIBabylonClient) QueryAllFpBtcPubKeys(consumerId string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryAllFpBtcPubKeys", consumerId)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryAllFpBtcPubKeys indicates an expected call of QueryAllFpBtcPubKeys.
func (mr *MockIBabylonClientMockRecorder) QueryAllFpBtcPubKeys(consumerId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryAllFpBtcPubKeys", reflect.TypeOf((*MockIBabylonClient)(nil).QueryAllFpBtcPubKeys), consumerId)
}

// QueryEarliestActiveDelBtcHeight mocks base method.
func (m *MockIBabylonClient) QueryEarliestActiveDelBtcHeight(fpPubkeyHexList []string) (*uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryEarliestActiveDelBtcHeight", fpPubkeyHexList)
	ret0, _ := ret[0].(*uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryEarliestActiveDelBtcHeight indicates an expected call of QueryEarliestActiveDelBtcHeight.
func (mr *MockIBabylonClientMockRecorder) QueryEarliestActiveDelBtcHeight(fpPubkeyHexList any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryEarliestActiveDelBtcHeight", reflect.TypeOf((*MockIBabylonClient)(nil).QueryEarliestActiveDelBtcHeight), fpPubkeyHexList)
}

// QueryFpPower mocks base method.
func (m *MockIBabylonClient) QueryFpPower(fpPubkeyHex string, btcHeight uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryFpPower", fpPubkeyHex, btcHeight)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryFpPower indicates an expected call of QueryFpPower.
func (mr *MockIBabylonClientMockRecorder) QueryFpPower(fpPubkeyHex, btcHeight any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryFpPower", reflect.TypeOf((*MockIBabylonClient)(nil).QueryFpPower), fpPubkeyHex, btcHeight)
}

// QueryMultiFpPower mocks base method.
func (m *MockIBabylonClient) QueryMultiFpPower(fpPubkeyHexList []string, btcHeight uint64) (map[string]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryMultiFpPower", fpPubkeyHexList, btcHeight)
	ret0, _ := ret[0].(map[string]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryMultiFpPower indicates an expected call of QueryMultiFpPower.
func (mr *MockIBabylonClientMockRecorder) QueryMultiFpPower(fpPubkeyHexList, btcHeight any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryMultiFpPower", reflect.TypeOf((*MockIBabylonClient)(nil).QueryMultiFpPower), fpPubkeyHexList, btcHeight)
}

// MockIBitcoinClient is a mock of IBitcoinClient interface.
type MockIBitcoinClient struct {
	ctrl     *gomock.Controller
	recorder *MockIBitcoinClientMockRecorder
}

// MockIBitcoinClientMockRecorder is the mock recorder for MockIBitcoinClient.
type MockIBitcoinClientMockRecorder struct {
	mock *MockIBitcoinClient
}

// NewMockIBitcoinClient creates a new mock instance.
func NewMockIBitcoinClient(ctrl *gomock.Controller) *MockIBitcoinClient {
	mock := &MockIBitcoinClient{ctrl: ctrl}
	mock.recorder = &MockIBitcoinClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBitcoinClient) EXPECT() *MockIBitcoinClientMockRecorder {
	return m.recorder
}

// GetBlockCount mocks base method.
func (m *MockIBitcoinClient) GetBlockCount() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockCount")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockCount indicates an expected call of GetBlockCount.
func (mr *MockIBitcoinClientMockRecorder) GetBlockCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockCount", reflect.TypeOf((*MockIBitcoinClient)(nil).GetBlockCount))
}

// GetBlockHashByHeight mocks base method.
func (m *MockIBitcoinClient) GetBlockHashByHeight(height uint64) (*chainhash.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockHashByHeight", height)
	ret0, _ := ret[0].(*chainhash.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockHashByHeight indicates an expected call of GetBlockHashByHeight.
func (mr *MockIBitcoinClientMockRecorder) GetBlockHashByHeight(height any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockHashByHeight", reflect.TypeOf((*MockIBitcoinClient)(nil).GetBlockHashByHeight), height)
}

// GetBlockHeaderByHash mocks base method.
func (m *MockIBitcoinClient) GetBlockHeaderByHash(blockHash *chainhash.Hash) (*wire.BlockHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockHeaderByHash", blockHash)
	ret0, _ := ret[0].(*wire.BlockHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockHeaderByHash indicates an expected call of GetBlockHeaderByHash.
func (mr *MockIBitcoinClientMockRecorder) GetBlockHeaderByHash(blockHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockHeaderByHash", reflect.TypeOf((*MockIBitcoinClient)(nil).GetBlockHeaderByHash), blockHash)
}

// GetBlockHeightByTimestamp mocks base method.
func (m *MockIBitcoinClient) GetBlockHeightByTimestamp(targetTimestamp uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockHeightByTimestamp", targetTimestamp)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockHeightByTimestamp indicates an expected call of GetBlockHeightByTimestamp.
func (mr *MockIBitcoinClientMockRecorder) GetBlockHeightByTimestamp(targetTimestamp any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockHeightByTimestamp", reflect.TypeOf((*MockIBitcoinClient)(nil).GetBlockHeightByTimestamp), targetTimestamp)
}

// MockICosmWasmClient is a mock of ICosmWasmClient interface.
type MockICosmWasmClient struct {
	ctrl     *gomock.Controller
	recorder *MockICosmWasmClientMockRecorder
}

// MockICosmWasmClientMockRecorder is the mock recorder for MockICosmWasmClient.
type MockICosmWasmClientMockRecorder struct {
	mock *MockICosmWasmClient
}

// NewMockICosmWasmClient creates a new mock instance.
func NewMockICosmWasmClient(ctrl *gomock.Controller) *MockICosmWasmClient {
	mock := &MockICosmWasmClient{ctrl: ctrl}
	mock.recorder = &MockICosmWasmClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICosmWasmClient) EXPECT() *MockICosmWasmClientMockRecorder {
	return m.recorder
}

// QueryConsumerId mocks base method.
func (m *MockICosmWasmClient) QueryConsumerId() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryConsumerId")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryConsumerId indicates an expected call of QueryConsumerId.
func (mr *MockICosmWasmClientMockRecorder) QueryConsumerId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryConsumerId", reflect.TypeOf((*MockICosmWasmClient)(nil).QueryConsumerId))
}

// QueryIsEnabled mocks base method.
func (m *MockICosmWasmClient) QueryIsEnabled() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryIsEnabled")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryIsEnabled indicates an expected call of QueryIsEnabled.
func (mr *MockICosmWasmClientMockRecorder) QueryIsEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryIsEnabled", reflect.TypeOf((*MockICosmWasmClient)(nil).QueryIsEnabled))
}

// QueryListOfVotedFinalityProviders mocks base method.
func (m *MockICosmWasmClient) QueryListOfVotedFinalityProviders(queryParams *cwclient.L2Block) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryListOfVotedFinalityProviders", queryParams)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryListOfVotedFinalityProviders indicates an expected call of QueryListOfVotedFinalityProviders.
func (mr *MockICosmWasmClientMockRecorder) QueryListOfVotedFinalityProviders(queryParams any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryListOfVotedFinalityProviders", reflect.TypeOf((*MockICosmWasmClient)(nil).QueryListOfVotedFinalityProviders), queryParams)
}
