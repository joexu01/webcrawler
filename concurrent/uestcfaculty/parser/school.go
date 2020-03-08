package parser

import (
	"regexp"
	"webcrawler/concurrent/config"
	"webcrawler/concurrent/engine"
)

const schoolListJsp = `http://faculty.uestc.edu.cn/xylb.jsp`

var (
	teacherRe = regexp.MustCompile(`<h2><a href="(http://faculty.uestc.edu.cn/[^/]+/zh_CN/index.htm)" target="_blank">([^<]+)</a></h2>[^<]+<p>([^<]+)</p>`)
	pageUrlRe = regexp.MustCompile(`<a href="(\?totalpage=[^&]+&PAGENUM=[0-9]+&urltype=[^&]+&wbtreeid=[^&]+&st=[^&]+&id=[^&]+&lang=zh_CN)" class="Next">下页</a>`)
)

func ParseSchool(contents []byte, _ string) engine.ParseResult {
	matches := teacherRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		teacherName := string(m[2])
		teacherPosition := string(m[3])
		result.Requests = append(result.Requests,
			engine.Request{
				Url:    url,
				Parser: NewProfileParser(teacherName, teacherPosition),
			})
	}

	matches = pageUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:    schoolListJsp + string(m[1]),
				Parser: engine.NewFuncParser(ParseSchool, config.ParseSchool),
			})
	}
	return result
}
