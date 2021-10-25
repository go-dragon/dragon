package middleware

import (
	"bytes"
	"dragon/core/dragon/tracker"
	"github.com/go-dragon/util"
	"io/ioutil"
	"net/http"
	"time"
)

func beforeBasicHandle(request *http.Request, response http.ResponseWriter) error {
	start := time.Now()
	spanId, _ := util.NewUUID()
	request.ParseForm()
	// 读取
	body, _ := ioutil.ReadAll(request.Body)
	var rawJsonParams map[string]interface{}
	util.FastJson.Unmarshal(body, &rawJsonParams)
	// 把刚刚读出来的再写进去
	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	trackMan := &tracker.Tracker{
		SpanId:    spanId,
		Uri:       request.RequestURI,
		Method:    request.Method,
		ReqHeader: request.Header,
		RawJson:   rawJsonParams,
		Form:      request.Form,
		StartTime: start,
		DateTime:  start.Format("2006-01-02 15:04:05"),
		CostTime:  "",
	}
	trackInfo := trackMan.Marshal()
	request.Header.Set(tracker.TrackKey, trackInfo)
	return nil
}

func beforeHandle(request *http.Request, response http.ResponseWriter) error {
	// 做一些如：日志跟踪，时间记录的基础处理
	beforeBasicHandle(request, response)
	//  todo 这里可以做限流的处理
	return nil
}
