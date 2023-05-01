package banks

import (
	"fmt"
	"github.com/philipelima/brasilapi-go/utils/parse"
	"github.com/philipelima/brasilapi-go/utils/request"
)

type BanksV1 struct {
	base string
	request *request.Request
}

type Bank struct {
	Isbp string `json:"isbp"`
	Name string `json:"name"`
	Code int64  `json:"code"`
	FullName string `json:"fullname"`
}


func V1() *BanksV1 {
	return &BanksV1{
		base: "https://brasilapi.com.br/api/banks/v1",
	}
}

func (b *BanksV1) All() (*[]Bank, error) {

	var banksResult *[]Bank

	b.request = request.NewRequest(b.base, nil)

	b.request.SetHeader("Content-Type", "application/json")

	banks, err := b.request.Get()

	if err != nil {
		return nil, err
	}

	parse.StrToStruct(banks, &banksResult)

	return banksResult, nil
}


func (b *BanksV1) Get(code string) (*Bank, error) {
	
	var result *Bank

	url := fmt.Sprintf("%s/%s", b.base, code)

	b.request = request.NewRequest(url, nil)

	b.request.SetHeader("Content-Type", "application/json")

	bank, err := b.request.Get()

	if err != nil {
		return nil, err
	}

	parse.StrToStruct(bank, &result)

	return result, nil
}

