package service

import (
	"net/http"

	"github.com/madhuri/pet-client/pkg/api"
)

type PetStore struct {
	Client *api.API
}

func NewPetStore(baseURL string) PetStoreClient {
	return &PetStore{
		Client: &api.API{
			Client:  http.DefaultClient,
			BaseURL: baseURL,
		},
	}
}
