package exchanges

import (
	"fmt"
	"log"

	"github.com/philipelima/brasilapi-go/utils/parse"
	"github.com/philipelima/brasilapi-go/utils/request"
)

type Exchange struct {
	Bairro string
	Cep string
	Cnpj string
	Codigo_cvm string
	Complemento string
	Data_inicio_situacao string
	Data_patrimonio_liquido string
	Data_registro string
	Email string
	Logradouro string
	Municipio string
	Nome_social string
	Nome_comercial string
	Pais string
	Status string
	Telefone string
	Type string
	Uf string
	Valor_patrimonio_liquido string
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