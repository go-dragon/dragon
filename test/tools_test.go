package test

import (
	"dragon/core/dragon/conf"
	"dragon/tools"
	"fmt"
	"github.com/go-dragon/util"
	"github.com/go-playground/validator/v10"
	"log"
	"sync"
	"testing"
)

// test
func TestTOPSign(t *testing.T) {

}

func TestFastJson(t *testing.T) {
	data := map[string]interface{}{
		"x": 1,
		"y": "world",
		"a": "world",
		"b": "world",
		"c": "world",
		"d": "world",
		"e": "world",
		"f": "world",
	}
	var wg sync.WaitGroup
	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go func() {
			util.FastJson.Marshal(&data)
			wg.Done()
		}()
	}
	wg.Wait()
	log.Println(111)
}

func BenchmarkFastJson(b *testing.B) {
	data := map[string]interface{}{
		"x": 1,
		"y": "world",
		"a": "world",
		"b": "world",
		"c": "world",
		"d": "world",
		"e": "world",
		"f": "world",
	}
	for i := 0; i < b.N; i++ {
		util.FastJson.Marshal(&data)
	}
}

func BenchmarkFastJsonDecode(b *testing.B) {
	data := `{"x":1, "y":"hello world"}`
	var res map[string]interface{}
	for i := 0; i < b.N; i++ {
		util.FastJson.Unmarshal([]byte(data), &res)
	}
	log.Println("res", fmt.Sprintf("%+v", res))
}

func TestUUidV4(t *testing.T) {
	const size = 10000000
	car := make(map[string]int, size)
	for i := 0; i < size; i++ {
		uuid := tools.UUidV4()
		if _, ok := car[uuid]; ok {
			// 如果uuid重复则报错
			log.Fatal("uuid repeat", uuid)
		}
		car[uuid] = 1
	}
	log.Println("uuidV4 generate success")
}

func BenchmarkUUidV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tools.UUidV4()
	}
}

func TestGetClientIp(t *testing.T) {
	conf.InitConf()
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "TestGetClientIp1", want: "192.168.31.112"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tools.GetClientIp()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClientIp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetClientIp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidator(t *testing.T) {
	var str = []byte(`{
    "MemberId":"4",
    "Cols":"DeliveryAddress,SkinType,AgeMin,AgeMax,Job,GoodAt,Wechat,CreateTime,UpdateTime"
}`)
	type validateData struct {
		MemberId int64  `validate:"required"`
		Cols     string `validate:"required"`
	}
	var data validateData
	err := util.FastJson.Unmarshal(str, &data)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("data", data)

	v := validator.New()
	err = v.Struct(&data)
	log.Println("err", err)
}
