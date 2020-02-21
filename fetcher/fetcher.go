package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetcher(url string) ([]byte, error) {
	resp, err := http.Get("url")
	if err != nil {
		return nil, err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d\n", resp.StatusCode)
	}

	// alternative:
	// if charset is gbk, we can simply replace the following
	// resp.Body with utf8Reader instead of calling function determineCharset
	// utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewEncoder())

	e := determineCharset(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewEncoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineCharset(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
