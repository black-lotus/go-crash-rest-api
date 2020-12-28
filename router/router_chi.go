package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type ChiRouter struct {
}

var (
	chiDispatcher *chi.Mux = chi.NewRouter()
)

func NewChiRouter() *ChiRouter {
	return &ChiRouter{}
}

func (*ChiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*ChiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*ChiRouter) SERVE(port string) {
	fmt.Printf("Chi HTTP server running on port %v \n", port)
	http.ListenAndServe(port, muxDispatcher)
}
