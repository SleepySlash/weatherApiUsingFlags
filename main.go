package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	getWeather := flag.String("weather", "", "the place you want to know the weather of")
	getWeatherDays := flag.Int("days", 1, "date of the day you want to know the weather for")
	verbose := flag.Bool("verbose", false, "Display the weather")

	flag.Parse()

	if *getWeather == "" {
		fmt.Println("Please enter a place name")
		return
	}
	output := weatherApi(*getWeather, *getWeatherDays)
	if *verbose {
		fmt.Printf("the weather of the place: %s\nfor days: %d\n   %s", *getWeather, *getWeatherDays, output)
	} else {
		fmt.Printf("the weather today of the place: %s\n   %s", *getWeather, output)
	}
}

func weatherApi(location string, days int) string {
	theUrl := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=3eb7808380a842e39b5131134241111&q=%s&days=%d", location, days)
	response, err := http.Get(theUrl)
	if err != nil {
		log.Fatalf("Failed to get the request %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
