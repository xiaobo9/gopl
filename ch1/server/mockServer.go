package server

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

type MockServer struct {
	client    *http.Client
	remoteURL string
	server    *http.Server
}

func (mock MockServer) closeServer() {
	log.Println("close mock server")
	mock.server.Close()
}

func Mock() *MockServer {
	// setup a mocked http client.
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// b, err := httputil.DumpRequest(r, true)
			_, err := httputil.DumpRequest(r, true)
			if err != nil {
				panic(err)
			}
			// fmt.Printf("%s\n", b)
			log.Println("mock server handler end")
		}))
	log.Println("mock server started: ", server.URL)
	var mock MockServer
	mock.client = server.Client()
	mock.remoteURL = server.URL
	return &mock
}
