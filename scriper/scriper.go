package scriper

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type Scriper struct {
	Colllection *colly.Collector
}

type Result struct {
	Props Props `json:"props"`
}

type Props struct {
	InitialPayload InitialPayload `json:"initialPayload"`
}

type InitialPayload struct {
	Tree Tree `json:"tree"`
}

type Tree struct {
	Items []Item `json:"items"`
}

type Item struct {
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	ContentType ContentType `json:"contentType"`
}

type ContentType string

const (
	Directory   ContentType = "directory"
	File        ContentType = "file"
	SymlinkFile ContentType = "symlink_file"
)

func New() *Scriper {

	// return colly.NewCollector()
	return &Scriper{
		colly.NewCollector(),
	}
}

func (c *Scriper) GetList(selector string) {
	c.Colllection.OnHTML("script[type=\"application/json\"]", func(h *colly.HTMLElement) {

		if strings.Contains(h.Text, "props") {

			var data Result
			err := json.Unmarshal([]byte(h.Text), &data)
			if nil != err {
				panic(err)
			}

			items := data.Props.InitialPayload.Tree.Items

			for _, item := range items {
				itemName := strings.Replace(item.Name, ".gitignore", "", 1)
				if item.ContentType == "file" {
					fmt.Println(itemName)
				}
			}
		}
	})
}

func (c *Scriper) GetIgnorFile(url string) {

	c.Colllection.OnHTML("script[type=\"application/json\"]", func(h *colly.HTMLElement) {

		if strings.Contains(h.Text, "payload") {

			start := strings.Index(h.Text, "rawLines")
			end := strings.Index(h.Text, "stylingDirectives")

			data :=
				strings.Split(h.Text[start:end], "rawLines\":")

			var result string
			removeFirst := strings.Replace(data[1], "\",", "\n", -1)
			result = strings.Replace(removeFirst, "\"", "", -1)

			fmt.Println(result[1 : len(result)-2])
		}
	})
}

func (c *Scriper) OnError() {

	c.Colllection.OnError(func(r *colly.Response, err error) {

		fmt.Println("cound not find language")
	})
}

func (c *Scriper) Visit(url string)  {

	_ = c.Colllection.Visit(url)
}
