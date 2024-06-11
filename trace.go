package main

import (
	"fmt"

	"github.com/Crown-Labs/xoracle-go-sdk/api"
	"github.com/Crown-Labs/xoracle-go-sdk/common"
)

func main() {
	api := api.NewApi([]int{
		common.TOKEN_INDEX.BTC,
		common.TOKEN_INDEX.ETH,
		common.TOKEN_INDEX.BUSD,
		common.TOKEN_INDEX.USDC,
		common.TOKEN_INDEX.SOL,
		common.TOKEN_INDEX.OP,
		common.TOKEN_INDEX.ARB,
	})
	networkId := common.NETWORK_ID.ARBITRUM_TESTNET

	// Trace that the function returns a non-empty slice of TokenIndexPrice structs
	fmt.Printf("\n🔍 xOracle API: %s\n", common.Config.XOracleAPI+common.Config.EndpointAPIPrice)
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

	tokenAddressPrices, err := api.GetTokenAddressPrice(networkId)
	if err != nil {
		fmt.Println(err)
	}

	for _, price := range tokenAddressPrices {
		fmt.Printf("TokenAddress: %s, Price: %s, Price 18 Decimals: %s\n",
			price.TokenAddress,
			price.Price.String(),
			common.ConvertPriceDecimals(price.Price, 18).String(),
		)
	}

	// Trace that the function returns a non-empty slice of TokenIndexInfo structs
	fmt.Printf("\n🔍 xOracle API: %s\n", common.Config.XOracleAPI+common.Config.EndpointAPITokenIndexInfo)
	tokenIndexInfos, err := api.GetTokenIndexInfo()
	if err != nil {
		fmt.Println(err)
	}
	for tokenIndex, tokenIndexInfo := range tokenIndexInfos {
		fmt.Printf("TokenIndex: %d, TokenName: %s, TokenSymbol: %s\n", tokenIndex, tokenIndexInfo.TokenName, tokenIndexInfo.TokenSymbol)
	}

	// Trace that the function returns a non-empty slice of NodeInfo structs
	fmt.Printf("\n🔍 xOracle API: %s\n", common.Config.XOracleAPI+common.Config.EndpointAPINodeInfo)
	nodeInfos, err := api.GetNodeInfo()
	if err != nil {
		fmt.Println(err)
	}
	for _, nodeInfo := range nodeInfos {
		fmt.Printf("NodeAddress: %s, NodeName: %s\n", nodeInfo.NodeAddress, nodeInfo.NodeName)
	}
}
