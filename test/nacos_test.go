package test

import (
	"dragon/core/dragon"
	"dragon/core/dragon/conf"
	"dragon/core/dragon/dnacos"
	"log"
	"testing"
)

func TestSelectOneHealthyInstance(t *testing.T) {
	//init config
	conf.InitConf()
	dragon.AppInit()
	addr, ins, err := dnacos.SelectOneHealthyInstance("dragon", "", nil)
	log.Println("addr", addr)
	log.Printf("%+v\n", ins)
	log.Println("err", err)
	if err != nil {
		log.Fatalln(err)
	}
}
