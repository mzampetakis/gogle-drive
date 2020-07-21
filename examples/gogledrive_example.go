package main

import (
	"fmt"
	"log"

	gogledrive "github.com/mzampetakis/gogle-drive"
)

func main() {
	searchAllFiles()
	searchImagesFiles()
}

func searchAllFiles() {
	fmt.Println("Your Google's Drive Files:")
	gdrive, err := gogledrive.New("../credentials.json")
	if err != nil {
		log.Fatal(err)
	}
	filterCriteria := gogledrive.ListFilter{}
	assets, err := gdrive.SearchFiles(filterCriteria)
	if err != nil {
		log.Fatal(err)
	}
	//Prints all your files in your google drive
	for name, fileID := range assets {
		fmt.Printf("Name: %s, ID: %s\n", name, fileID)
	}
}

func searchImagesFiles() {
	fmt.Println("Your Google's Drive Images:")
	gdrive, err := gogledrive.New("../credentials.json")
	if err != nil {
		log.Fatal(err)
	}
	mimeTypeCriteris := "image/*"
	filterCriteria := gogledrive.ListFilter{
		MimeType: &mimeTypeCriteris,
	}
	assets, err := gdrive.SearchFiles(filterCriteria)
	if err != nil {
		log.Fatal(err)
	}
	//Prints all your image files in your google drive
	for name, fileID := range assets {
		fmt.Printf("Name: %s, ID: %s\n", name, fileID)
	}
}
