package common

import (
	geth_common "github.com/ethereum/go-ethereum/common"
)

// Constants
var NETWORK_ID = NetworkId{
	DEVELOP_TESTNET:  1112,
	ARBITRUM_TESTNET: 421614,
}

var TOKEN_DECIMALS = 8

var TOKEN_INDEX = TokenIndex{
	BTC:   0,
	ETH:   1,
	BNB:   2,
	USDT:  3,
	BUSD:  4,
	USDC:  5,
	DAI:   6,
	XRP:   10,
	DOGE:  11,
	TRX:   12,
	ADA:   20,
	MATIC: 21,
	SOL:   22,
	DOT:   23,
	AVAX:  24,
	FTM:   25,
	NEAR:  26,
	ATOM:  27,
	OP:    28,
	ARB:   29,
}

var AllowTokenIndex = map[int]bool{
	TOKEN_INDEX.BTC:   true,
	TOKEN_INDEX.ETH:   true,
	TOKEN_INDEX.BNB:   true,
	TOKEN_INDEX.USDT:  true,
	TOKEN_INDEX.BUSD:  true,
	TOKEN_INDEX.USDC:  true,
	TOKEN_INDEX.DAI:   true,
	TOKEN_INDEX.XRP:   true,
	TOKEN_INDEX.DOGE:  true,
	TOKEN_INDEX.TRX:   true,
	TOKEN_INDEX.ADA:   true,
	TOKEN_INDEX.MATIC: true,
	TOKEN_INDEX.SOL:   true,
	TOKEN_INDEX.DOT:   true,
	TOKEN_INDEX.AVAX:  true,
	TOKEN_INDEX.FTM:   true,
	TOKEN_INDEX.NEAR:  true,
	TOKEN_INDEX.ATOM:  true,
	TOKEN_INDEX.OP:    true,
	TOKEN_INDEX.ARB:   true,
}

// Config
var Config = ConfigType{
	XOracleAPI:                 AppendUrlWithSlash("https://api.xoracle.io"),
	EndpointAPIPrice:           "prices/xoracle",
	EndpointAPINodePrice:       "", // prices/xoracle/node
	EndpointAPITokenIndexPrice: "", // prices/xoracle/:index
	EndpointAPIPricefeed:       "", // prices/pricefeed/:timestamp
	EndpointAPITokenIndexInfo:  "prices/tokenIndexInfo",
	EndpointAPINodeInfo:        "prices/nodeInfo",

	Chains: ChainsType{
		NETWORK_ID.DEVELOP_TESTNET: {
			NetworkId: NETWORK_ID.DEVELOP_TESTNET,
			TokenAddress: map[int]geth_common.Address{
				TOKEN_INDEX.BTC:   geth_common.HexToAddress("0x9E4317b6a9e990D38A7D6850e00A794B49F77248"),
				TOKEN_INDEX.ETH:   geth_common.HexToAddress("0x078c04b8cfC949101905fdd5912D31Aad0a244Cb"),
				TOKEN_INDEX.BNB:   geth_common.HexToAddress("0x800f889DA4bba83E966536113F4b99461C1d35D2"),
				TOKEN_INDEX.USDT:  geth_common.HexToAddress("0xa07C8259fBFa95FB7DaDc536030de6b1e9fA3e49"),
				TOKEN_INDEX.USDC:  geth_common.HexToAddress("0x1eBfbAf2646Eb623334Ea2B5f45D18cDAc3dbfFE"),
				TOKEN_INDEX.MATIC: geth_common.HexToAddress("0xB5F3587C5a9D04233254dDf7f221AFeA444d7fA6"),
				TOKEN_INDEX.OP:    geth_common.HexToAddress("0x76806Bb3b09F5fE634a75B6990EfEde67133B7Cf"),
				TOKEN_INDEX.ARB:   geth_common.HexToAddress("0xEA041B9a61d59E9A4D521Ee72F70Be29F0D43e9b"),
			},
		},
		NETWORK_ID.ARBITRUM_TESTNET: {
			NetworkId: NETWORK_ID.ARBITRUM_TESTNET,
			TokenAddress: map[int]geth_common.Address{
				TOKEN_INDEX.BTC:  geth_common.HexToAddress("0x3b0C46252B44bd40c54f1557970171b0D70d2e53"),
				TOKEN_INDEX.ETH:  geth_common.HexToAddress("0x2C1b868d6596a18e32E61B901E4060C872647b6C"),
				TOKEN_INDEX.USDT: geth_common.HexToAddress("0xbceC86a186274C4215030b1a5903FAC944b540dd"),
				TOKEN_INDEX.USDC: geth_common.HexToAddress("0x539912979f90fb0eb170c61fCE865d1714659Cf8"),
				TOKEN_INDEX.SOL:  geth_common.HexToAddress("0xB80808027a0Ef560528bF8f1C7788a942bBc0CE7"),
				TOKEN_INDEX.OP:   geth_common.HexToAddress("0x1e0BBfD45acda3DafD0d9CC6AdC2B1F8Ba5C1b6A"),
				TOKEN_INDEX.ARB:  geth_common.HexToAddress("0x45950C3dB9f713b10cfd1208A2f0007175697FED"),
			},
		},
	},
}
