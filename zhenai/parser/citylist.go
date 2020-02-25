package parser

import (
	"regexp"
	"webcrawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re, _ := regexp.Compile(
		cityListRe)
	// [^>]* -- zero or more character(s) but >
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, match := range matches {
		//result.Items = append(result.Items, "City " + string(match[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(match[1]),
			ParserFunc: ParseCity,
		})
		//fmt.Printf("City: %s  URL: %s\n", match[2], match[1])
		//break
	}
	//fmt.Printf("Matches: %d\n", len(matches))
	return result
}
