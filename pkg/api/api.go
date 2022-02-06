package api

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type API struct {
	Client  *http.Client
	BaseURL string
}

func (api *API) DoStuff(method string, path string, body io.Reader) ([]byte, error) {

	var (
		resp *http.Response
		err  error
	)
	apiPath := fmt.Sprintf("%s%s", api.BaseURL, path)
	switch method {
	case http.MethodGet:
		resp, err = api.Client.Get(apiPath)
	case http.MethodPost:
		if body == nil {
			return nil, errors.New("request body is empty")
		}
		resp, err = api.Client.Post(apiPath, "application/json", body)
	}
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with error %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	Respbody, err := ioutil.ReadAll(resp.Body)
	// handling error and doing stuff with body that needs to be unit tested
	return Respbody, err
}
