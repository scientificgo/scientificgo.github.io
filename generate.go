package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func repositories(org string) (data []map[string]interface{}) {
	path := fmt.Sprintf("https://api.github.com/orgs/%v/repos", org)
	response, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Read response data in to memory
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
	}

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal("Error unmarshalling body into JSON. ", err)
	}

	return data
}

func copyFile(path, dest string) {
	log.Printf("Copying %v to %v", path, dest)

	response, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Read response data in to memory
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
	}

	ioutil.WriteFile(dest, body, 0644)
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting working directory. ", err)
	}

	for _, repo := range repositories("scientificgo") {
		name := repo["name"].(string)
		if !strings.Contains(wd, name) {
			path := fmt.Sprintf("https://raw.githubusercontent.com/scientificgo/%v/master/README.md", name)
			dest := fmt.Sprintf("docs/%v.md", name)
			copyFile(path, dest)
		}
	}
}
