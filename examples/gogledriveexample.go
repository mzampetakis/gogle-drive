package main

import (
	"fmt"
	"log"

	gogledrive "github.com/mzampetakis/gogle-drive"
)

func main() {
	gdrive, err := gogledrive.New("../credentials.json")
	if err != nil {
		log.Fatal(err)
	}
	filterCriteria := gogledrive.ListFilter{}
	assets, err := gdrive.SearchFiles(filterCriteria)
	if err != nil {
		log.Fatal(err)
	}
	for name, fileID := range assets {
		fmt.Printf("Name: %s, ID: %s\n", name, fileID)
	}
}
