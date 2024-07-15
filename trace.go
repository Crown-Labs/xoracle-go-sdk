package main

import (
	"fmt"

	"github.com/Crown-Labs/xoracle-go-sdk/api"
	"github.com/Crown-Labs/xoracle-go-sdk/common"
)

func main() {
	api := api.NewApi("testnet", []int{
		common.TOKEN_INDEX.BTC,
		common.TOKEN_INDEX.ETH,
		common.TOKEN_INDEX.BUSD,
		common.TOKEN_INDEX.USDC,
		common.TOKEN_INDEX.SOL,
		common.TOKEN_INDEX.OP,
		common.TOKEN_INDEX.ARB,
	})

	// Trace that the function returns a non-empty slice of TokenIndexPrice structs
	fmt.Printf("\n🔍 xOracle API: %s\n", common.TestnetConfig.XOracleAPI+common.TestnetConfig.EndpointAPIPrice)
	tokenIndexPrices, err := api.GetTokenIndexPrice()
	if err != nil {
		fmt.Println(err)
	}

	for _, price := range tokenIndexPrices {
		fmt.Printf("TokenIndex: %d, Price: %s, Price 18 Decimals: %s\n",
			price.TokenIndex,
			price.Price.String(),
			common.ConvertPriceDecimals(price.Price, 18).String(),
		)
	}

	// Trace that the function returns a non-empty slice of TokenIndexInfo structs
	fmt.Printf("\n🔍 xOracle API: %s\n", common.TestnetConfig.XOracleAPI+common.TestnetConfig.EndpointAPITokenIndexInfo)
	tokenIndexInfos, err := api.GetTokenIndexInfo()
	if err != nil {
		fmt.Println(err)
	}
	for tokenIndex, tokenIndexInfo := range tokenIndexInfos {
		fmt.Printf("TokenIndex: %d, TokenName: %s, TokenSymbol: %s\n", tokenIndex, tokenIndexInfo.TokenName, tokenIndexInfo.TokenSymbol)
	}

	// Trace that the function returns a non-empty slice of NodeInfo structs
	fmt.Printf("\n🔍 xOracle API: %s\n", common.TestnetConfig.XOracleAPI+common.TestnetConfig.EndpointAPINodeInfo)
	nodeInfos, err := api.GetNodeInfo()
	if err != nil {
		fmt.Println(err)
	}
	for _, nodeInfo := range nodeInfos {
		fmt.Printf("NodeAddress: %s, NodeName: %s\n", nodeInfo.NodeAddress, nodeInfo.NodeName)
	}
}
