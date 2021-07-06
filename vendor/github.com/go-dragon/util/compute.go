package util

import (
	"github.com/shopspring/decimal"
	"regexp"
)

// 找出字符串中所有数字, 如果没有匹配的，则返回nil
func FindNumsInString(s string) []string {
	re := regexp.MustCompile("[0-9]+")
	return re.FindAllString(s, -1)
}

// 截取2位小数，四舍五不入。 这里通过 ✖️100取math.Floor再➗100方式实现的两位小数
func TwoDecimalPlaces(value float64) float64 {
	dc := decimal.NewFromFloat(value)
	dc = dc.Mul(decimal.NewFromFloat(100))
	dc = dc.Floor().Div(decimal.NewFromFloat(100))
	res, _ := dc.Float64()
	return res
}
