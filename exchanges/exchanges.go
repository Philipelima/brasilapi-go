package exchanges

import (
	"fmt"
	"log"

	"github.com/philipelima/brasilapi-go/utils/parse"
	"github.com/philipelima/brasilapi-go/utils/request"
)

type Exchange struct {
	Bairro string `json:"bairro"`
	Cep string	  `json:"cep"`
	Cnpj string	  `json:"cnpj"`
	Codigo_cvm string `json:"codigo_cvm"`
	Complemento string `json:"complemento"`
	Data_inicio_situacao string `json:"data_inicio_situacao"`
	Data_patrimonio_liquido string `json:"data_patrimonio_liquido"`
	Data_registro string `json:"data_registro"`
	Email string  `json:"email"`
	Logradouro string `json:"logradouro"`
	Municipio string `json:"municipio"`
	Nome_social string `json:"nome_social"`
	Nome_comercial string `json:"nome_comercial"`
	Pais string `json:"pais"`
	Status string `json:"status"`
	Telefone string `json:"telefone"`
	Type string `json:"type"`
	Uf string `jso:"uf"`
	Valor_patrimonio_liquido string `json:"valor_patrimonio_liquido"`
}

type ExchangesV1 struct {
	base string 
}



func V1() *ExchangesV1 {
	return &ExchangesV1{
		base: "https://brasilapi.com.br/api/cvm/corretoras/v1",
	}
}


func (e *ExchangesV1) All() (*[]Exchange, error){

	var result *[]Exchange

	request := request.NewRequest(e.base, nil)

	request.SetHeader("Content-Type", "application/json")

	exchanges, err := request.Get()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	parse.StrToStruct(exchanges, &result)

	return result, nil
}



func (e *ExchangesV1) Get(cnpj string) (*Exchange, error){

	var result *Exchange

	url := fmt.Sprintf("%s/%s", e.base, cnpj)

	request := request.NewRequest(url, nil)

	request.SetHeader("Content-Type", "application/json")

	exchanges, err := request.Get()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	parse.StrToStruct(exchanges, &result)

	return result, nil
}