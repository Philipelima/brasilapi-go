package test

import (
	"log"
	"testing"

	"github.com/philipelima/brasilapi-go/cnpj"
)


func TestCnpj(t *testing.T) {

	cnpjNumber := "06990590000123"

	cnpj, err := cnpj.V1().Get(cnpjNumber)

	if err != nil {
		t.Fatal(err)
	}

	if isAValidCnpjStruct(cnpj) == false {
		log.Fatal("it's not a valid cnpj struct", cnpj)
	}

	
	if cnpj.Cnpj != cnpjNumber {
		log.Fatal("it's not a valid cnpj struct", cnpj)
	}
}


func isAValidCnpjStruct(i interface{}) bool {
	switch i.(type) {
		case *cnpj.Cnpj:
			return true
		default:
			return false
	}
}