package parser

import (
	"regexp"
	"webcrawler/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re, _ := regexp.Compile(cityRe)
	// [^>]* -- zero or more character(s) but >
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, match := range matches {
		name := string(match[2])
		result.Items = append(result.Items, "User " + name)
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(match[1]),
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents, name)
			},
		})
		//fmt.Printf("City: %s  URL: %s\n", match[2], match[1])
		//break
	}
	//fmt.Printf("Matches: %d\n", len(matches))
	return result
}
