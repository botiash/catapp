package service

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/botiash/catapp/internal/app/model"
	"github.com/botiash/catapp/internal/infrastructure/api"
)

type CatService struct {
	catAPI api.CatAPI
}

func NewCatService(catAPI api.CatAPI) *CatService {
	return &CatService{
		catAPI: catAPI,
	}
}

func (s *CatService) Run() error {
	breeds, err := s.catAPI.FetchBreeds()
	if err != nil {
		return err
	}

	groupedBreeds := s.GroupByOrigin(breeds)
	s.SortByLength(groupedBreeds)
	err = s.SaveToFile(groupedBreeds, "out.json")
	if err != nil {
		return err
	}

	fmt.Println("Data processed and saved successfully.")
	return nil
}

func (s *CatService) GroupByOrigin(breeds []model.Cat) map[string][]string {
	groupedBreeds := make(map[string][]string)

	for _, breed := range breeds {
		groupedBreeds[breed.Country] = append(groupedBreeds[breed.Country], breed.Breed)
	}

	return groupedBreeds
}

func (s *CatService) SortByLength(groupedBreeds map[string][]string) {
	for country, breeds := range groupedBreeds {
		sort.Slice(breeds, func(i, j int) bool {
			return len(breeds[i]) < len(breeds[j])
		})
		fmt.Printf("Country: %s, Sorted Breeds: %v\n", country, breeds)
	}
}

func (s *CatService) SaveToFile(data interface{}, filename string) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Data saved to %s\n", filename)
	return nil
}
