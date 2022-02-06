package handlers

import "net/http"

type PetHandler interface {
	CreatePet(w http.ResponseWriter, r *http.Request)
	GetPet(w http.ResponseWriter, r *http.Request)
	GetPets(w http.ResponseWriter, r *http.Request)
}
