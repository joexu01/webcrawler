package parser

import (
	"regexp"
	"webcrawler/engine"
	"webcrawler/model"
)

var (
	genderRe = regexp.MustCompile(`<[^>]+>性别：([^<])</[^>]+>`)
	graduatedRe = regexp.MustCompile(`<[^>]+>毕业院校：([^<]+)</[^>]+>`)
	educationRe = regexp.MustCompile(`<[^>]+>学历：([^<]+)</[^>]+>`)
	degreeRe = regexp.MustCompile(`<[^>]+>学位：([^<]+)</[^>]+>`)
)

func ParseTeacherProfile(contents []byte, name string, position string) engine.ParseResult {
	teacher := model.TeacherProfile{
		Name:     name,
		Position: position,
	}
	teacher.Gender = extractString(contents, genderRe)
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
