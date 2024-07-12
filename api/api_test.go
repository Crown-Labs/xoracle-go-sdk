package api

import (
	"math/big"
	"testing"

	"github.com/Crown-Labs/xoracle-go-sdk/common"
	"github.com/stretchr/testify/assert"
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

func TestGetTokenIndexPriceByTimestamp(t *testing.T) {
	api := NewApi([]int{
		common.TOKEN_INDEX.BTC,
		common.TOKEN_INDEX.ETH,
		common.TOKEN_INDEX.USDT,
		common.TOKEN_INDEX.USDC,
		common.TOKEN_INDEX.SOL,
		common.TOKEN_INDEX.OP,
		common.TOKEN_INDEX.ARB,
	})
	// Get timestamp
	var timestamp int64 = 1720569600 // 2024-07-10 00:00:00 UTC
	assert := assert.New(t)
	kvMap := map[int]*big.Int{
		0:  big.NewInt(5802596226469),
		1:  big.NewInt(306450811266),
		3:  big.NewInt(99964254),
		5:  big.NewInt(99987895),
		22: big.NewInt(14149421526),
		28: big.NewInt(161950000),
		29: big.NewInt(70967075),
	}

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
		for k2, v2 := range kvMap {
			if v.TokenIndex == k2 {
				assert.Equal(v.Price, v2)
			}
		}
	}
}
