package middleware

import (
	"bytes"
	"dragon/core/dragon/tracker"
	"github.com/go-dragon/util"
	"io/ioutil"
	"net/http"
	"time"
)

// 这里的next是router.go里Routes
func LogInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		spanId, _ := util.NewUUID()
		r.ParseForm()
		// 读取
		body, _ := ioutil.ReadAll(r.Body)
		var rawJsonParams map[string]interface{}
		util.FastJson.Unmarshal(body, &rawJsonParams)
		// 把刚刚读出来的再写进去
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		trackMan := &tracker.Tracker{
			SpanId:    spanId,
			Uri:       r.RequestURI,
			Method:    r.Method,
			ReqHeader: r.Header,
			RawJson:   rawJsonParams,
			Form:      r.Form,
			StartTime: start,
			DateTime:  start.Format("2006-01-02 15:04:05"),
			CostTime:  "",
		}
		trackInfo := trackMan.Marshal()
		r.Header.Set(tracker.TrackKey, trackInfo)

		// before_req hook
		beforeReq(r, w)

		next.ServeHTTP(w, r)

		// after_req hook
		afterReq(r, w)
	})
}
