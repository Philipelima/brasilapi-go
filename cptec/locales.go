package cptec

import (
	"fmt"
	"log"

	"github.com/philipelima/brasilapi-go/utils/parse"
	"github.com/philipelima/brasilapi-go/utils/request"
)


type City struct {
	Name string  `json:"nome"`
	State string `json:"estado"`
	Id int64 `json:"id"`
}

func (c *CptecV1) Cities() (*[]City, error){

	var cities *[]City

	url := c.bases["locales"]

	request := request.NewRequest(url, nil)
	request.SetHeader("Content-Type", "application/json")

	response, err := request.Get()

	if err != nil {
		log.Fatal(err)
		return nil, err 
	}

	parse.StrToStruct(response, &cities)

	return cities, err
}	


func (c *CptecV1) City(cityName string) (*City, error) {

	var city *City

	url := fmt.Sprintf("%s/%s", c.bases["locales"], cityName)

	request := request.NewRequest(url, nil)
	request.SetHeader("Content-Type", "application/json")

	response, err := request.Get()

	if err != nil {
		log.Fatal(err)
		return nil, err 
	}

	parse.StrToStruct(response, &city)

	return city, nil 
}