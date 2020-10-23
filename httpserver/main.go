package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetWeather(u string) (*Weather, error) {
	req, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	w := &Weather{}
	err = json.Unmarshal(b, w)
	if err != nil {
		return nil, err
	}

	return w, nil
}
func main() {

	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	apiKey := os.Getenv("APITOKEN")
	if apiKey == ""{
		log.Fatal("Not detected api token")
	}

	request := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=Paris&appid=%s", apiKey)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		weather , err := GetWeather(request)
		if err != nil || weather == nil {
			http.Error(w, err.Error(),http.StatusInternalServerError)
		}
		_, _ = io.WriteString(w, weather.String())
	})

	log.Fatal(http.ListenAndServe(":" + port, nil))
}
