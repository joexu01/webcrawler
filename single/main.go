package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		_, _ = fmt.Fprintf(os.Stderr, "Error: status code %d", resp.StatusCode)
		return
	}

	// alternative:
	// if charset is gbk, we can simply replace the following
	// resp.Body with utf8Reader
	// utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewEncoder())

	e, _ := determineCharset(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewEncoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", all)
}

func determineCharset(r io.Reader) (encoding.Encoding, string) {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return nil, ""
	}
	e, name, _ := charset.DetermineEncoding(bytes, "")
	return e, name
}
