package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/madhuri/pet-client/model"
	"github.com/madhuri/pet-client/pkg/constant"
	"github.com/sirupsen/logrus"
)

type PetStoreClient interface {
	GetPet(ctx context.Context, petID string) (*model.Pet, error)
	PostPet(ctx context.Context, petReq *model.PetPostReq) (*model.PetPostResp, error)
	GetPets(ctx context.Context) ([]*model.Pet, error)
}

func (ps *PetStore) GetPet(ctx context.Context, petID string) (*model.Pet, error) {

	respBody, err := ps.Client.DoStuff(http.MethodGet, fmt.Sprintf(constant.GetPetEndpoint, petID), nil)
	if err != nil {
		logrus.Error("failed to get pet by id", err)
		return nil, err
	}
	resp := &model.Pet{}
	err = json.Unmarshal(respBody, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ps *PetStore) PostPet(ctx context.Context, petReq *model.PetPostReq) (*model.PetPostResp, error) {

	data, err := json.Marshal(petReq)
	if err != nil {
		logrus.Error("failed to get pet by id", err)
	}
	respBody, err := ps.Client.DoStuff(http.MethodPost, constant.PostPetEndpoint, bytes.NewBuffer(data))
	if err != nil {
		logrus.Error("failed to get pet by id", err)
		return nil, err
	}
	resp := &model.PetPostResp{}
	err = json.Unmarshal(respBody, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ps *PetStore) GetPets(ctx context.Context) ([]*model.Pet, error) {
	respBody, err := ps.Client.DoStuff(http.MethodGet, constant.PostPetEndpoint, nil)
	if err != nil {
		logrus.Error("failed to get pet by id", err)
		return nil, err
	}
	resp := []*model.Pet{}
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
