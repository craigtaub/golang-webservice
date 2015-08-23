package handlers

import (
		// "fmt"
		"io/ioutil"
		"net/http"
		"encoding/json"
)

const ARRAYURL string ="http://localhost:8080/return-jsonarray/"
const URL string = "http://localhost:8080/return-json/"

type Page struct {
    Title string
    Body  []byte
}

func (p *Page) save() error {
    filename := "data/" + p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {

    filename := "data/" + title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}



func (p *Page) load() string {

    res, err := http.Get(URL)

    if err != nil {
        panic(err.Error())
    }

    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        panic(err.Error())
    }

		var data Page // stuct type
    json.Unmarshal(body, &data)
    // fmt.Printf("Results: %v\n", data)

		return data.Title
}


func (p *Page) loadComplexPage() string {

    res, err := http.Get(ARRAYURL)

    if err != nil {
        panic(err.Error())
    }

    body, err := ioutil.ReadAll(res.Body) //body of response

    if err != nil {
        panic(err.Error())
    }

		data := make([]Page,0) // Can hand it any struct
		json.Unmarshal(body 	, &data)

		return data[1].Title
}
