package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/geosoft1/token"
)

var (
	port      = flag.String("port", "8080", "port")
	token_len = flag.Int("len", 8, "token length")
)

func main() {
	flag.Parse()
	http.HandleFunc("/token_gen",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, token.GetToken(*token_len))
		})
	http.ListenAndServe(":"+*port, nil)
}
