package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func main() {
	director := func(request *http.Request) {
		request.URL.Scheme = "https"
		request.URL.Host = ":18443"
		request.URL.Path = "/directed"
	}
	modifier := func(res *http.Response) error {
		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			return fmt.Errorf("reading body error: %w", err)
		}

		newBody := bytes.NewBuffer(body)
		newBody.WriteString("via Proxy")
		res.Body = ioutil.NopCloser(newBody)
		res.Header.Set("Content-Lngth", strconv.Itoa(newBody.Len()))
		return nil
	}
	rp := &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifier,
	}
	server := http.Server{
		Addr:    "0.0.0.0:18888",
		Handler: rp,
	}
	log.Println("Origin server start at :18888")
	log.Fatalln(server.ListenAndServe())
}
