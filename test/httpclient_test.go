package test

import (
	"dragon/httpclient"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"
)

// test GET
func TestClient_GET(t *testing.T) {
	type fields struct {
		HttpCli *http.Client
	}
	type args struct {
		url     string
		params  map[string]string
		headers map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *httpclient.Response
	}{
		{
			name:   "test GET https://qwu.zero-w.cn/",
			fields: fields{HttpCli: http.DefaultClient},
			args:   args{url: "https://qwu.zero-w.cn/", params: nil, headers: nil},
			want: &httpclient.Response{
				Content: `{"code":200,"result":null,"msg":null}`,
				Status:  http.StatusOK,
				Err:     nil,
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &httpclient.Client{
				HttpCli: tt.fields.HttpCli,
			}
			if got := c.GET(tt.args.url, tt.args.params, tt.args.headers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GET() = %v, want %v", got, tt.want)
			}
		})
	}

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
