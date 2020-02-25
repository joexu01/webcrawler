package parser

import (
	"regexp"
	"webcrawler/engine"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)
func ParseCity(contents []byte) engine.ParseResult {
	// [^>]* -- zero or more character(s) but >
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, match := range matches {
		name := string(match[2])
		//result.Items = append(result.Items, "User " + name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(match[1]),
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents, name)
			},
		})
		//fmt.Printf("City: %s  URL: %s\n", match[2], match[1])
		//break
	}
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
