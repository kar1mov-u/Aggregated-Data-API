package src

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type WeatherResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name string `json:"name"`
}
type Current struct {
	Temperature float64 `json:"temp_c"`
	Condition   struct {
		Text string `json:"text"`
	} `json:"condition"`
	Feels float64 `json:"feelslike_c"`
}

type FinalResponseWeather struct {
	Temp        float64 `json:"temp"`
	Description string  `json:"description"`
	Feels       float64 `json:"feels"`
}

type NewsResp struct {
	Articles []Articles `json:"articles"`
}

type Articles struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Url    string `json:"url"`
}

func WeatherApiCall(location string, wg *sync.WaitGroup, ch chan FinalResponseWeather) {
	defer wg.Done()
	var weahterResp WeatherResponse
	var finalResp FinalResponseWeather

	API_KEY := GetKey("WEATHER_API_KEY")
	URL := "https://api.weatherapi.com/v1/current.json?q=" + location + "&" + "key=" + API_KEY
	resp, err := http.Get(URL)

	if err != nil {
		log.Fatal("Cuoldnt make request to url", URL)
	}
	//Serialize to the Struct
	err = json.NewDecoder(resp.Body).Decode(&weahterResp)

	//Create a desired struct from response
	finalResp = FinalResponseWeather{Temp: weahterResp.Current.Temperature, Description: weahterResp.Current.Condition.Text, Feels: weahterResp.Current.Feels}
	if err != nil {
		log.Fatalf("Couldnt Parse the JSON:%v", err)
	}
	ch <- finalResp

}

func NewsApiCall(location string, wg *sync.WaitGroup, ch chan NewsResp) {
	wg.Done()
	API_KEY := GetKey("NEWS_API_KEY")
	URL := "https://newsapi.org/v2/everything?q=" + location + "&pageSize=5&apiKey=" + API_KEY
	var newsResp NewsResp

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatalf("Failed to send request to NEWS")
	}
	//Serialize response to the Struct

	err = json.NewDecoder(resp.Body).Decode(&newsResp)
	if err != nil {
		log.Fatal("Failed to Serialize Json")
	}
	ch <- newsResp

}
