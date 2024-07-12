package sdk

import (
	"fmt"
	"time"

	"context"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	bbnclient "github.com/babylonchain/babylon/client/client"
	bbncfg "github.com/babylonchain/babylon/client/config"
	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"go.uber.org/zap"

	"github.com/babylonchain/babylon-finality-gadget/btcclient"
	sdkconfig "github.com/babylonchain/babylon-finality-gadget/sdk/config"
	"github.com/babylonchain/babylon-finality-gadget/testutils"
)

type BtcClient interface {
	GetBlockHeightByTimestamp(targetTimestamp uint64) (uint64, error)
}

// SdkClient is a client that can only perform queries to a Babylon node
// It only requires the client config to have `rpcAddr`, but not other fields
// such as keyring, chain ID, etc..
type SdkClient struct {
	config    *sdkconfig.Config
	bbnClient *bbnclient.Client
	BtcClient BtcClient
}

// NewClient creates a new BabylonFinalityGadgetClient according to the given config
func NewClient(config *sdkconfig.Config) (*SdkClient, error) {
	rpcAddr, err := config.GetRpcAddr()
	if err != nil {
		return nil, err
	}

	bbnConfig := bbncfg.DefaultBabylonConfig()
	bbnConfig.RPCAddr = rpcAddr

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	// Note: We can just ignore the below info which is printed by bbnclient.New
	// service injective.evm.v1beta1.Msg does not have cosmos.msg.v1.service proto annotation
	bbnClient, err := bbnclient.New(
		&bbnConfig,
		logger,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Babylon client: %w", err)
	}

	var btcClient BtcClient
	// Create BTC client
	switch config.ChainType {
	case sdkconfig.BabylonLocalnet:
		btcClient, err = testutils.NewMockBTCClient(config.BTCConfig, logger)
	default:
		btcClient, err = btcclient.NewBTCClient(config.BTCConfig, logger)
	}
	if err != nil {
		return nil, err
	}
	return &SdkClient{
		bbnClient: bbnClient,
		config:    config,
		BtcClient: btcClient,
	}, nil
}

// querySmartContractState queries the smart contract state given the contract address and query data
func (babylonClient *SdkClient) querySmartContractState(contractAddress string, queryData []byte) (*wasmtypes.QuerySmartContractStateResponse, error) {
	// hardcode the timeout to 20 seconds. We can expose it to the params once needed
	timeout := 20 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	sdkClientCtx := sdkclient.Context{Client: babylonClient.bbnClient.RPCClient}
	wasmQueryClient := wasmtypes.NewQueryClient(sdkClientCtx)

	req := &wasmtypes.QuerySmartContractStateRequest{
		Address:   contractAddress,
		QueryData: queryData,
	}
	return wasmQueryClient.SmartContractState(ctx, req)
}
