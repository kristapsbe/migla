package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Entries struct {
	Entries []Entry `json:"entries"`
}

type Entry struct {
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Signature string    `json:"signature"`
	Image     string    `json:"image"`
	TimeFrom  time.Time `json:"time_from"`
	TimeTo    time.Time `json:"time_to"`
	Weight    int       `json:"weight"`
	Opacity   float32   `json:"opacity"`
}

func main() {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)

	var entries Entries

	json.Unmarshal(byteValue, &entries)

	for i := 0; i < len(entries.Entries); i++ {
		fmt.Println("Title: " + entries.Entries[i].Title)
		fmt.Println("Weight: " + strconv.Itoa(entries.Entries[i].Weight))
		fmt.Println("TimeFrom: " + entries.Entries[i].TimeFrom.String())
	}
}
