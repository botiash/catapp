package api

import (
	"encoding/json"
	"fmt"
	"io"
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
	response, err := http.Get(catAPIURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var breedsResponse struct {
		Data []model.Cat `json:"data"`
	}

	err = json.Unmarshal(body, &breedsResponse)
	if err != nil {
		return nil, err
	}

	return breedsResponse.Data, nil
}
