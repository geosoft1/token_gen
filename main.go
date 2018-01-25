package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/geosoft1/token"
)

func main() {
	ip := flag.String("ip", "", "ip")
	port := flag.Int("port", 8084, "port")
	token_len := flag.Int("len", 6, "token length")
	flag.Parse()
	http.HandleFunc("/token",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, token.GetToken(*token_len))
		})
	http.ListenAndServe(fmt.Sprintf("%s:%d", *ip, *port), nil)
}
