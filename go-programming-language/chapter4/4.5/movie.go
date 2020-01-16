package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{
		Title: "Casablanca", Year: 1982, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
	},
	{
		Title: "Paul Newman", Year: 1986, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"},
	},
}

func struct2Json() {
	//data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatal("JSON Marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

}

func json2Struct() {
	var titles []struct{ Title string }
	data, _ := json.Marshal(movies)
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatal("JSON Unmarshal failed : %s", err)
	}
	fmt.Printf("%s\n", titles)
}
