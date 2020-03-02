package engine

type ParserFunc func(contents []byte, url string) ParseResult

type Request struct {
	Url        string
	ParserFunc func(contents []byte, url string) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Payload interface{}
}

type NilParser struct{}

func (NilParser) Parse(
	_ []byte, _ string) ParseResult {
	return ParseResult{}
}
