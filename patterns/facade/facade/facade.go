package facade

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// CurrentWeatherDataRetriever - интерфейс, являющийся фасадом/оберткой для объекта CurrentWeatherData
// в котором производится основная логика
type CurrentWeatherDataRetriever interface {
	GetByCityAndCountryCode(city, countryCode string) (Weather, error)
	GetByGeoCoordinates(lat, lon float64) (Weather, error)
}

// CurrentWeatherData - объект с апи ключом для сервиса прогноза погоды
type CurrentWeatherData struct {
	APIkey string
}

// Weather - объект, который мы получаем из сервиса по запросу в сети
type Weather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure float64 `json:"pressure"`
		Humidity float64 `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Rain struct {
		ThreeHours float64 `json:"3h"`
	} `json:"rain"`
	Dt  uint32 `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

const (
	commonRequestPrefix              = "https://api.openweathermap.org/data/2.5/"
	weatherByCityName                = commonRequestPrefix + "weather?q=%s,%s&APPID=%s"
	weatherByGeographicalCoordinates = commonRequestPrefix + "weather?lat=%f&lon=%f&APPID=%s"
)

// GetByGeoCoordinates возвращает погоду в определенной ширине и долготе
func (c *CurrentWeatherData) GetByGeoCoordinates(lat, lon float64) (weather *Weather, err error) {
	return c.doRequest(fmt.Sprintf(weatherByGeographicalCoordinates, lat, lon, c.APIkey))
}

// GetByCityAndCountryCode возвращает погоду в определенном городе страны
func (c *CurrentWeatherData) GetByCityAndCountryCode(city, countryCode string) (weather *Weather, err error) {
	return c.doRequest(fmt.Sprintf(weatherByCityName, city, countryCode, c.APIkey))
}

// Методы выше имплементируют нужный нам фасад, потому они могут быть вызваны со стороны клиента
// Имеют "однострочную" логику, основная же логика инкапсулирована и находится ниже
// responseParser парсит тело ответа в объект Weather
func (c *CurrentWeatherData) responseParser(body io.Reader) (*Weather, error) {
	w := &Weather{}
	err := json.NewDecoder(body).Decode(w)
	if err != nil {
		return nil, err
	}

	return w, nil
}

// doRequest - осуществляет Get запрос по определенному uri и возвращает объект weather
func (c *CurrentWeatherData) doRequest(uri string) (weather *Weather, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		body, errMsg := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if errMsg == nil {
			errMsg = fmt.Errorf("%s", string(body))
		}

		err = fmt.Errorf("Status code was %d. Error meassge was:\n%s\n", resp.StatusCode, errMsg)
		return
	}

	weather, err = c.responseParser(resp.Body)
	resp.Body.Close()

	return
}
