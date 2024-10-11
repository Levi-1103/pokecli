package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func DisplayLocations(pageURL *string) (*string, *string, error) {

	baseUrl := "https://pokeapi.co/api/v2"
	endPoint := "/location-area"
	fullUrl := baseUrl + endPoint

	if pageURL != nil {
		fullUrl = *pageURL
	}

	res, err := http.Get(fullUrl)
	if err != nil {
		return nil, nil, fmt.Errorf("error couldn't get locations: %s", err)
	}
	defer res.Body.Close()

	var locations LocationAreas

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println(locations.Next)

	for _, locationArea := range locations.Results {
		fmt.Println("-", locationArea.Name)
	}

	next := locations.Next
	prev := locations.Previous

	return next, prev, nil
}
