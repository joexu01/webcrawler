package main

import (
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
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
	// resp.Body with utf8Reader instead of calling function determineCharset
	// utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewEncoder())

	e, _ := determineCharset(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewEncoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	printCityList(all)
}



func printCityList(contents []byte) {
	re, _ := regexp.Compile(
		`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	// [^>]* -- zero or more character(s) but >
	matches := re.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		fmt.Printf("City: %s  URL: %s\n", match[2], match[1])
	}
	fmt.Printf("Matches: %d\n", len(matches))
}
