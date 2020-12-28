package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type MuxRouter struct {
}

var (
	muxDispatcher *mux.Router = mux.NewRouter()
)

func NewMuxRouter() *MuxRouter {
	return &MuxRouter{}
}

func (*MuxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*MuxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*MuxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v \n", port)
	http.ListenAndServe(port, muxDispatcher)
}
