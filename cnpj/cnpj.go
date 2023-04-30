package cnpj

import (
	"encoding/json"
	"fmt"
	"github.com/philipelima/brasilapi-go/utils/request"
)


type CnpjV1 struct {
	base    string 
	request *request.Request
}


type Cnae struct {

	Codigo     int64 `json:"codigo"`
	Descricao string `json:"descricao"`
}

type Qsa struct {

	Indentificador_de_socio    int64 `json:"identificador_de_socio"`
	Nome_socio 				  string `json:"nome_socio"`
	Cnpj_cpf_do_socio 		  string `json:"cnpj_cpf_do_socio"`
	Codigo_qualificacao_socio  int64 `json:"codigo_qualificacao_socio"`
	Percentual_capital_social  int64 `json:"percentual_capital_social"`
	Data_entrada_sociedade    string `json:"data_entrada_sociedade"`
	Cpf_representante_legal  *string `json:"cpf_representante_legal"`
	Nome_representante_legal *string `json:"nome_representante_legal"`
	Codigo_qualificacao_representante_legal *int64 `json:"codigo_qualificacao_representante_legal"`
}

type Cnpj struct {

	Cnpj string	`json:"cnpj"`
	Identificador_matriz_filial   int64 `json:"identificador_matriz_filial"`
	Descricao_matriz_filial   string `json:"descricao_matriz_filial"`
	Razao_social string  `json:"razao_social"`
	Nome_fantasia string `json:"nome_fantasia"`
	Situacao_cadastral int64 `json:"situacao_cadastral"`
	Descricao_situacao_cadastral string `json:"descricao_situacao_cadastral"`
	Data_situacao_cadastral  string `json:"data_situacao_cadastral"`
	Motivo_situacao_cadastral  int32 `json:"motivo_situacao_cadastral"`
	Nome_cidade_exterior  *string `json:"nome_cidade_exterior"`
	Codigo_natureza_juridica int64 `json:"codigo_natureza_juridica"`
	Data_inicio_atividade  string `json:"data_inicio_atividade"`
	Cnae_fiscal   int64 `json:"cnae_fiscal"`
	Cnae_fiscal_descricao  string `json:"cnae_fiscal_descricao"`
	Descricao_tipo_logradouro string  `json:"descricao_tipo_logradouro"`
	Logradouro  string `json:"logradouro"`
	Numero  string `json:"numero"`
	Complemento string `json:"complemento"`
	Bairro  string `json:"bairro"`
	Cep string `json:"cep"`
	Uf  string `json:"uf"`
	Codigo_municipio int64 `json:"codigo_municipio"`
	Municipio  string `json:"municipio"`
	Ddd_telefone_1 string `json:"ddd_telefone_1"`
	Ddd_telefone_2 *string `json:"ddd_telefone_2"`
	Ddd_fax  *string	`json:"ddd_fax"` 
	Qualificacao_do_responsavel int64 `json:"qualificacao_do_responsavel"`
	Capital_social  int64 `json:"capital_social"`
	Porte  string `json:"porte"`
	Descricao_porte  string `json:"descricao_porte"`
	Opcao_pelo_simples bool `json:"opcao_pelo_simples"`
	Data_opcao_pelo_simples *string `json:"data_opcao_pelo_simples"`
	Data_exclusao_do_simples *string `json:"data_exclusao_do_simples"`
	Opcao_pelo_mei  bool `json:"opcao_pelo_mei"`
	Situacao_especial *string `json:"situacao_especial"` 
	Data_situacao_especial *string `json:"data_situacao_especial"`

	Cnae_secudarios []Cnae
	Qsa  []Qsa
}


func V1() *CnpjV1 {
	return &CnpjV1{
		base: "https://brasilapi.com.br/api/cnpj/v1/%s",	
	}
}

func (c *CnpjV1) Get(cnpj string) (*Cnpj, error) {

	url := fmt.Sprintf(c.base, cnpj)

	c.request = request.NewRequest(url, nil)

	cnpj, err := c.request.Get()

	if err != nil {
		return nil, err
	}

	return cnpjResponse(cnpj)
}


func cnpjResponse(response string) (*Cnpj, error) {
 
	var cnpj Cnpj

	err := json.Unmarshal([]byte(response), &cnpj)

	if err != nil {
		return nil, err
	}

	return &cnpj, nil
}