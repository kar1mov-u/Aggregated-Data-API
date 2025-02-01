package src

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type WeatherResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name      string `json:"name"`
	Country   string `country:"country"`
	LocalTime string `localtime:"localtime"`
}
type Current struct {
	Temperature float64 `json:"temp_c"`
	Condition   struct {
		Text string `json:"text"`
	} `json:"condition"`
	Feels float64 `json:"feelslike_c"`
}

type FinalResponse struct {
	Temp        float64 `json:"temp"`
	Description string  `json:"description"`
	Feels       float64 `json:"feels"`
}

func WeatherApiCall(location string) (FinalResponse, error) {
	var weahterResp WeatherResponse
	var finalResp FinalResponse
	API_KEY := GetKey("WEATHER_API_KEY")
	fmt.Println(API_KEY)
	URL := "https://api.weatherapi.com/v1/current.json?q=" + location + "&" + "key=" + API_KEY
	fmt.Println(URL)
	resp, err := http.Get(URL)

	if err != nil {
		log.Fatal("Cuoldnt make request to url", URL)
		return finalResp, err
	}

	err = json.NewDecoder(resp.Body).Decode(&weahterResp)

	finalResp = FinalResponse{Temp: weahterResp.Current.Temperature, Description: weahterResp.Current.Condition.Text, Feels: weahterResp.Current.Feels}
	if err != nil {
		log.Fatalf("Couldnt Parse the JSON:%v", err)
	}

	return finalResp, nil
}
