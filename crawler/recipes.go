package crawler

import (
	"fmt"
	"sync"

	"github.com/gocolly/colly/v2"
)

func ExtractRecipe(input chan RecipeUrl, output chan Recipe, wg *sync.WaitGroup) {
	var url RecipeUrl
	c := colly.NewCollector()
	c.OnHTML("div[id=info-user]", func(e *colly.HTMLElement) {
		recipe := Recipe{
			Category:    url.Category,
			Link:        url.Link,
			Ingredients: []Ingredient{},
		}
		e.ForEach("li", func(i int, e *colly.HTMLElement) {
			ingredient := Ingredient{
				Name: e.Text,
			}
			recipe.Ingredients = append(recipe.Ingredients, ingredient)
		})
		output <- recipe
	})
	for url = range input {
		c.Visit(fmt.Sprintf("https://www.tudogostoso.com.br%s", url.Link))
	}
	wg.Done()
}
