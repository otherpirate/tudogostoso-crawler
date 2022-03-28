package main

import (
	"sync"

	"github.com/otherpirate/tudogostoso-crawler/crawler"
)

func main() {
	wg := &sync.WaitGroup{}
	categories := make(chan string)
	go crawler.FindAllCategories(categories, wg)

	urls := make(chan crawler.RecipeUrl)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go crawler.FindAllRecipes(categories, urls, wg)
	}
	recipes := make(chan crawler.Recipe)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go crawler.ExtractRecipe(urls, recipes, wg)
	}
	crawler.Clean()
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go crawler.Save(recipes, wg)
	}
	wg.Wait()
}
