package api

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"sort"

	"github.com/Crown-Labs/xoracle-go-sdk/common"
	geth_common "github.com/ethereum/go-ethereum/common"
)

type Api struct {
	acceptTokenIndex map[int]bool
	config           common.ConfigType
}

// NewApi returns a new Api instance.
// It accepts a slice of token indexes, which are used to filter the results of the GetTokenIndexPrice and GetTokenAddressPrice functions.
func NewApi(network string, tokenIndexes []int) *Api {
	var acceptTokenIndex = make(map[int]bool)
	configApi, configAllowTokenIndex := GetConfig(network)

	if len(tokenIndexes) > 0 {
		for _, tokenIndex := range tokenIndexes {
			if !configAllowTokenIndex[tokenIndex] {
				continue
			}
			acceptTokenIndex[tokenIndex] = true
		}
	} else {
		acceptTokenIndex = configAllowTokenIndex
	}

	return &Api{
		acceptTokenIndex: acceptTokenIndex,
		config:           configApi,
	}
}

func GetConfig(network string) (common.ConfigType, map[int]bool) {
	if network == "mainnet" {
		return common.MainnetConfig, common.MainnetAllowTokenIndex
	} else if network == "testnet" {
		return common.TestnetConfig, common.TestnetAllowTokenIndex
	}
	return common.TestnetConfig, common.TestnetAllowTokenIndex
}

// GetTokenIndexPrice retrieves the current prices of all tokens in the XOracle index.
// It returns a slice of TokenIndexPrice structs, each containing the token index and its corresponding price.
func (a *Api) GetTokenIndexPrice() ([]common.TokenIndexPrice, error) {
	url := a.config.XOracleAPI + a.config.EndpointAPIPrice
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[int]string
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	var tokenIndexPrices []common.TokenIndexPrice
	for k, v := range data {
		tokenIndex := k
		if !a.acceptTokenIndex[tokenIndex] {
			continue
		}

		price, ok := new(big.Int).SetString(v, 10)
		if !ok {
			return nil, fmt.Errorf("invalid price: %s", v)
		}

		tokenIndexPrices = append(tokenIndexPrices, common.TokenIndexPrice{
			TokenIndex: tokenIndex,
			Price:      price,
		})
	}

	// Sort by token index
	sort.Slice(tokenIndexPrices, func(i, j int) bool {
		return tokenIndexPrices[i].TokenIndex < tokenIndexPrices[j].TokenIndex
	})

	return tokenIndexPrices, nil
}

func (a *Api) GetTokenIndexInfo() (map[int]common.TokenIndexInfo, error) {
	url := a.config.XOracleAPI + a.config.EndpointAPITokenIndexInfo
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[int]map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	var tokenInfos = make(map[int]common.TokenIndexInfo)
	for k, v := range data {
		tokenInfos[k] = common.TokenIndexInfo{
			TokenName:   v["name"].(string),
			TokenSymbol: v["symbol"].(string),
		}
	}

	return tokenInfos, nil
}

func (a *Api) GetNodeInfo() ([]common.NodeInfo, error) {
	url := a.config.XOracleAPI + a.config.EndpointAPINodeInfo
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	var nodes []common.NodeInfo
	for _, v := range data {
		nodes = append(nodes, common.NodeInfo{
			NodeAddress: geth_common.HexToAddress(v["address"].(string)),
			NodeName:    v["name"].(string),
		})
	}

	return nodes, nil
}

func (a *Api) GetTokenIndexPriceByTimestamp(timestamp int64) ([]common.TokenIndexPrice, error) {
	url := a.config.XOracleAPI + a.config.EndpointAPIPrice
	if timestamp > 0 {
		url += fmt.Sprintf("?timestamps=%d", timestamp)
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[int]string
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	var tokenIndexPrices []common.TokenIndexPrice
	for k, v := range data {
		tokenIndex := k
		if !a.acceptTokenIndex[tokenIndex] {
			continue
		}

		price, ok := new(big.Int).SetString(v, 10)
		if !ok {
			return nil, fmt.Errorf("invalid price: %s", v)
		}

		tokenIndexPrices = append(tokenIndexPrices, common.TokenIndexPrice{
			TokenIndex: tokenIndex,
			Price:      price,
		})
	}

	// Sort by token index
	sort.Slice(tokenIndexPrices, func(i, j int) bool {
		return tokenIndexPrices[i].TokenIndex < tokenIndexPrices[j].TokenIndex
	})

	return tokenIndexPrices, nil
}
