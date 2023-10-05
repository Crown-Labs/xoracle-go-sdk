package api

import (
	"testing"
	"xoracle-go-sdk/common"
)

func TestGetTokenIndexPrice(t *testing.T) {
	api := NewApi(nil)

	// Test that the function returns a non-empty slice of TokenIndexPrice structs
	tokenIndexPrices, err := api.GetTokenIndexPrice()
	if err != nil {
		t.Errorf("GetTokenIndexPrice returned an error: %v", err)
	}
	if len(tokenIndexPrices) == 0 {
		t.Errorf("GetTokenIndexPrice returned an empty slice")
	}

	// Test that each TokenIndexPrice struct has a non-zero token index and price
	for _, tip := range tokenIndexPrices {
		if tip.Price.Sign() == 0 {
			t.Errorf("GetTokenIndexPrice returned a TokenIndexPrice with a zero price")
		}
	}
}

func TestGetTokenAddressPrice(t *testing.T) {
	api := NewApi([]int{
		common.TOKEN_INDEX.BTC,
		common.TOKEN_INDEX.ETH,
		common.TOKEN_INDEX.BNB,
		common.TOKEN_INDEX.USDT,
		common.TOKEN_INDEX.BUSD,
		common.TOKEN_INDEX.USDC,
		common.TOKEN_INDEX.MATIC,
		common.TOKEN_INDEX.OP,
		common.TOKEN_INDEX.ARB,
	})

	for networkId, chain := range common.Config.Chains {
		// Test that the function returns a non-empty slice of TokenAddressPrice structs
		tokenAddressPrices, err := api.GetTokenAddressPrice(networkId)
		if err != nil {
			t.Errorf("GetTokenAddressPrice returned an error: %v", err)
		}
		if len(tokenAddressPrices) == 0 {
			t.Errorf("GetTokenAddressPrice returned an empty result")
		}

		// Test that each TokenAddressPrice struct has a non-zero token address and price
		for _, tip := range tokenAddressPrices {
			// Check tip.TokenAddress is in addresses
			isFound := false
			for _, address := range chain.TokenAddress {
				if tip.TokenAddress == address {
					isFound = true
					break
				}
			}

			if !isFound {
				t.Errorf("GetTokenAddressPrice returned a TokenAddressPrice with an invalid address")
			}

			if tip.Price.Sign() == 0 {
				t.Errorf("GetTokenIndexPrice returned a TokenIndexPrice with a zero price")
			}
		}
	}
}

func TestGetNodeInfo(t *testing.T) {
	api := NewApi(nil)

	// Test that the function returns a non-empty slice of NodeInfo structs
	nodeInfo, err := api.GetNodeInfo()
	if err != nil {
		t.Errorf("GetNodeInfo returned an error: %v", err)
	}
	if len(nodeInfo) == 0 {
		t.Errorf("GetNodeInfo returned an empty slice")
	}

	// Test that each NodeInfo struct has a non-zero node address and name
	for _, ni := range nodeInfo {
		if ni.NodeAddress == "" {
			t.Errorf("GetNodeInfo returned a NodeInfo with an empty node address")
		}
		if ni.NodeName == "" {
			t.Errorf("GetNodeInfo returned a NodeInfo with an empty node name")
		}
	}
}
