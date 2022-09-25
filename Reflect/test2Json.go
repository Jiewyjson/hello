package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int
	Price  int
	Actors []string
}

func main() {
	movie := Movie{"The Matrix", 1999, 10, []string{"Keanu Reeves", "Laurence Fishburne", "Carrie-Anne Moss"}}
	//fmt.Println(movie)

	// Marshal
	jsonstr, err := json.Marshal(movie)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("jsonstr= %s", jsonstr)
	}
	fmt.Println("--------")
	// Unmarshal
	var movie2 Movie
	err = json.Unmarshal(jsonstr, &movie2)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("movie2= %v", movie2)
	}
}
