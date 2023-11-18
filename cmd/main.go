package main

import (
	"fmt"

	"github.com/botiash/catapp/internal/app/service"
	"github.com/botiash/catapp/internal/infrastructure/api"
)

func main() {
	catAPI := api.NewCatAPI()
	catService := service.NewCatService(*catAPI)

	if err := catService.Run(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
