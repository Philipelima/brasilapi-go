package request

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Request struct {
	url     string
	method  string
	headers map[string]string
	payload *map[string]string 
	client  *http.Client
}


func NewRequest(url string, payload *map[string]string) *Request {
	return &Request{
		url: url,
		payload: payload,
		client:  &http.Client{},
		headers: make(map[string]string),
	}
}


func (r *Request) SetHeader(name, value string) {
	r.headers[name] = value
}


func (r *Request) RemoveHeader(name string) bool {

	_, ok := r.headers[name]

	if ok {
		delete(r.headers, name)
		return true
	}

	fmt.Printf("Sorry, the key %s dosen't exists on request headers.", name)
	return false
}


func (r *Request) Get() (string, error){

	r.method = "GET"


	if r.payload != nil {

		values := url.Values{}

		for param, value := range *r.payload {
			values.Add(param, value)
		} 

		query := values.Encode()
		r.url = fmt.Sprintf("%s/?%s", r.url, query)

	}

	return r.do()
}


func (r *Request) do() (string, error) {

	request, err := http.NewRequest(r.method, r.url, nil)

	if err != nil {
		log.Println("Sorry, we can't do this request. ERROR: ", err.Error())
		return "", err
	}

	for  name, value := range r.headers {
		request.Header.Add(name, value)
	}

	response, err := r.client.Do(request)

	if err != nil {
		log.Println("Sorry, we can't do this request. ERROR: ", err.Error())
		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println("Sorry, we can't do this request. ERROR: ", err.Error())
		return "", err
	}

	return string(body), nil
}

