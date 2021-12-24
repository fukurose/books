package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func handlerUpdrade(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Connection") != "Upgrade" || r.Header.Get("Upgrade") != "OriginProtocol" {
		w.WriteHeader(400)
		return
	}
	fmt.Println("Upgrade to OriginProtocol")

	hijacker := w.(http.Hijacker)
	conn, readWriter, err := hijacker.Hijack()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	response := http.Response{
		StatusCode: 101,
		Header:     make(http.Header),
	}
	response.Header.Set("Upgrade", "OriginalProtocol")
	response.Header.Set("Connection", "Upgrade")
	response.Write(conn)

	for i := 1; i <= 10; i++ {
		fmt.Fprintf(readWriter, "%d\n", i)
		fmt.Println("->", i)
		readWriter.Flush()
		recv, err := readWriter.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("<- %s", string(recv))
		time.Sleep(500 * time.Millisecond)
	}

}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handlerUpdrade)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
