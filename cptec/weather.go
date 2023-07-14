package cptec

import (
	"fmt"
	"log"

	"github.com/philipelima/brasilapi-go/utils/parse"
	"github.com/philipelima/brasilapi-go/utils/request"
)


type CurrentWeather struct {
	IcaoCode   string `json:"codigo_icao"`
	UpadatedAt string `json:"atualizado_em"`
	AtmosphericPressure string `json:"pressao_atmosferica"`
	Visibility string `json:"visibilidade"`
	Wind int64 `json:"vento"`
	WindDirection int64 `json:"direcao_vento"`
	Moisture  int64 `json:"umidade"`
	Condition string `json:"condicao"`
	ConditionDescription string `json:"condicao_Desc"`
}


func (c *CptecV1) CurrentWeatherInCapitals() (*[]CurrentWeather, error) {

	var currentWeather *[]CurrentWeather

	url := fmt.Sprintf("%s/%s", c.bases["current_weather"], "capital")

	request := request.NewRequest(url, nil)
	request.SetHeader("Content-Type", "application/json")

	response, err := request.Get()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	parse.StrToStruct(response, currentWeather)

	return currentWeather, err
}


func (c *CptecV1) CurrentWeatherInAirport(icaoCode string) (*CurrentWeather, error) {

	var currentWeather *CurrentWeather

	url := fmt.Sprintf("%s/%s/%s", c.bases["current_weather"], "aeroporto", icaoCode)

	request := request.NewRequest(url, nil)
	request.SetHeader("Content-Type", "application/json")

	response, err := request.Get()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	parse.StrToStruct(response, currentWeather)

	return currentWeather, err
}


func (c *CptecV1) CurrentWeatherInCity(code int64) (*CurrentWeather, error) {

	var currentWeather *CurrentWeather

	url := fmt.Sprintf("%s/%s/%d", c.bases["current_weather"], "previsao", code)

	request := request.NewRequest(url, nil)
	request.SetHeader("Content-Type", "application/json")

	response, err := request.Get()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	parse.StrToStruct(response, currentWeather)

	return currentWeather, err
}


