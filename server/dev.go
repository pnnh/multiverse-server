package server

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os" 

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type devResponseWriter struct {
	http.ResponseWriter
	size   int
	status int
}

var _ gin.ResponseWriter = &devResponseWriter{}

const (
	noWritten     = -1
	defaultStatus = http.StatusOK
)

func (w *devResponseWriter) reset(writer http.ResponseWriter) {
	w.ResponseWriter = writer
	w.size = noWritten
	w.status = defaultStatus
}

func (w *devResponseWriter) WriteHeader(code int) {
	if code > 0 && w.status != code {
		w.Written()
		w.status = code
	}
}

func (w *devResponseWriter) WriteHeaderNow() {
	if !w.Written() {
		w.size = 0
		w.ResponseWriter.WriteHeader(w.status)
	}
}

func (w *devResponseWriter) Write(data []byte) (n int, err error) {
	w.WriteHeaderNow()
	n, err = w.ResponseWriter.Write(data)
	w.size += n
	return
}

func (w *devResponseWriter) WriteString(s string) (n int, err error) {
	w.WriteHeaderNow()
	n, err = io.WriteString(w.ResponseWriter, s)
	w.size += n
	return
}

func (w *devResponseWriter) Status() int {
	return w.status
}

func (w *devResponseWriter) Size() int {
	return w.size
}

func (w *devResponseWriter) Written() bool {
	return w.size != noWritten
}

func (w *devResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if w.size < 0 {
		w.size = 0
	}
	return w.ResponseWriter.(http.Hijacker).Hijack()
}

func (w *devResponseWriter) CloseNotify() <-chan bool {
	return w.ResponseWriter.(http.CloseNotifier).CloseNotify()
}

func (w *devResponseWriter) Flush() {
	w.WriteHeaderNow()
	w.ResponseWriter.(http.Flusher).Flush()
}

func (w *devResponseWriter) Pusher() (pusher http.Pusher) {
	if pusher, ok := w.ResponseWriter.(http.Pusher); ok {
		return pusher
	}
	return nil
}

func debugStaticWebHandler(gctx *gin.Context) {
	// target := "localhost:8080"
	// //devUrl := fmt.Sprintf("http://localhost:3000%s", realPath)
	// proxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
	// 	req.URL.Scheme = "http"
	// 	req.URL.Host = target
	// 	req.URL.Path = r.URL.Path //"/svc/css/index.scss"
	// 	//r.Host = target
	// }}

	// proxy.ServeHTTP(w, r)

	rs1 := "http://127.0.0.1:3500"
	targetUrl, err := url.Parse(rs1)
	if err != nil {
		logrus.Fatalln("debugStaticWebHandler: %w", err)
		return
	}
	httputil.NewSingleHostReverseProxy(targetUrl).ServeHTTP(gctx.Writer, gctx.Request)

	//gctx.Redirect(http.StatusFound, devUrl)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func staticWebHandler(gctx *gin.Context) { 
	path := gctx.Request.URL.Path

	if path == "/favicon.ico" {
		http.ServeFile(gctx.Writer, gctx.Request, "browser/images/favicon.ico")
		return
	}

	path = "browser/web" + path
	if fileExists(path) {
		http.ServeFile(gctx.Writer, gctx.Request, path)
	} else {
		http.ServeFile(gctx.Writer, gctx.Request, "browser/web/index.html")
	}
}
