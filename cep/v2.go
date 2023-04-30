package cep

import (
	"fmt"
	"github.com/philipelima/brasilapi-go/utils/request"
)

type CepV2 struct {
	base    string
	request *request.Request
}


func V2() *CepV2 {
	return &CepV2{
		base: "https://brasilapi.com.br/api/cep/v2/%s",
	}
}

func (c *CepV2) Get(cep string) (string, error) {
	
	url := fmt.Sprintf(c.base, cep)

	c.request = request.NewRequest(url, nil)

	c.request.SetHeader("Content-Type", "application/json")

    return c.request.Get()
}
