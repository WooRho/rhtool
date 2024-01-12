package rhtool_math

import (
	"github.com/shopspring/decimal"
)

func DecimalAdd(x, y float64) float64 {
	xDecimal := decimal.NewFromFloat(x)
	yDecimal := decimal.NewFromFloat(y)
	z, _ := xDecimal.Add(yDecimal).Float64()
	return z
}

func DecimalSub(x, y float64) float64 {
	xDecimal := decimal.NewFromFloat(x)
	yDecimal := decimal.NewFromFloat(y)
	z, _ := xDecimal.Sub(yDecimal).Float64()
	return z
}

func DecimalMul(x, y float64) float64 {
	xDecimal := decimal.NewFromFloat(x)
	yDecimal := decimal.NewFromFloat(y)
	z, _ := xDecimal.Mul(yDecimal).Float64()
	return z
}

func DecimalDiv(x, y float64) float64 {
	xDecimal := decimal.NewFromFloat(x)
	yDecimal := decimal.NewFromFloat(y)
	z, _ := xDecimal.Div(yDecimal).Float64()
	return z
}
