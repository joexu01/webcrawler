package engine

import (
	"log"
	"webcrawler/concurrent/fetcher"
)

func Worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching url: %s", r.Url)
	body, err := fetcher.Fetcher(r.Url)
	if err != nil {
		log.Printf("Fetcher: an error occured when fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parser.Parse(body, r.Url), nil
}
