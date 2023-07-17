package cep

import (
	"fmt"
	"log"

	"github.com/philipelima/brasilapi-go/utils/parse"
	"github.com/philipelima/brasilapi-go/utils/request"
)

type CepV2 struct {
	base    string
	request *request.Request
}

type CepV2Response struct {
	Cep   string  `json:"cep"`
	State string  `json:"state"`
	City  string  `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street  string `json:"street"`
	Service string `json:"service"`
}


func V2() *CepV2 {
	return &CepV2{
		base: "https://brasilapi.com.br/api/cep/v2/%s",
	}
}

func (c *CepV2) Get(cep string) (*CepV2Response, error) {
	
	var cepData *CepV2Response


	url := fmt.Sprintf(c.base, cep)


	c.request = request.NewRequest(url, nil)

	c.request.SetHeader("Content-Type", "application/json")

	response, err := c.request.Get()

	if err != nil {
		log.Fatal(err)
	}

    
	parse.StrToStruct(response, &cepData)

    return cepData, nil 
}
