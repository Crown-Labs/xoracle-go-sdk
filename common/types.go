package common

import (
	"math/big"

	geth_common "github.com/ethereum/go-ethereum/common"
)

// Constants
var (
	BigInt0        = big.NewInt(0)
	BigInt1        = big.NewInt(1)
	BigInt10       = big.NewInt(10)
	BigFloat0      = big.NewFloat(0)
	BigFloatBase18 = big.NewFloat(1e18)
	BigFloatBase9  = big.NewFloat(1e9)
)

type TokenIndex struct {
	BTC   int
	ETH   int
	BNB   int
	USDT  int
	BUSD  int
	USDC  int
	DAI   int
	XRP   int
	DOGE  int
	TRX   int
	ADA   int
	MATIC int
	SOL   int
	DOT   int
	AVAX  int
	FTM   int
	NEAR  int
	ATOM  int
	OP    int
	ARB   int
}

type ConfigType struct {
	XOracleAPI                 string
	EndpointAPIPrice           string
	EndpointAPINodePrice       string
	EndpointAPITokenIndexPrice string
	EndpointAPIPricefeed       string
	EndpointAPITokenIndexInfo  string
	EndpointAPINodeInfo        string
}

type TokenIndexPrice struct {
	TokenIndex int
	Price      *big.Int
}

type TokenIndexInfo struct {
	TokenIndex  int
	TokenName   string
	TokenSymbol string
}

type NodeInfo struct {
	NodeAddress geth_common.Address
	NodeName    string
}

type NetworkType int

type Network struct {
	MAINNET NetworkType
	TESTNET NetworkType
}
