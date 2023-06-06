package tenkiGetter

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
)

type Forecast struct {
	office      string
	url         string
	rawForecast *simplejson.Json
}

func (forecast *Forecast) GetForecast(config *Config) (*Forecast, error) {
	request, err := http.NewRequest("GET", forecast.buildUrl(config), nil)
	if err != nil {
		return nil, err
	}
	data, err := sendRequest(request, config)
	if err != nil {
		return nil, err
	}
	return parseForecast(data)
}

func (forecast *Forecast) GetData(tag string) string {
	return forecast.rawForecast.Get(tag).MustString()
}

func parseForecast(data []byte) (*Forecast, error) {
	var returnValue Forecast
	dec, err := simplejson.NewJson(data)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		return nil, err
	}
	returnValue.rawForecast = dec
	return &returnValue, nil
}

func NewForecast() *Forecast {
	return &Forecast{}
}

func (forecast *Forecast) buildUrl(config *Config) string {
	config.OfficeCode = "270000"
	return fmt.Sprintf("https://www.jma.go.jp/bosai/forecast/data/%s/%s.json", config.RunMode.GetMode(), config.OfficeCode)
}

func sendRequest(request *http.Request, config *Config) ([]byte, error) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()
	return handleResponse(response)
}

func handleResponse(response *http.Response) ([]byte, error) {
	if response.StatusCode/100 != 2 {
		data, _ := io.ReadAll(response.Body)
		fmt.Println("response body:", string(data))
		return []byte{}, fmt.Errorf("response status code %d", response.StatusCode)
	}
	return io.ReadAll(response.Body)
}
