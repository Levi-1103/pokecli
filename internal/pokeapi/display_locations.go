package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Levi-1103/pokecli/internal/pokecache"
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

func DisplayLocations(pageURL *string, cache pokecache.Cache) (*string, *string, error) {

	baseUrl := "https://pokeapi.co/api/v2"
	endPoint := "/location-area"
	fullUrl := baseUrl + endPoint

	if pageURL != nil {
		fullUrl = *pageURL
	}

	var locations LocationAreas

	cachedData, ok := cache.Get(fullUrl)

	if ok {
		var locations LocationAreas
		err := json.Unmarshal(cachedData, &locations)
		if err != nil {
			return nil, nil, err
		}
		next := locations.Next
		prev := locations.Previous

		return next, prev, nil

	}

	res, err := http.Get(fullUrl)
	if err != nil {
		return nil, nil, fmt.Errorf("error couldn't get locations: %s", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	err = json.Unmarshal(data, &locations)
	if err != nil {
		return nil, nil, err
	}

	for _, locationArea := range locations.Results {
		fmt.Println("-", locationArea.Name)
	}

	next := locations.Next
	prev := locations.Previous

	cache.Add(fullUrl, data)

	return next, prev, nil
}
