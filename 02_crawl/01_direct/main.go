package main

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly"
)

const site = "https://www.jayantha.in"

type Info struct {
	Name  string
	Skill string
}

func main() {
	c := colly.NewCollector()

	messages := []Info{}

	c.OnHTML(".my-60", func(e *colly.HTMLElement) {
		messages = append(messages, Info{
			Name:  e.ChildText(".font-bold"),
			Skill: e.ChildText(".inline"),
		})
	})

	err := c.Visit(site)
	if err != nil {
		panic(err)
	}

	c.Wait()

	bs, err := json.MarshalIndent(messages, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
	fmt.Println("Number of total tweets:", len(messages))
}
