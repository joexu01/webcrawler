package view

import (
	"html/template"
	"io"
	"webcrawler/front/model"
)

type SearchResultView struct {
	resultTemplate *template.Template
}

func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{resultTemplate:template.Must(template.ParseFiles(filename))}
}

func (s SearchResultView)Render(w io.Writer, data model.SearchResult) error {
	return s.resultTemplate.Execute(w, data)
}
