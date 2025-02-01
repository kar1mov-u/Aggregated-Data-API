package src

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	//Extract city name
	city_name := r.URL.Query().Get("location")
	if city_name == "" {
		http.Error(w, "City name cannot be empty", http.StatusBadRequest)
		return
	}
	fmt.Println(city_name)

	//made request to weather
	weatherResp, err := WeatherApiCall(city_name)
	if err != nil {
		http.Error(w, "FAiled on weather api request:%v", http.StatusInternalServerError)
		return
	}
	fmt.Println(weatherResp)
	err = json.NewEncoder(w).Encode(weatherResp)
	if err != nil {
		http.Error(w, "Failed to Wrire Json", http.StatusBadRequest)
		return
	}

	//made request to news

}
