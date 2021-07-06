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

// HttpContext
type HttpContext struct {
	Request *http.Request
	Writer  http.ResponseWriter
	Params  httprouter.Params
}

// output struct
type Output struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// output data structure
type OutData struct {
	Output
	SpanId string `json:"span_id"`
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
	resp.Header().Set("content-type", "application/json; charset=utf-8")
	resp.Header().Set("x-server", "dragon")
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT,PATCH")
	resp.Header().Set("Access-Control-Allow-Credentials", "true")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Content-Length, Accept-Encoding, Origin")

	trackInfo := h.Request.Header.Get(tracker.TrackKey)
	trackMan := tracker.UnMarshal(trackInfo)
	defer func() {
		dlogger.Info(trackMan) // 最后写日志跟踪
	}()
	trackMan.Resp.Header = resp.Header()
	outData := OutData{
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
	resp.Header().Set("x-server", "dragon")
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT,PATCH")
	resp.Header().Set("Access-Control-Allow-Credentials", "true")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Content-Length, Accept-Encoding, Origin")

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

// wrap the router for controller
func WrapController(handler func(ctx *HttpContext)) httprouter.Handle {

	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		ctx := &HttpContext{
			Request: request,
			Writer:  writer,
			Params:  params,
		}
		handler(ctx)
	}
}
