package processor

import (
	"fmt"
	"searcher/internal/index"
	"searcher/pkg/crawler/spider"
)

type Proc struct {
	I *index.Index
	S *spider.Service
}

func New() *Proc {
	p := Proc{
		I: index.New(),
		S: spider.New(),
	}

	return &p
}

func (p *Proc) FindUrls(t string) {

	urlsID := p.I.DocsID(t)

	for y, url := range urlsID {
		fmt.Printf("[%v]: %s\n", y, p.I.Docs[url].URL)
	}
}
