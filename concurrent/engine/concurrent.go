package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	Submit(Request)
	ReadyNotifier
	WorkChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			log.Printf("duplicate url: %s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func(i Item) { e.ItemChan <- i }(item)
		}

		// URL deduplication
		for _, r := range result.Requests {
			if isDuplicate(r.Url) {
				log.Printf("duplicate url: %s", r.Url)
				continue
			}
			e.Scheduler.Submit(r)
		}
	}
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, notifier ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i'm ready
			notifier.WorkerReady(in)

			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
