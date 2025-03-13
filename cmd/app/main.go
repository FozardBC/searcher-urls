package main

import (
	"fmt"
	"searcher/internal/crawler"
	"searcher/internal/processor"
	"strings"
)

func main() {

	p := processor.New()

	// urls := []string{"https://google.com", "https://vk.com"}
	// p.D.Clear()

	// Scan(p, urls)

	// p.Save()

	err := p.Load()
	if err != nil {
		fmt.Print(err)
	}
}

func Scan(p *processor.Proc, urls []string) {

	var err error

	var data []crawler.Document

	for _, url := range urls {
		data, err = p.S.Scan(url, 2) // получаем все отсканированные данные ID, URL, Body
		if err != nil {
			fmt.Print(err)
		}

		p.I.Docs = append(p.I.Docs, data...)
	}

	splitFunc := func(r rune) bool {
		return r == '/' || r == '?' || r == '&' || r == '#' || r == ' ' || r == '.' || r == ':' || r == '{' || r == '}'
	}

	for _, doc := range p.I.Docs {
		words := strings.FieldsFunc(doc.URL, splitFunc)

		for _, w := range words {

			p.I.AddWord(w, doc.ID)
		}
	}

}
