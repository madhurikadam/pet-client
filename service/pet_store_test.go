package service

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/madhuri/pet-client/model"
	"github.com/madhuri/pet-client/pkg/api"
)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestPetStore_PostPet_Success(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`{"id":1,"type":"cat","price":200.2}`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	t.Setenv("PET_BASE_URL", "http://example.com")
	api := PetStore{Client: &api.API{
		Client: client,
	}}
	body, err := api.GetPet(context.TODO(), fmt.Sprintf("%d", 1))
	assert.NoError(t, err)
	assert.Equal(t, model.Pet{ID: 1, Type: "cat", Price: 200.2}, *body)
}

func TestPetStore_PostPet_Error(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		return &http.Response{
			StatusCode: 404,
			// Send response to be tested
			Body: nil,
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	t.Setenv("PET_BASE_URL", "http://example.com")
	api := PetStore{Client: &api.API{
		Client: client,
	}}
	body, err := api.GetPet(context.TODO(), fmt.Sprintf("%d", 1))
	assert.Error(t, err)
	assert.Nil(t, body)
}

//we can write test cases for other apis as well/
