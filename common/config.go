package common

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

var MainnetAllowTokenIndex = map[int]bool{
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

var TestnetAllowTokenIndex = map[int]bool{
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
var MainnetConfig = ConfigType{
	XOracleAPI:                 "https://api.xoracle.io/",
	EndpointAPIPrice:           "prices/xoracle",
	EndpointAPINodePrice:       "", // prices/xoracle/node
	EndpointAPITokenIndexPrice: "", // prices/xoracle/:index
	EndpointAPIPricefeed:       "", // prices/pricefeed/:timestamp
	EndpointAPITokenIndexInfo:  "prices/tokenIndexInfo",
	EndpointAPINodeInfo:        "prices/nodeInfo",
}

var TestnetConfig = ConfigType{
	XOracleAPI:                 "https://api-testnet.xoracle.io/",
	EndpointAPIPrice:           "prices/xoracle",
	EndpointAPINodePrice:       "", // prices/xoracle/node
	EndpointAPITokenIndexPrice: "", // prices/xoracle/:index
	EndpointAPIPricefeed:       "", // prices/pricefeed/:timestamp
	EndpointAPITokenIndexInfo:  "prices/tokenIndexInfo",
	EndpointAPINodeInfo:        "prices/nodeInfo",
}
