package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T){

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Not Found")
	})

	router.GET("/panic", func (writer http.ResponseWriter, request *http.Request, pararms httprouter.Params)  {
		panic("error")
	})

	request := httptest.NewRequest("GET", "http://localhost:8000/404", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Not Found", string(body))
}