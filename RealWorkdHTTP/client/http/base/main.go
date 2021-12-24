package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
)

const LOCALHOST = "http://localhost:18888"

func get() {
	values := url.Values{
		"query": {"hello world"},
	}
	resp, err := http.Get(LOCALHOST + "?" + values.Encode())
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
	resp, err := http.PostForm(LOCALHOST, values)
	if err != nil {
		panic(err)
	}
	log.Println("status:", resp.Status)
}

func postFile() {
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(LOCALHOST, "text/plain", file)
	if err != nil {
		panic(err)
	}
	log.Println("status:", resp.Status)
}

func muiltiPost() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")

	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Dispositon", `form-data; name="thumbnail"; filename="photo.jpg"`)

	fileWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post(LOCALHOST, writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("status:", resp.Status)
}

func main() {
	get()
	post()
	postFile()
	muiltiPost()
}
