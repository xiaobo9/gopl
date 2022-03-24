package fetch

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/xiaobo9/gopl/utils"
)

func fetch1(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}

func fetch2(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(resp.Status)

	written, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("输出了 %v", written)
}
func FetchSingle() {
	for _, url := range os.Args[1:] {
		fetch2(utils.FixUrlPrefix(url))
	}
}
