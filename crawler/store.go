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
	os.MkdirAll("./store/", 0775)
}

func Save(input chan Recipe, wg *sync.WaitGroup) {
	for recipe := range input {
		file, err := json.MarshalIndent(recipe, "", " ")
		if err != nil {
			log.Printf("[Save][Parser] Could not parser %+v | Err: %+v", recipe, err)
			continue
		}
		fileName := strings.ReplaceAll(recipe.Link, "/receita/", "")
		fileName = strings.ReplaceAll(fileName, ".html", ".json")
		err = ioutil.WriteFile("./store/"+fileName, file, 0644)
		if err != nil {
			log.Printf("[Save][Parser] Could not save file %s - %+v | Err: %+v", fileName, recipe, err)
			continue
		}
		log.Printf("[Save][Store] Saved %s", fileName)
	}
	wg.Done()
}
