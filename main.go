package main

import (
	"flag"
	"fmt"
	"github.com/elazarl/goproxy"
	"net/http"
)

func main() {
	port := 3128
	verbose := false
	flag.IntVar(&port, "p", port, "port number. default [3128]")
	flag.BoolVar(&verbose, "v", verbose, "verbose output. default [false]")
	proxy := goproxy.NewProxyHttpServer()
	if verbose {
		proxy.Verbose = true
	}
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), proxy)
	if err != nil {
		proxy.Logger.Fatal(err)
	}
}
