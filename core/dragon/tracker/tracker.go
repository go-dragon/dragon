package tracker

import (
	"github.com/go-dragon/util"
	"net/http"
	"net/url"
	"time"
)

// resp header tracker key
const TrackKey = "dragonTrack"

type Tracker struct {
	SpanId    string                 `json:"span_id"`
	Uri       string                 `json:"uri"`
	Method    string                 `json:"method"`
	ReqHeader http.Header            `json:"req_header"`
	RawJson   map[string]interface{} `json:"raw_json"`
	Form      url.Values             `json:"form"`
	Resp      struct {
		Header http.Header `json:"header"`
		Data   string      `json:"data"`
	} `json:"resp"`
	HttpClient struct {
		Req struct {
			Uri    string `json:"uri"`
			string `json:"body"`
			Params map[string]string `json:"params"`
		} `json:"req"`
		//Req *http.Request `json:"reqdata"`
		Resp string `json:"resp"`
	} `json:"httpclient"`
	StartTime time.Time   `json:"start_time"`
	DateTime  string      `json:"date_time"`
	CostTime  string      `json:"cost_time"`
	ErrInfo   interface{} `json:"err_info"`
}

func (tracker *Tracker) Marshal() string {
	m, _ := util.FastJson.Marshal(tracker)
	return string(m)
}

func UnMarshal(s string) *Tracker {
	track := Tracker{}
	util.FastJson.Unmarshal([]byte(s), &track)
	return &track
}
