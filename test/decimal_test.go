package test

import (
	"github.com/shopspring/decimal"
	"log"
	"testing"
)

func TestDecimal(t *testing.T) {
	d1 := decimal.NewFromFloat(6.35)
	d2 := decimal.NewFromFloat(5.43)
	d := d1.Add(d2)
	log.Println("6.35 + 5.43")
	log.Println(d.Float64())

	d1 = decimal.NewFromFloat(6.35)
	d2 = decimal.NewFromFloat(5.43)
	d = d1.Mul(d2)
	log.Println("6.35 * 5.43")
	log.Println(d.String())

	ds1, _ := decimal.NewFromString("6.4")
	ds2, _ := decimal.NewFromString("5.3")
	d = ds1.Mul(ds2)
	log.Println("6.4 * 5.3")
	log.Println(d.String())
}
