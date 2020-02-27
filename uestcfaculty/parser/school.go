package parser

import (
	"regexp"
	"webcrawler/engine"
)

const schoolListJsp = `http://faculty.uestc.edu.cn/xylb.jsp`

var (
	teacherRe = regexp.MustCompile(`<h2><a href="(http://faculty.uestc.edu.cn/[^/]+/zh_CN/index.htm)" target="_blank">([^<]+)</a></h2>[^<]+<p>([^<]+)</p>`)
	pageUrlRe = regexp.MustCompile(`<a href="(\?totalpage=[^&]+&PAGENUM=[0-9]+&urltype=[^&]+&wbtreeid=[^&]+&st=[^&]+&id=[^&]+&lang=zh_CN)" class="Next">下页</a>`)
)
func ParseSchool(contents []byte) engine.ParseResult {
	matches := teacherRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		teacherName := string(m[2])
		teacherPosition := string(m[3])
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: func(contents []byte) engine.ParseResult {
					return ParseTeacherProfile(contents, teacherName, teacherPosition)
				},
			})
	}

	matches = pageUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        schoolListJsp + string(m[1]),
				ParserFunc: ParseSchool,
			})
	}
	return result
}