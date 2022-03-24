package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/xiaobo9/gopl/ch1/lissajous"
)

var mu sync.Mutex
var count int

func Server() {
	http.HandleFunc("/", handler) //each request calls handler
	http.HandleFunc("/count", counter)
	http.HandleFunc("/file", handlerFormFile)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/gif", func(w http.ResponseWriter, r *http.Request) {
		lissajous.LissajousOut(w)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echos the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	formFile, fileHeader, _ := r.FormFile("asd")
	formFile.Close()
	fmt.Println(fileHeader.Filename)
	mu.Lock()
	count++
	mu.Unlock()

	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

// handler echos the Path component of the request URL r.
func handlerFormFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(2)
	file, _, err := r.FormFile("file")

	// osFile := (*os.File)(unsafe.Pointer(&file))
	// fmt.Printf("tmp file name: %s", osFile.Name())

	if err != nil {
		fmt.Println(err)
		return
	}
	size, err := io.Copy(ioutil.Discard, file)
	// ... rest of your handler
	fmt.Printf("readfile, size: %v, err: %v", size, err)
}
