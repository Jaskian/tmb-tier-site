package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/antchfx/xmlquery"
)

func main() {
	b, err := ioutil.ReadFile("data/ids.csv")
	mustNot(err)

	itemIds := strings.Split(string(b), "\n")

	messages := make(chan item)

	for _, itemId := range itemIds {
		go callWowhead(itemId, messages)
	}

	var sb strings.Builder
	for range itemIds {
		r := <-messages
		sb.WriteString(fmt.Sprintf("%d : %d,\n", r.itemid, r.ilvl))
	}

	trimmed := strings.Trim(sb.String(), "\n")
	ioutil.WriteFile("data/itemlevels.txt", []byte(trimmed), fs.ModePerm)
}

func callWowhead(itemId string, ch chan<- item) {
	i, err := strconv.Atoi(itemId)
	mustNot(err)
	url := fmt.Sprintf("https://www.wowhead.com/wotlk/item=%d&xml", i)

	doc, err := xmlquery.LoadURL(url)
	mustNot(err)

	node := xmlquery.FindOne(doc, "//wowhead/item/level")
	ilvl, err := strconv.Atoi(node.InnerText())
	mustNot(err)

	ch <- item{i, ilvl}
}

type item struct {
	itemid int
	ilvl   int
}

func mustNot(e error) {
	if e != nil {
		panic(e)
	}
}
