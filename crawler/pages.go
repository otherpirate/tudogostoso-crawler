package crawler

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"

	"github.com/gocolly/colly/v2"
)

func FindAllRecipes(input chan string, output chan RecipeUrl, wg *sync.WaitGroup) {
	c := colly.NewCollector()
	matcher, err := regexp.Compile("/receita/")
	if err != nil {
		log.Fatalf("Could compile regexp %+v", err)
	}
	category := ""
	hasContent := false
	c.OnHTML("div[class=col-lg-5]", func(e *colly.HTMLElement) {
		urls := e.ChildAttrs("a", "href")
		for _, url := range urls {
			if !strings.HasSuffix(url, ".html") {
				return
			}
			match := matcher.MatchString(url)
			if !match {
				return
			}
			hasContent = true
			output <- RecipeUrl{
				Category: category,
				Link:     url,
			}
		}
	})
	for link := range input {
		category = strings.ReplaceAll(link, "/categorias/", "")
		page := 1
		base := fmt.Sprintf("https://www.tudogostoso.com.br%s", link)
		for {
			hasContent = false
			url := base
			if page != 1 {
				url = fmt.Sprintf("%s?page=%d", base, page)
			}
			log.Printf("[FindAllRecipes][Visiting] %s", url)
			c.Visit(url)
			if !hasContent {
				break
			}
			page++
		}
	}
	wg.Done()
}
