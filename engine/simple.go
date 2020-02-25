package engine

import (
	"log"
	"webcrawler/fetcher"
)

type SimpleEngine struct {}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item: %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching url: %s", r.Url)
	body, err := fetcher.Fetcher(r.Url)
	if err != nil {
		log.Printf("Fetcher: an error occured when fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
