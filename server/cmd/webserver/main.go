package webserver

import "net/http"

type WebServer struct {
	ServeMux http.ServeMux
}

func (ws *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws.ServeMux.ServeHTTP(w, r)
}

func InitializeWebServer() *WebServer {
	// Initialize the web server and return it
	return &WebServer{}
}
