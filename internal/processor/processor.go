package processor

import (
	"fmt"
	"searcher/internal/crawler/spider"
	"searcher/internal/database"
	"searcher/internal/database/files"
	"searcher/internal/index"
)

type Proc struct {
	I *index.Index
	S *spider.Service
	D database.Database
	U string
}

func New() *Proc {
	p := Proc{
		I: index.New(),
		S: spider.New(),
		D: files.New(),
	}

	return &p
}

func (p *Proc) Save() error {

	data := fmt.Sprintf("\tURL:%s\n%s", p.U, p.I.ExportData())

	_, err := p.D.Write([]byte(data))
	if err != nil {
		return err
	}
	return nil
}

func (p *Proc) Load() error {

	data := make([]byte, 1024)

	n, err := p.D.Read(data)
	if err != nil {
		return err
	}

	data = data[:n]

	err = p.I.ImportData(string(data))
	if err != nil {
		return err
	}

	return nil

}

func (p *Proc) FindUrls(t string) {

	urlsID := p.I.DocsID(t)

	for y, url := range urlsID {
		fmt.Printf("[%v]: %s\n", y, p.I.Docs[url].URL)
	}
}
