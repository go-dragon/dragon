package dragon

import (
	"dragon/core/dragon/dlogger"
	"dragon/core/dragon/tracker"
	"fmt"
	"github.com/go-dragon/util"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// IHttpContext interface
type IHttpContext interface {
	GetReqParams() map[string]string
	BindPostJson(data interface{}) error
	Json(data *Output, statusCode int)
	String(data string, statusCode int)
}

type HttpContext struct {
	Request *http.Request
	Writer  http.ResponseWriter
	PathParams  httprouter.Params
}

// Output struct
type Output struct {
	Code int
	Msg  string
	Data interface{}
}

// output data structure
type outData struct {
	Output
	SpanId string
}

// GetReqParams get request params (get and post params)
func (h *HttpContext) GetReqParams() map[string]string {
	requests := make(map[string]string)
	var err error
	err = h.Request.ParseForm()
	if err != nil {
		log.Println(err)
		return requests
	}

	for k, v := range h.Request.Form {
		if len(v) >= 1 {
			// only select the 1st param
			requests[k] = v[0]
		}
	}

	return requests
}

// BindQueryParams parse query params (like: https://foo.com?v1=xxx&v2=xxx) bind to struct
// if query contains >=1 repeated key params, only bind first key value
func (h *HttpContext) BindQueryParams(data interface{}) error {
	var err error
	h.Request.ParseForm()
	queryParams := make(map[string]string, len(h.Request.Form))
	for k, v := range h.Request.Form{
		queryParams[k] = v[0]
	}
	if err != nil {
		return err
	}
	queryBytes, err := util.FastJson.Marshal(queryParams)
	if err != nil {
		return err
	}
	err = util.FastJson.Unmarshal(queryBytes, data)
	if err != nil {
		return err
	}
	return nil
}

// BindPostJson parse raw json bind to struct
func (h *HttpContext) BindPostJson(data interface{}) error {
	var body []byte
	var err error
	body, err = ioutil.ReadAll(h.Request.Body)
	if err != nil {
		return err
	}
	err = util.FastJson.Unmarshal(body, data)
	if err != nil {
		return err
	}
	return nil
}

func (h *HttpContext) Json(data *Output, statusCode int) {
	resp := h.Writer
	resp.Header().Set("X-Server", "dragon")
	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")
	resp.Header().Set("Access-Control-Allow-Credentials", "true")
	resp.Header().Set("Access-Control-Allow-Headers", "*")

	trackInfo := h.Request.Header.Get(tracker.TrackKey)
	trackMan := tracker.UnMarshal(trackInfo)
	defer func() {
		dlogger.Info(trackMan) // 最后写日志跟踪
	}()
	trackMan.Resp.Header = resp.Header()
	outData := outData{
		Output: *data,
		SpanId: trackMan.SpanId,
	}
	js, err := util.FastJson.Marshal(outData)
	// 生成耗时
	trackMan.CostTime = time.Since(trackMan.StartTime).String()

	if err != nil {
		trackMan.ErrInfo = err
		fmt.Fprint(resp, "error")
		return
	}
	// trackMan data log
	resp.WriteHeader(statusCode)
	trackMan.Resp.Data = string(js)

	// output
	_, err = resp.Write(js)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		trackMan.Resp.Header = resp.Header()
		trackMan.ErrInfo = err
		fmt.Fprint(resp, "")
		return
	}
}

// String send string to client
func (h *HttpContext) String(data string, statusCode int) {
	resp := h.Writer
	resp.Header().Set("X-Server", "dragon")
	resp.Header().Set("Content-Type", "text/html; charset=utf-8")
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")
	resp.Header().Set("Access-Control-Allow-Credentials", "true")
	resp.Header().Set("Access-Control-Allow-Headers", "*")

	trackInfo := h.Request.Header.Get(tracker.TrackKey)
	trackMan := tracker.UnMarshal(trackInfo)
	defer func() {
		dlogger.Info(trackMan) // 最后写日志跟踪
	}()
	trackMan.Resp.Header = resp.Header()
	// 生成耗时
	trackMan.CostTime = time.Since(trackMan.StartTime).String()
	// trackMan data log
	resp.WriteHeader(statusCode)
	trackMan.Resp.Data = data
	// output
	_, err := resp.Write([]byte(data))
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		trackMan.Resp.Header = resp.Header()
		trackMan.ErrInfo = err
		fmt.Fprint(resp, "")
		return
	}
}

// WrapController wrap the router for controller
func WrapController(handler func(ctx *HttpContext)) httprouter.Handle {

	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		ctx := &HttpContext{
			Request: request,
			Writer:  writer,
			PathParams:  params, // path params like /path/:userId, not query params!!!
		}
		handler(ctx)
	}
}
