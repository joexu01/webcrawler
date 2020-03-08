package worker

import "webcrawler/concurrent/engine"

type CrawlService struct {}

func (CrawlService) Process(req Request, result *ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	engParseResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	*result = SerializeResult(engParseResult)
	return nil
}
