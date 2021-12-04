package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func get() {
	values := url.Values{
		"query": {"hello world"},
	}
	resp, err := http.Get("http://localhost:18888" + "?" + values.Encode())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	log.Println(resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}

func post() {
	values := url.Values{
		"test": {"value"},
	}
	resp, err := http.PostForm("http://localhost:18888", values)
	if err != nil {
		panic(err)
	}
	log.Println("status:", resp.Status)
}

func main() {
	get()
	post()
}
