package interchaintest

import (
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/types"
	balancertypes "github.com/notional-labs/fee-abstraction/tests/interchaintest/osmosistypes/gamm/balancer"
	gammtypes "github.com/notional-labs/fee-abstraction/tests/interchaintest/osmosistypes/gamm/types"
	feeabstype "github.com/notional-labs/fee-abstraction/v2/x/feeabs/types"
	"github.com/strangelove-ventures/interchaintest/v4/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v4/chain/cosmos/wasm"
	"github.com/strangelove-ventures/interchaintest/v4/ibc"
)

type QueryFeeabsModuleBalacesResponse struct {
	Balances types.Coins
	Address  string
}

type QueryHostChainConfigRespone struct {
	HostChainConfig cosmos.HostChainFeeAbsConfig `protobuf:"bytes,1,opt,name=host_chain_config,json=hostChainConfig,proto3" json:"host_chain_config" yaml:"host_chain_config"`
}

const (
	votingPeriod     = "10s"
	maxDepositPeriod = "10s"
)

var (
	FeeabsMainRepo = "ghcr.io/notional-labs/fee-abstraction"

	feeabsImage = ibc.DockerImage{
		Repository: "ghcr.io/notional-labs/fee-abstraction",
		Version:    "2.0.4",
		UidGid:     "1025:1025",
	}

	feeabsConfig = ibc.ChainConfig{
		Type:                "cosmos",
		Name:                "feeabs",
		ChainID:             "feeabs-2",
		Images:              []ibc.DockerImage{feeabsImage},
		Bin:                 "feeappd",
		Bech32Prefix:        "feeabs",
		Denom:               "stake",
		CoinType:            "118",
		GasPrices:           "0.0stake",
		GasAdjustment:       1.1,
		TrustingPeriod:      "112h",
		NoHostMount:         false,
		SkipGenTx:           false,
		PreGenesis:          nil,
		ModifyGenesis:       cosmos.ModifyGenesisProposalTime(votingPeriod, maxDepositPeriod),
		ConfigFileOverrides: nil,
		EncodingConfig:      feeabsEncoding(),
	}

	pathFeeabsGaia      = "feeabs-gaia"
	pathFeeabsOsmosis   = "feeabs-osmosis"
	pathOsmosisGaia     = "osmosis-gaia"
	genesisWalletAmount = int64(10_000_000)
)

// feeabsEncoding registers the feeabs specific module codecs so that the associated types and msgs
// will be supported when writing to the blocksdb sqlite database.
func feeabsEncoding() *simappparams.EncodingConfig {
	cfg := wasm.WasmEncoding()

	// register custom types
	feeabstype.RegisterInterfaces(cfg.InterfaceRegistry)

	return cfg
}

func osmosisEncoding() *simappparams.EncodingConfig {
	cfg := wasm.WasmEncoding()

	gammtypes.RegisterInterfaces(cfg.InterfaceRegistry)
	balancertypes.RegisterInterfaces(cfg.InterfaceRegistry)

	return cfg
}