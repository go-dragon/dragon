package test

import (
	"dragon/httpclient"
	"fmt"
	"log"
	"testing"
)

// test GET
func TestGET(t *testing.T) {
	res := httpclient.DefaultClient.GET("https://qwu.zero-w.cn", nil, nil)

	if res.Err != nil {
		log.Println(res.Err)
	}
	fmt.Println(res.Content)
	fmt.Println(res.Status)
}

// benchmark GET
func BenchmarkGET(b *testing.B) {
	cli := httpclient.DefaultClient
	for i := 0; i < b.N; i++ {
		res := cli.GET("https://qwu.zero-w.cn/", nil, nil)
		if res.Err != nil {
			log.Println(res.Err)
		}
		//log.Println(res.Content)
	}
}

// test POST
func TestPOST(t *testing.T) {
	srv := httpclient.DefaultClient
	res := srv.POST("https://qwu.zero-w.cn/", map[string]string{
		"cur_code": "440300000000",
		"page":     "0",
		"pageSize": "10",
	}, nil)
	if res.Err != nil {
		log.Println(res.Err)
	}
	fmt.Println(res.Content)
	fmt.Println(res.Status)
}

func TestPUT(t *testing.T) {
	srv := httpclient.DefaultClient
	res := srv.PUT("http://talent.qh-1.cn/pc/httpclient/orgs", map[string]string{
		"cur_code": "440300000000",
		"page":     "0",
		"pageSize": "10",
	}, nil)
	if res.Err != nil {
		log.Println(res.Err)
	}
	fmt.Println(res.Content)
	fmt.Println(res.Status)
}

func TestPATCH(t *testing.T) {
	srv := httpclient.DefaultClient
	res := srv.PATCH("http://talent.qh-1.cn/pc/httpclient/orgs", map[string]string{
		"cur_code": "440300000000",
		"page":     "0",
		"pageSize": "10",
	}, nil)
	if res.Err != nil {
		log.Println(res.Err)
	}
	fmt.Println(res.Content)
	fmt.Println(res.Status)
}

func TestDELETE(t *testing.T) {
	srv := httpclient.DefaultClient
	res := srv.DELETE("http://talent.qh-1.cn/pc/httpclient/orgs", map[string]string{
		"cur_code": "440300000000",
		"page":     "0",
		"pageSize": "10",
	}, nil)
	if res.Err != nil {
		log.Println(res.Err)
	}
	fmt.Println(res.Content)
	fmt.Println(res.Status)
}

func TestClient_POSTJson(t *testing.T) {
	cli := httpclient.DefaultClient
	rsp := cli.POSTJson("https://www.baidu.com/", `{"x":1, "y":2}`)
	log.Println(rsp.Content)
}
