package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	res.WriteHeader(http.StatusOK)

	//noinspection GoUnhandledErrorResult
	fmt.Fprintf(res, "key : %s - path : %s", vars["key"], vars["path"])
}
