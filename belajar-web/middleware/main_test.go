package server

import (
	"fmt"
	_ "net"
	"net/http"
	"testing"
)

type LogMiddleware struct{
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request){
	fmt.Println("Before Execute Handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Execute Handler")
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprintf(writer, "Hello Middleware")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: logMiddleware,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}