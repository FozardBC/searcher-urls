package index

import (
	"fmt"
	"searcher/internal/crawler"
	"strconv"
	"strings"
)

type Index struct {
	Words map[string][]int
	Docs  []crawler.Document
}

func New() *Index {
	return &Index{Words: make(map[string][]int)}
}

func (i *Index) AddWord(word string, docId int) {
	i.Words[word] = append(i.Words[word], docId)
}

func (i *Index) DocsID(word string) []int {
	return i.Words[word]
}

func (i *Index) ImportData(data string) error {
	//добавить различный парсинг в зависимости от источника датабазы

	err := i.pFile(data)
	if err != nil {
		return fmt.Errorf("can't parse file: %w", err)
	}

	return nil
}

// функция для парсинга текста если он из файла
func (i *Index) pFile(data string) error {
	docs := strings.Split(data, "\tURL:") // тут первая строка всегда окажется ссылкой
	docs = docs[1:]

	for _, d := range docs {
		parts := strings.Split(d, "___\n")
		if len(parts) != 2 {
			return fmt.Errorf("wrong data format")
		}

		w := strings.Split(parts[0], "\n")
		w = w[1:]

		err := i.impWords(w)
		if err != nil {
			return fmt.Errorf("can't import words: %w", err)
		}

		err = i.impDocs(strings.Split(parts[1], "\n"))
		if err != nil {
			return fmt.Errorf("can't import docs: %w", err)
		}

	}

	return nil
}

func (i *Index) impDocs(docs []string) error {
	for _, d := range docs {
		if d == "" || d == " " {
			continue
		}

		line := strings.Split(d, "]")
		if len(line) != 2 {
			return fmt.Errorf("wrong data format: %s", d)
		}

		line[0] = strings.Trim(line[0], "[]")
		id, err := strconv.Atoi(line[0])
		if err != nil {
			return err
		}

		i.Docs = append(i.Docs, crawler.Document{ID: id, URL: line[1]})
	}
	return nil
}

func (i *Index) impWords(words []string) error {
	for _, w := range words {
		if w == "" || w == " " {
			continue
		}

		line := strings.Split(w, ":")
		if len(line) != 2 {
			return fmt.Errorf("wrong data format: %s", w)
		}

		line[1] = strings.Trim(line[1], "[]") // конвертируем айдишники в []int
		for _, id := range strings.Split(line[1], ",") {
			id, err := strconv.Atoi(id)
			if err != nil {
				return err
			}

			i.AddWord(line[0], id)
		}

	}
	return nil
}
