# xOracle go sdk
The `xoracle-go-sdk` package provides a Go client for the xOracle API.

## Installation
To install the package, you need to have Go installed on your system. You can download and install Go from the official website: `https://golang.org/dl/`

Once you have Go installed, you can install the package using the following command:
```javascript
go get github.com/Crown-Labs/xoracle-go-sdk
```
This will download and install the package and its dependencies.

## Available Functions
- `GetTokenIndexPrice()`
- `GetTokenAddressPrice(networkId int)`
- `GetTokenIndexInfo()`
- `GetNodeInfo()`

## Usage
To use the package, you need to import it in your Go code:
```javascript
package main

import (
	"fmt"

	"github.com/Crown-Labs/xoracle-go-sdk/api"
	"github.com/Crown-Labs/xoracle-go-sdk/common"
)

func main() {
	// Create new xoracle instance
	xOracle := api.NewApi([]int{
		common.TOKEN_INDEX.BTC,
		common.TOKEN_INDEX.ETH,
		common.TOKEN_INDEX.BNB,
		common.TOKEN_INDEX.BUSD,
		common.TOKEN_INDEX.USDC,
		common.TOKEN_INDEX.MATIC,
		common.TOKEN_INDEX.OP,
		common.TOKEN_INDEX.ARB,
	})

	// Get token prices
	networkId := common.NETWORK_ID.LINEA_TESTNET
	tokenAddressPrices, err := xOracle.GetTokenAddressPrice(networkId)
	if err != nil {
		fmt.Println(err)
	}

	// Logs token prices
	for _, price := range tokenAddressPrices {
		fmt.Printf("TokenAddress: %s, Price: %s, Price 30 Decimals: %s\n",
			price.TokenAddress,
			price.Price.String(),
			common.ConvertPriceDecimals(price.Price, 30).String(),
		)
	}
}
```

In this code, we import the `xoracle-go-sdk` package, create an instance of the `xOracle Api` using the `NewApi` function, and use its methods to interact with the xOracle API. We get the token index prices using the `GetTokenIndexPrice` method, the token address prices for a specific network using the `GetTokenAddressPrice` method, and the token index info using the `GetTokenIndexInfo` method.

## Contributing
If you want to contribute to the project, you can fork the repository and create a pull request with your changes. Please make sure to follow the coding style and conventions used in the project.

## License
This project is licensed under the GPL-3.0 license. See the LICENSE file for details.