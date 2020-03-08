package parser

import (
	"regexp"
	"webcrawler/concurrent/config"
	"webcrawler/concurrent/engine"
)

const domain = `http://faculty.uestc.edu.cn/`

var schoolUrlRe = regexp.MustCompile(
	`<a href="(xylb.jsp\?urltype=tsites.CollegeTeacherList&wbtreeid=1021&st=0&id=[0-9]+[^>]+)><p[^>]+>([^<]+)</p></a>`)

func ParseSchoolList(contents []byte, _ string) engine.ParseResult {
	matches := schoolUrlRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    domain + string(m[1]),
			Parser: engine.NewFuncParser(ParseSchool, config.ParseSchool),
		})
	}
	return result
}
