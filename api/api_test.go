package api

import (
	"bytes"
	"testing"
	"time"

	"github.com/Crown-Labs/xoracle-go-sdk/common"
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
	for _, v := range tokenIndexPrices {
		if v.Price.Sign() == 0 {
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
		for _, v := range tokenAddressPrices {
			// Check tip.TokenAddress is in addresses
			isFound := false
			for _, address := range chain.TokenAddress {
				if bytes.Equal(v.TokenAddress.Bytes(), address.Bytes()) {
					isFound = true
					break
				}
			}

			if !isFound {
				t.Errorf("GetTokenAddressPrice returned a TokenAddressPrice with an invalid address")
			}

			if v.Price.Sign() == 0 {
				t.Errorf("GetTokenIndexPrice returned a TokenIndexPrice with a zero price")
			}
		}
	}
}

func TestGetTokenIndexInfo(t *testing.T) {
	api := NewApi(nil)

	// Test that the function returns a non-empty slice of TokenIndexInfo structs
	tokenIndexInfo, err := api.GetTokenIndexInfo()
	if err != nil {
		t.Errorf("GetTokenIndexInfo returned an error: %v", err)
	}
	if len(tokenIndexInfo) == 0 {
		t.Errorf("GetTokenIndexInfo returned an empty slice")
	}

	// Test that each TokenIndexInfo struct has a non-zero token index, name and symbol
	for k, v := range tokenIndexInfo {
		if !common.AllowTokenIndex[k] {
			t.Errorf("GetTokenIndexInfo returned a TokenIndexInfo with an invalid token index")
		}
		if v.TokenName == "" {
			t.Errorf("GetTokenIndexInfo returned a TokenIndexInfo with an empty token name")
		}
		if v.TokenSymbol == "" {
			t.Errorf("GetTokenIndexInfo returned a TokenIndexInfo with an empty token symbol")
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
	for _, v := range nodeInfo {
		if v.NodeAddress.Bytes() == nil {
			t.Errorf("GetNodeInfo returned a NodeInfo with an empty node address")
		}
		if v.NodeName == "" {
			t.Errorf("GetNodeInfo returned a NodeInfo with an empty node name")
		}
	}
}

func GetTokenIndexPriceByTimestamp(t *testing.T) {
	api := NewApi(nil)
	// Get timestamp
	timestamp := time.Now().Unix()

	// Test that the function returns a non-empty slice of TokenIndexPrice structs
	tokenIndexPrices, err := api.GetTokenIndexPriceByTimestamp(timestamp)
	if err != nil {
		t.Errorf("GetTokenIndexPrice returned an error: %v", err)
	}
	if len(tokenIndexPrices) == 0 {
		t.Errorf("GetTokenIndexPrice returned an empty slice")
	}

	// Test that each TokenIndexPrice struct has a non-zero token index and price
	for _, v := range tokenIndexPrices {
		if v.Price.Sign() == 0 {
			t.Errorf("GetTokenIndexPrice returned a TokenIndexPrice with a zero price")
		}
	}
}

func TestGetTokenAddressPriceByTimestamp(t *testing.T) {
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
		// Get timestamp
		timestamp := time.Now().Unix()
		// Test that the function returns a non-empty slice of TokenAddressPrice structs
		tokenAddressPrices, err := api.GetTokenAddressPriceByTimestamp(networkId, timestamp)
		if err != nil {
			t.Errorf("GetTokenAddressPrice returned an error: %v", err)
		}
		if len(tokenAddressPrices) == 0 {
			t.Errorf("GetTokenAddressPrice returned an empty result")
		}

		// Test that each TokenAddressPrice struct has a non-zero token address and price
		for _, v := range tokenAddressPrices {
			// Check tip.TokenAddress is in addresses
			isFound := false
			for _, address := range chain.TokenAddress {
				if bytes.Equal(v.TokenAddress.Bytes(), address.Bytes()) {
					isFound = true
					break
				}
			}

			if !isFound {
				t.Errorf("GetTokenAddressPrice returned a TokenAddressPrice with an invalid address")
			}

			if v.Price.Sign() == 0 {
				t.Errorf("GetTokenIndexPrice returned a TokenIndexPrice with a zero price")
			}
		}
	}
}
