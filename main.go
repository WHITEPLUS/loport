package main

import (
	"flag"
	"fmt"
	"github.com/elazarl/goproxy"
	"net/http"
)

func main() {
	port := flag.Int("p", 3128, "port number. default [3128]")
	verbose := flag.Bool("v", false, "verbose output. default [false]")
	flag.Parse()

	proxy := goproxy.NewProxyHttpServer()
	if *verbose {
		proxy.Verbose = true
	}
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), proxy)
	if err != nil {
		proxy.Logger.Printf("%s\n", err.Error())
	}
}
