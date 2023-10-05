package common

import (
	"math/big"
	"testing"
)

func TestConvertPriceDecimals(t *testing.T) {
	price8 := big.NewInt(100000000) // 8 decimals

	price6 := ConvertPriceDecimals(price8, 6)
	expectedPrice6 := big.NewInt(1000000) // 6 decimals
	if price6.Cmp(expectedPrice6) != 0 {
		t.Errorf("ConvertPriceDecimals returned %v, expected %v", price6, expectedPrice6)
	}

	price18 := ConvertPriceDecimals(price8, 18)
	expectedPrice18 := big.NewInt(1000000000000000000) // 18 decimals
	if price18.Cmp(expectedPrice18) != 0 {
		t.Errorf("ConvertPriceDecimals returned %v, expected %v", price18, expectedPrice18)
	}

	price30 := ConvertPriceDecimals(price8, 30)
	expectedPrice30 := new(big.Int).Exp(BigInt10, big.NewInt(30), nil) // 30 decimals
	if price30.Cmp(expectedPrice30) != 0 {
		t.Errorf("ConvertPriceDecimals returned %v, expected %v", price30, expectedPrice30)
	}

	price8_2 := ConvertPriceDecimals(price8, 8)
	expectedPrice8 := big.NewInt(100000000) // 8 decimals
	if price8_2.Cmp(expectedPrice8) != 0 {
		t.Errorf("ConvertPriceDecimals returned %v, expected %v", price8_2, expectedPrice8)
	}
}
