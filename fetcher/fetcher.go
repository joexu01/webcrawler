package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

//var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetcher(url string) ([]byte, error) {
	//<-rateLimiter
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
	req.Header.Set("Cookie",
		`sid=86ed93c9-5a72-42bd-9426-9559c53c9846; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1582443794; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1582443794`)
	req.Header.Set("Host",
		"www.zhenai.com")
	req.Header.Set("Proxy-Connection",
		"keep-alive")


	resp, err := client.Do(req)
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

	bodyReader := bufio.NewReader(resp.Body)
	e := determineCharset(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewEncoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineCharset(r *bufio.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
