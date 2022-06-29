package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type MovieRating struct {
	Title     string  `json:"title"`
	Comment   string  `json:"comment"`
	Rating    float32 `json:"rating"`
	DateAdded string  `json:"date_added"`
}

func GetRatings(filename string) (ratings []MovieRating, filesize int64, err error) {
	fileInfo, err := os.Stat(filename)
	fileSize := fileInfo.Size()

	if err != nil {
		return nil, 0, err
	}

	rawData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fileSize, errors.New("Error reading file")
	}
	var movieRatings []MovieRating
	json.Unmarshal(rawData, &movieRatings)
	return movieRatings, fileSize, nil
}

func UpdateJsonFile(filename string, data []MovieRating) {
	if len(data) == 0 {
		log.Fatal("The data you've provided is empty. Exiting because if we continue, your existing ratings will be overwritten")
	}
	CopyFile(filename, filename+".bak")
	d, _ := json.MarshalIndent(data, "", "  ")
	err := ioutil.WriteFile(filename, d, 0644)

	if err != nil {
		log.Fatal("Error writing to file", err.Error())
	}
}

func GetNewRating() MovieRating {
	dt := time.Now()
	var title, comment, stringRating, dateAdded string
	dateAdded = dt.Format("02 01 2006") // dd/mm/yyyy

	GetInput("Movie Title: ", &title)
	GetInput("Comment: ", &comment)
	// Could just use Scanln for rating but this looks cleaner
	GetInput("Rating: ", &stringRating)

	floatRating, err := strconv.ParseFloat(stringRating, 32)
	if err != nil {
		fmt.Println("You can only enter ints or floats for a rating. Your rating has been set to 0.")
	}

	title = strings.Trim(title, " ")
	comment = strings.Trim(comment, " ")

	if floatRating > 10 {
		floatRating = 10
	}

	return MovieRating{Title: title, Comment: comment, Rating: float32(floatRating), DateAdded: dateAdded}
}
