package common

import (
	"math/big"

	geth_common "github.com/ethereum/go-ethereum/common"
)

func ParseEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), BigFloatBase18)
}

func ParseAddressFormat(address string) string {
	return geth_common.HexToAddress(address).Hex()
}

func AppendUrlWithSlash(url string) string {
	if url[len(url)-1:] != "/" {
		url += "/"
	}
	return url
}

func ConvertPriceDecimals(price *big.Int, decimals int) *big.Int {
	if decimals > TOKEN_DECIMALS {
		return new(big.Int).Mul(price, new(big.Int).Exp(BigInt10, big.NewInt(int64(decimals-TOKEN_DECIMALS)), nil))
	} else if decimals < TOKEN_DECIMALS {
		return new(big.Int).Div(price, new(big.Int).Exp(BigInt10, big.NewInt(int64(TOKEN_DECIMALS-decimals)), nil))
	}
	return price
}
