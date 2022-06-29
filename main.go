package main

import (
	"log"
	"path/filepath"
)

func main() {
	inputFile := "movieRatings.json"
	outputFile := "movieRatings.json"
	data, fileSize, err := GetRatings(inputFile)

	if err != nil {
		log.Fatal(err)
	}

	if len(data) == 0 && fileSize > 0 {
		log.Fatal("There's an error in the json file:", inputFile)
	}

	newRating := GetNewRating()

	var newData []MovieRating
	newData = append(data, newRating)
	path, err := GetCurrentPath()
	if err != nil {
		log.Fatal(err)
	}

	UpdateJsonFile(filepath.Join(path, outputFile), newData)
	// UpdateJsonFile("movieRatings.json", newData)
}
