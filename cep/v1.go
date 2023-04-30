package cep

import (
	"encoding/json"
	"fmt"

	"github.com/philipelima/brasilapi-go/utils/request"
)

type CepV1 struct {
	base    string
	request *request.Request
}


type CepV1Response struct {
	Cep   string  `json:"cep"`
	State string  `json:"state"`
	City  string  `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street  string `json:"street"`
	Service string `json:"service"`
}


func V1() *CepV1 {
	return &CepV1{
		base: "https://brasilapi.com.br/api/cep/v1/%s",
	}
}

func (c *CepV1) Get(cep string) (*CepV1Response, error) {
	
	url := fmt.Sprintf(c.base, cep)

	c.request = request.NewRequest(url, nil)

	c.request.SetHeader("Content-Type", "application/json")

	cep, err := c.request.Get()

	if err != nil {
		return nil, err
	}

    return newCepV1Response(cep)
}


func newCepV1Response(response string) (*CepV1Response, error) {
	
	var cepResponse CepV1Response

	err := json.Unmarshal([]byte(response), &cepResponse)

	if err != nil {
		return nil, err
	}

	return &cepResponse, nil
}


func (cr *CepV1Response) GetCep() string {
	return cr.Cep
}

func (cr *CepV1Response) GetStreet() string {
	return cr.Street
}

func (cr *CepV1Response) GetNeighborhood() string {
	return cr.Neighborhood
}


func (cr *CepV1Response) GetCity() string {
	return cr.City
}


func (cr *CepV1Response) GetState() string {
	return cr.State
}

func (cr *CepV1Response) GetService() string {
	return cr.Service
}