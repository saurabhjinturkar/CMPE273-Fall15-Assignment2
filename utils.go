package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func fetchGeocode(l Location) Coordinate {
	baseUrl := "https://maps.googleapis.com/maps/api/geocode/json?address="
	address := l.Address + " " + l.City + " " + l.State
	authUrl := "&key=AIzaSyCv2TRAAF2y9GokOKq3UTovq5KYT2R2qDA"
	url := strings.Replace(baseUrl+address+authUrl, " ", "%20", -1)
	fmt.Println("URL:", url)
	println("Hitting url:", baseUrl, address, authUrl)
	res, err := http.Get(url)
	defer res.Body.Close()
	checkErr(err)
	fmt.Println(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	checkErr(err)
	resp := make(map[string]interface{})
	err = json.Unmarshal(body, &resp)
	checkErr(err)
	fmt.Println(resp)
	results := resp["results"].([]interface{})[0].(map[string]interface{})["geometry"].(map[string]interface{})["location"]
	lat := results.(map[string]interface{})["lat"]
	lng := results.(map[string]interface{})["lng"]

	fmt.Println(lat, lng)
	return Coordinate{lat.(float64), lng.(float64)}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
