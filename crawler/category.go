package crawler

import (
	"log"
	"sync"

	"github.com/gocolly/colly/v2"
)

func FindAllCategories(output chan string, wg *sync.WaitGroup) {
	wg.Add(1)
	c := colly.NewCollector()
	c.OnHTML("h3", func(e *colly.HTMLElement) {
		link := e.DOM.Find("a")
		if link != nil {
			output <- link.AttrOr("href", "")
		}
	})
	err := c.Visit("https://www.tudogostoso.com.br/categorias")
	if err != nil {
		log.Fatalf("[FindAllCategories][Visit] Error %v", err)
	}
	wg.Done()
}
