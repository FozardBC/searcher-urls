package app

import (
	"fmt"
	"searcher/internal/index"
	"searcher/internal/processor"
	"strings"
)

func Start() {

	p := processor.New()

	//urls := []string{"https://google.com", "https://vk.com"}
	//Scan(p, urls)

	err := p.Load()
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print("scaned")

}

func Scan(p *processor.Proc, urls []string) {

	var err error

	p.D.Clear()

	for _, url := range urls {
		p.I = index.New()
		p.U = url

		p.I.Docs, err = p.S.Scan(url, 2) // получаем все отсканированные данные ID, URL, Body
		if err != nil {
			fmt.Print(err)
		}

		splitFunc := func(r rune) bool {
			return r == '/' || r == '?' || r == '&' || r == '#' || r == ' ' || r == '.' || r == ':'
		}

		for _, doc := range p.I.Docs {
			words := strings.FieldsFunc(doc.URL, splitFunc)

			for _, w := range words {

				p.I.AddWord(w, doc.ID)
			}
		}
		p.Save()
	}
}
