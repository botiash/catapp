package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/botiash/catapp/internal/app/model"
)

const catAPIURL = "https://catfact.ninja/breeds"

type CatAPI struct {
}

func NewCatAPI() *CatAPI {
	return &CatAPI{}
}

func (api *CatAPI) FetchBreeds() ([]model.Cat, error) {
	var allBreeds []model.Cat
	page := 1
	for {
		response, err := http.Get(fmt.Sprintf("%s?page=%d", catAPIURL, page))
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()
		if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("API request failed with status code: %d", response.StatusCode)
		}
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		var breedsResponse struct {
			Data  []model.Cat `json:"data"`
			Total int         `json:"total"`
		}
		err = json.Unmarshal(body, &breedsResponse)
		if err != nil {
			return nil, err
		}
		allBreeds = append(allBreeds, breedsResponse.Data...)
		if len(allBreeds) >= breedsResponse.Total {
			break
		}
		page++
	}

	return allBreeds, nil
}
