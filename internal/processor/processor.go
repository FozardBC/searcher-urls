package processor

import (
	"encoding/json"
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

	b, err := json.Marshal(p.I)
	if err != nil {
		return err
	}

	_, err = p.D.Write([]byte(b))
	if err != nil {
		return err
	}

	return nil
}

func (p *Proc) Load() error {

	data := make([]byte, 4096)

	n, err := p.D.Read(data)
	if err != nil {
		return fmt.Errorf("can't read data: %w", err)
	}

	err = json.Unmarshal(data[:n], &p.I)
	if err != nil {
		return fmt.Errorf("can't unmarshal data: %w", err)
	}

	return nil

}
