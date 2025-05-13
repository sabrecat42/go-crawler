package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
	defer file.Close()

	urlToScrape := "https://news.ycombinator.com/item?id=17203903"

	c := colly.NewCollector(colly.AllowedDomains("news.ycombinator.com"))

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "ru-RU")
	})

	c.OnHTML("td.title span.titleline", func(h *colly.HTMLElement) {
		// fmt.Println("Title:", h.Text)
		line := fmt.Sprintf("Title: %s\n", h.Text)
		fmt.Print(line)
		_, err := file.WriteString(line)
		if err != nil {
			log.Printf("Failed to write title: %s", err)
		}
	})

	c.OnHTML("div.comment div.commtext", func(h *colly.HTMLElement) {
		// selection = h.DOM
		// fmt.Println("Comment:", h.Text)
		line := fmt.Sprintf("Comment: %s\n\n", h.Text)
		fmt.Print(line)
		_, err := file.WriteString(line)
		if err != nil {
			log.Printf("Failed to write comment: %s", err)
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error while scrapping: %s\n", err.Error())
	})

	c.Visit(urlToScrape)

}
