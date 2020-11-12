package tools

import (
	"github.com/shopspring/decimal"
	"math/big"
)

func BigIntStrToFloatStr(number string, p int32) string {
	n := new(big.Int)
	n, _ = n.SetString(number, 10)
	precision := new(big.Int)
	switch p {
	case 6:
		precision, _ = new(big.Int).SetString("1000000", 10)
	case 8:
		precision, _ = new(big.Int).SetString("100000000", 10)
	case 18:
		precision, _ = new(big.Int).SetString("1000000000000000000", 10)
	}
	bf := decimal.NewFromBigInt(n, 10).DivRound(decimal.NewFromBigInt(precision, 10), p)
	return bf.String()
}
