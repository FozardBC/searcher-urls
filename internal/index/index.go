package index

import "searcher/pkg/crawler"

type Index struct {
	words map[string][]int
	Docs  []crawler.Document
}

func New() *Index {
	return &Index{words: make(map[string][]int)}
}

func (i *Index) AddWord(word string, docId int) {
	i.words[word] = append(i.words[word], docId)
}

func (i *Index) DocsID(word string) []int {
	return i.words[word]
}
