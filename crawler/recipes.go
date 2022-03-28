package crawler

import (
	"fmt"
	"log"
	"strings"
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
				Name:        getName(e.Text),
				Quantity:    getQuantity(e.Text),
				Unit:        getUnit(e.Text),
				IsOptional:  isOptional(e.Text),
				Description: e.Text,
			}
			recipe.Ingredients = append(recipe.Ingredients, ingredient)
		})
		output <- recipe
	})
	for url = range input {
		url := fmt.Sprintf("https://www.tudogostoso.com.br%s", url.Link)
		log.Printf("[ExtractRecipe][Visiting] %s", url)
		c.Visit(url)
	}
	wg.Done()
}

func isOptional(text string) bool {
	return strings.Contains(text, "(opcional)")
}

func getName(text string) string {
	return text
}

func getQuantity(text string) string {
	return text
}

func getUnit(text string) string {
	if strings.Contains(text, "colher") {
		return "Colher"
	} else if strings.Contains(text, "xícara") {
		return "Xícara"
	} else if strings.Contains(text, "pitada") {
		return "Pitada"
	} else if strings.Contains(text, "lata") {
		return "Lata"
	} else if strings.Contains(text, "g ") || strings.Contains(text, "grama ") {
		return "Grama"
	} else if strings.Contains(text, "kg ") || strings.Contains(text, "kilo ") {
		return "Kilograma"
	} else if strings.Contains(text, "ml ") || strings.Contains(text, "mililitro ") {
		return "Mililitro"
	} else if strings.Contains(text, "lt ") || strings.Contains(text, "litro ") {
		return "Litros"
	} else if strings.Contains(text, "caixa ") || strings.Contains(text, "caixinha ") {
		return "Caixa"
	} else {
		return "Unidade"
	}
}
