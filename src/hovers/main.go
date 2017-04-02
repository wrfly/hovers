package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
	"golang.org/x/net/proxy"
)

func main() {
	PROXY_ADDR := flag.String("socks5", "127.0.0.1:1080", "socks5 proxy")
	verbose := flag.Bool("v", false, "verbose")
	addr := flag.String("p", "8080", "proxy listen port")
	flag.Parse()

	// create a socks5 dialer
	dialer, err := proxy.SOCKS5("tcp", *PROXY_ADDR, nil, proxy.Direct)
	if err != nil {
		log.Fatalf("can't connect to the proxy: %v", err)
		return
	}

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose
	proxy.Tr.Dial = dialer.Dial
	log.Fatal(http.ListenAndServe(":"+*addr, proxy))
}
