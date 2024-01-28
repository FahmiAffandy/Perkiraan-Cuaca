package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiKey = "2ac786df17cb6b29163aa099f1bc7929"
const apiUrl = "http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s"

type WeatherResponse struct {
	Main struct {
		Temperature float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func getWeather(city string) (*WeatherResponse, error) {
	url := fmt.Sprintf(apiUrl, city, apiKey)
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get weather data: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON response: %v", err)
	}

	return &weatherResponse, nil
}

func main() {
	city := "Bandung" // Ganti dengan kota yang ingin Anda periksa cuacanya
	weather, err := getWeather(city)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	description := ""
	if len(weather.Weather) > 0 {
		description = weather.Weather[0].Description
	}

	temperature := weather.Main.Temperature - 273.15 // Konversi dari Kelvin ke Celsius

	fmt.Printf("Cuaca di %s: %s\n", city, description)
	fmt.Printf("Suhu: %.2f°C\n", temperature)
}
