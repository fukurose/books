package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	proxyURL, err := url.Parse("http://localhost:18888")
	if err != nil {
		panic(err)
	}

	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
	log.Println(proxyURL)
	resp, err := client.Get("https://exmaple.com")
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
