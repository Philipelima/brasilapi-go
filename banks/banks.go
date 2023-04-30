package banks

import (
	"encoding/json"
	"fmt"
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

	b.request = request.NewRequest(b.base, nil)

	b.request.SetHeader("Content-Type", "application/json")

	banks, err := b.request.Get()

	if err != nil {
		return nil, err
	}

	return banksResponse(banks)
}


func (b *BanksV1) Get(code string) (*Bank, error) {

	url := fmt.Sprintf("%s/%s", b.base, code)

	b.request = request.NewRequest(url, nil)

	b.request.SetHeader("Content-Type", "application/json")

	bank, err := b.request.Get()

	if err != nil {
		return nil, err
	}

	return bankResponse(bank)
}

func banksResponse(response string) (*[]Bank, error) {

	var banks []Bank

	err := json.Unmarshal([]byte(response), &banks)

	if err != nil {
		return nil, err
	}

	return &banks, nil
}


func bankResponse(response string) (*Bank, error) {

	var bank Bank

	err := json.Unmarshal([]byte(response), &bank)

	if err != nil {
		return nil, err
	}

	return &bank, nil
}