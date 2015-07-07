package app

import (
   "fmt"
	"net/http"
)

func init () {
   http.HandleFunc("/", defaultHandler)
}

func defaultHandler (w http.ResponseWriter, r *http.Request) {
   fmt.Fprint (w, "hello world (dynamic)!")
}






