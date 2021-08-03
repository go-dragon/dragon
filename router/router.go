package router

/**
 * @link https://github.com/julienschmidt/httprouter
 */
import (
	"dragon/core/dragon"
	"dragon/core/dragon/dlogger"
	"dragon/core/dragon/tracker"
	"dragon/handler"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

type notFoundHandler struct {
}

var (
	Routes      *httprouter.Router
	userHandler handler.IUserHandler
)

func init() {
	userHandler = &handler.UserHandler{} //product handler 初始化
	Routes = httprouter.New()
	Routes.NotFound = notFoundHandler{}
	Routes.PanicHandler = panicHandler
	dRouter := dragon.NewDRouter(Routes)
	// -----------------------------用户相关-----------------------------
	dRouter.POST("/test", userHandler.Test)
	// -----------------------------用户相关-----------------------------
}

// not found route handle
func (notFoundHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "text/html; charset=utf-8")
	resp.Header().Set("X-Server", "dragon")
	trackInfo := req.Header.Get(tracker.TrackKey)
	trackMan := tracker.UnMarshal(trackInfo)
	trackMan.Resp.Header = resp.Header()
	trackMan.Resp.Data = "<h2>Dragon Not Found</h2>"
	trackMan.CostTime = time.Since(trackMan.StartTime).String()
	dlogger.Info(trackMan) // 最后写日志跟踪
	resp.Write([]byte("<h2>Dragon Not Found</h2>"))
}

// all panic handler
func panicHandler(resp http.ResponseWriter, req *http.Request, err interface{}) {
	resp.Header().Set("Content-Type", "text/html; charset=utf-8")
	resp.Header().Set("X-Server", "dragon")
	resp.WriteHeader(http.StatusInternalServerError)
	trackInfo := req.Header.Get(tracker.TrackKey)
	trackMan := tracker.UnMarshal(trackInfo)
	trackMan.Resp.Header = resp.Header()
	trackMan.Resp.Data = "<h2>500 Internal Server Error</h2>"
	trackMan.ErrInfo = err
	dlogger.Error(trackMan) // 写入日志跟踪
	resp.Write([]byte("<h2>500 Internal Server Error</h2>"))
	if err != nil {
		log.Println(err)
	}
	recover()
}
