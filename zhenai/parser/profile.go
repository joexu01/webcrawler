package parser

import (
	"regexp"
	"strconv"
	"webcrawler/engine"
	"webcrawler/model"
)

var ageRe = regexp.MustCompile(
	`<div.+purple[^>]+>([\d]+)岁</div>`)
//var nameRe = regexp.MustCompile(
//	`<h1 data-v-5b109fc3="" class="nickName">([^<]+)</h1>`)
var heightRe = regexp.MustCompile(
	`<div.+purple[^>]+>([\d]+)cm</div>`)
var locationRe = regexp.MustCompile(
	`<div.+purple[^>]+>工作地:([^<]+)</div>`)
var incomeRe = regexp.MustCompile(
	`<div.+purple[^>]+>月收入:([^<]+)</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	profile.Name = name
	profile.Income = extractString(contents, incomeRe)
	profile.Location = extractString(contents, locationRe)

	result := engine.ParseResult{
		Requests: nil,
		Items:    []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
