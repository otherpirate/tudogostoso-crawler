package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/otherpirate/tudogostoso-crawler/crawler"
)

const WORKERS = 1

func main() {
	startTime := time.Now()
	fmt.Println("Starting")
	wg := &sync.WaitGroup{}
	categories := make(chan string)
	crawler.FindAllCategories(categories, wg)

	urls := make(chan crawler.RecipeUrl)
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go crawler.FindAllRecipes(categories, urls, wg)
	}
	recipes := make(chan crawler.Recipe)
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go crawler.ExtractRecipe(urls, recipes, wg)
	}
	crawler.Clean()
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go crawler.Save(recipes, wg)
	}
	fmt.Println("Running")
	wg.Wait()
	fmt.Println("Finished")
	fmt.Printf("Took: %v", time.Since(startTime))
}
