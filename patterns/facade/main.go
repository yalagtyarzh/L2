package main

import (
	"fmt"
	"log"

	"facade/facade"
)

const apikey = "xd"

func main() {
	weatherMap := facade.CurrentWeatherData{APIkey: apikey}

	weather, err := weatherMap.GetByCityAndCountryCode("Moscow", "RU")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Температура в Москве: %f крадусов по цельсию", weather.Main.Temp)
}
