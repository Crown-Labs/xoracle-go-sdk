package api

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/Crown-Labs/xoracle-go-sdk/common"
)

type Api struct {
	acceptTokenIndex map[int]bool
	config           common.ConfigType
}

// NewApi returns a new Api instance.
// It accepts a slice of token indexes, which are used to filter the results of the GetTokenIndexPrice and GetTokenAddressPrice functions.
func NewApi(tokenIndexes []int) *Api {
	var acceptTokenIndex = make(map[int]bool)

	if len(tokenIndexes) > 0 {
		for _, tokenIndex := range tokenIndexes {
			if !common.AllowTokenIndex[tokenIndex] {
				continue
			}
			acceptTokenIndex[tokenIndex] = true
		}
	} else {
		acceptTokenIndex = common.AllowTokenIndex
	}

	return &Api{
		acceptTokenIndex: acceptTokenIndex,
		config:           common.Config,
	}
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

	var data map[string]string
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	var tokenIndexPrices []common.TokenIndexPrice
	for k, v := range data {
		tokenIndex, err := strconv.Atoi(k)
		if err != nil {
			return nil, err
		}

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

	return tokenIndexPrices, nil
}

// GetTokenAddressPrice returns an array of TokenAddressPrice structs, which contain the token address and its price for a given network ID.
func (a *Api) GetTokenAddressPrice(networkId int) ([]common.TokenAddressPrice, error) {
	tokenIndexPrices, err := a.GetTokenIndexPrice()
	if err != nil {
		return nil, err
	}

	var tokenAddressPrice []common.TokenAddressPrice
	for _, tokenIndexPrice := range tokenIndexPrices {
		address := a.config.Chains[networkId].TokenAddress[tokenIndexPrice.TokenIndex]
		if address == "" {
			continue
		}

		tokenAddressPrice = append(tokenAddressPrice, common.TokenAddressPrice{
			TokenAddress: address,
			Price:        tokenIndexPrice.Price,
		})
	}

	return tokenAddressPrice, nil
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
			NodeAddress: v["address"].(string),
			NodeName:    v["name"].(string),
		})
	}

	return nodes, nil
}
