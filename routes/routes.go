package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wanderporto/api-json-crud/utils"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(jsonMiddleware)

	addBookHandler(r)
	addUserHandler(r)

	return r
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func httpInfo(r *http.Request) {
	fmt.Printf("%s\t %s\t %s%s\t %s\n", r.Method, r.Proto, r.Host, r.URL, utils.GetDateTime())
}
