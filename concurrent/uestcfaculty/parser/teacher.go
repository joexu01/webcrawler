package parser

import (
	"regexp"
	"strings"
	"webcrawler/config"
	"webcrawler/engine"
	"webcrawler/model"
)

var (
	genderRe      = regexp.MustCompile(`<[^>]+>性别：([^<])</[^>]+>`)
	graduatedRe   = regexp.MustCompile(`<[^>]+>毕业院校：([^<]+)</[^>]+>`)
	educationRe   = regexp.MustCompile(`<[^>]+>学历：([^<]+)</[^>]+>`)
	degreeRe      = regexp.MustCompile(`<[^>]+>学位：([^<]+)</[^>]+>`)
	isIncumbentRe = regexp.MustCompile(`<[^>]+>在职信息：([^<]+)</[^>]+>`)
	schoolRe      = regexp.MustCompile(`<[^>]+>所在单位：([^<]+)</[^>]+>`)
	disciplineRe  = regexp.MustCompile(`<[^>]+>学科：([^<]+)</[^>]+>`)
	locationRe    = regexp.MustCompile(`<[^>]+>办公地点：([^<]+)</[^>]+>`)
	emailRe       = regexp.MustCompile(`<[^>]+>电子邮箱：([a-zA-Z0-9.]+)(@[a-zA-Z0-9.]+\.[a-zA-Z0-9.]+)</[^>]+>`)
	IdRe          = regexp.MustCompile(`uestc.edu.cn/([^/]+)/`)
)

func ParseTeacherProfile(contents []byte, name string, position string, url string) engine.ParseResult {
	teacher := model.TeacherProfile{
		Name:          name,
		Position:      position,
		Gender:        extractString(contents, genderRe),
		GraduatedFrom: extractString(contents, graduatedRe),
		Education:     extractString(contents, educationRe),
		Degree:        extractString(contents, degreeRe),
		IsIncumbent:   extractString(contents, isIncumbentRe),
		School:        extractString(contents, schoolRe),
		Discipline:    extractString(contents, disciplineRe),
		Location:      extractString(contents, locationRe),
		Email:         extractString(contents, emailRe),
		PersonalUrl:   url,
	}

	result := engine.ParseResult{
		Requests: nil,
		Items: []engine.Item{
			{
				Url:     url,
				Id:      getTeacherId(url, IdRe),
				Payload: teacher,
			},
		},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		result := string(match[1])
		result = strings.ReplaceAll(result, `<br>`, "")
		return result
	}
	return ""
}

// this function is deprecated!
//func TeacherParser(name string, position string) engine.ParserFunc {
//	return func(c []byte, url string) engine.ParseResult {
//		return ParseTeacherProfile(c, name, position, url)
//	}
//}

func getTeacherId(url string, idRe *regexp.Regexp) string {
	id := idRe.FindSubmatch([]byte(url))
	return string(id[1])
}

type ProfileParser struct {
	Name     string
	Position string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseTeacherProfile(contents, p.Name, p.Position, url)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ParseTeacherProfile, p.Name
}

func NewProfileParser(name string, position string) *ProfileParser {
	return &ProfileParser{
		Name:     name,
		Position: position,
	}
}
