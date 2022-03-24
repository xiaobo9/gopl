package server

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

const chunkSize = 256

func PostMock() {
	PostWithMockServer()
}

func PostFile() {
	SendPostRequest("http://localhost:8000/file", "temp/go.sum")
}

func PostWithMockServer() {
	// prepare the reader instances to encode
	values := map[string]io.Reader{
		"file":  openFile("post.go"), // lets assume its this file
		"other": strings.NewReader("hello world!"),
	}

	mock := Mock()
	defer mock.closeServer()
	err := upload(mock.client, mock.remoteURL, values)
	if err != nil {
		panic(err)
	}
}

func upload(client *http.Client, url string, values map[string]io.Reader) (err error) {
	// Prepare a form that you will submit to that URL.
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	for key, reader := range values {
		var fw io.Writer
		if closer, ok := reader.(io.Closer); ok {
			defer closer.Close()
		}
		// Add a file
		if file, ok := reader.(*os.File); ok {
			if fw, err = writer.CreateFormFile(key, file.Name()); err != nil {
				return
			}
		} else {
			// Add other fields
			if fw, err = writer.CreateFormField(key); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, reader); err != nil {
			return err
		}

	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	writer.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &buffer)
	if err != nil {
		fmt.Println("create request", err)
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Submit the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("client Do", err)
		return
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}
	return
}

func SendPostRequest(url string, filename string) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	bodyWriter.WriteField("name", "name")

	fillFormFile(bodyWriter, "file", filename)

	bodyWriter.Close()

	response, err := http.Post(url, bodyWriter.FormDataContentType(), bodyBuf)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(string(content))
}

// fill form file with field name from named file
func fillFormFile(writer *multipart.Writer, fieldName string, fileName string) {
	fw, err := writer.CreateFormFile(fieldName, fileName)
	if err != nil {
		panic(fmt.Sprintf("Error creating writer: %v", err))
	}

	file := openFile(fileName)
	defer file.Close()
	if _, err = io.Copy(fw, file); err != nil {
		panic(fmt.Sprintf("Error with io.Copy: %v", err))
	}

	writer.Close()
}

// 打开文件
func openFile(f string) *os.File {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return file
}
