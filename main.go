package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {
	ping := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("ping")
		writer.Write([]byte("pong"))
	})

	slow := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("call slow api")

		time.Sleep(4 * time.Second)

		// very long string
		sb := strings.Builder{}
		str := fmt.Sprintf("this is a very long long string %d ", rand.Int())
		for i := 1; i < 10000; i++ {
			sb.WriteString(str)
		}
		writer.Write([]byte(sb.String()))

		fmt.Println("end api")
	})

	fast := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("call fast api")
		writer.Write([]byte("done"))
	})

	mux := http.NewServeMux()
	mux.Handle("/ping", ping)
	mux.Handle("/slow", slow)
	mux.Handle("/fast", fast)

	server := &http.Server{
		Addr:    "0.0.0.0:3000",
		Handler: mux,
	}
	fmt.Println("starting server ...")
	server.ListenAndServe()
}
