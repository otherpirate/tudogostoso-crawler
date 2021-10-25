package crawler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

func Clean() {
	os.RemoveAll("./store/")
}

func Save(input chan Recipe, wg *sync.WaitGroup) {
	for recipe := range input {
		file, err := json.MarshalIndent(recipe, "", " ")
		if err != nil {
			log.Printf("[Save][Parser] Could not parser %+v", recipe)
			continue
		}
		fileName := strings.ReplaceAll(recipe.Link, "/receita/", "")
		fileName = strings.ReplaceAll(fileName, ".html", ".json")
		err = ioutil.WriteFile("./store/"+fileName, file, 0644)
		if err != nil {
			log.Printf("[Save][Parser] Could not save file %s - %+v", fileName, recipe)
			continue
		}
	}
	wg.Done()
}
