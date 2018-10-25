package utils

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func QueryParamInt(r *http.Request, name string) int64 {
	i, _ := strconv.ParseInt(r.URL.Query().Get(name), 10, 64)
	return i
}

func PathParamInt(r *http.Request, name string) int64 {
	i, _ := strconv.ParseInt(mux.Vars(r)[name], 10, 64)
	return i
}
