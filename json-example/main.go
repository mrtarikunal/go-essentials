package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	//`json:"first_name"` json daki karşılık gelen fieldı belrtyrz
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
[
    {
        "first_name": "Clark",
        "last_name": "Kent",
        "hair_color": "black",
        "has_dog": true
    },
    {
        "first_name": "Bruce",
        "last_name": "Wayne",
        "hair_color": "black",
        "has_dog": false
    }
]`

	var unmarshalled []Person
	//slice tanımladık. slice Person type larından olşyr

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	//Unmarshal metodu json decode edyr ve tanımlanan structa atyr.
	//burda myJson nu api den dönen bir json data düşün. onu decode edip unmarshalled listine type da belrttğmz fieldlara atyr
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	// write json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "black"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "     ")
	//developmentta MarshalIndent kullan. ama prod da Marshal kullan
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))
	//MarshalIndent normalde slice of byte olştr, string ile stringe çevirdik
}
