package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/madhuri/pet-client/model"
	"github.com/madhuri/pet-client/service"
)

//BankAndAccountHandler is the class implementation for CompositeIface Interface
type PetHandlerImpl struct {
	PetService service.PetStoreClient
}

func (h *PetHandlerImpl) CreatePet(w http.ResponseWriter, r *http.Request) {

	postReq := model.PetPostReq{}
	bodyBytes, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		responseController(w, http.StatusInternalServerError, readErr)
		return
	}

	strBufferValue := string(bodyBytes)
	err := json.Unmarshal([]byte(strBufferValue), &postReq)
	if err != nil {
		responseController(w, http.StatusInternalServerError, err)
		return
	}

	resp, err := h.PetService.PostPet(r.Context(), &postReq)
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseController(w, http.StatusOK, resp)
}

func (h *PetHandlerImpl) GetPet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	petID := params["id"]

	resp, err := h.PetService.GetPet(r.Context(), petID)
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseController(w, http.StatusOK, resp)
}

func (h *PetHandlerImpl) GetPets(w http.ResponseWriter, r *http.Request) {
	resp, err := h.PetService.GetPets(r.Context())
	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseController(w, http.StatusOK, resp)
}

func responseController(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func NewPetHandel(petStore service.PetStoreClient) PetHandler {
	return &PetHandlerImpl{
		PetService: petStore,
	}
}
