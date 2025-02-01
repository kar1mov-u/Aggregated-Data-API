package src

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type FinalResp struct {
	Location string               `json:"location"`
	Temp     FinalResponseWeather `json:"weather"`
	News     NewsResp             `json:"news"`
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	//Extract city name
	var wg sync.WaitGroup
	ch1 := make(chan FinalResponseWeather)
	ch2 := make(chan NewsResp)
	var ch1_res FinalResponseWeather
	var ch2_res NewsResp
	var res FinalResp

	city_name := r.URL.Query().Get("location")
	if city_name == "" {
		http.Error(w, "City name cannot be empty", http.StatusBadRequest)
		return
	}

	//MADE A REQUEST TO WEATHER
	wg.Add(1)
	go WeatherApiCall(city_name, &wg, ch1)

	//made request to news
	wg.Add(1)
	go NewsApiCall(city_name, &wg, ch2)
	for i := 0; i < 2; i++ {
		select {
		case ch1_res = <-ch1:
			res.Temp = ch1_res
		case ch2_res = <-ch2:
			res.News = ch2_res
		}
	}
	wg.Wait()

	res.Location = city_name

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed)

}
