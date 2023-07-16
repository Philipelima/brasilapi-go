package test

import (
	"testing"
	"github.com/philipelima/brasilapi-go/banks"
)


func TestAll(t *testing.T) {

	allBanks, err := banks.V1().All()

	if err != nil {
		t.Error(err)
	}

	for _, bank := range *allBanks {
		if isABankStruct(bank) == false {
			t.Fatal("there's a not bank struct")
		}
	}
}


func isABankStruct(i interface{}) bool {
	switch i.(type) {
		case banks.Bank:
			return true
		case *banks.Bank:
			return true
		default:
			return false
	}
}