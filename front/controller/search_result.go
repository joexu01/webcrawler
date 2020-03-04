package controller

import (
	"context"
	"github.com/olivere/elastic/v7"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"webcrawler/engine"
	"webcrawler/front/model"
	"webcrawler/front/view"
)

type SearchResultHandler struct {
	v      view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(templateName string) SearchResultHandler{
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	
	return SearchResultHandler{
		v:      view.CreateSearchResultView(templateName),
		client: client,
	}
}

// http://localhost:9200/search?q=key_word&from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.FormValue("q"))
	from, err := strconv.Atoi(r.FormValue("from"))
	if err != nil {
		from = 0
	}

	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = h.v.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := h.client.Search("teacher").
		Query(elastic.NewQueryStringQuery(q)).From(from).
		Do(context.Background())
	if err != nil {
		return result, err
	}
	result.Query = q
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}
