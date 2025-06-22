package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type weather struct {
	Timezone        string
	Resolvedaddress string

	Days []struct {
		Datetime    string
		Tempmax     float32
		Tempmin     float32
		Temp        float32
		Conditions  string
		Description string
	}
}

func main() {

	url := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/28422?unitGroup=us&key=HFFE44LHLML3ATWX6R694ZVLY&contentType=json"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	dataBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var result weather
	err = json.Unmarshal(dataBytes, &result)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Address:", result.Resolvedaddress)
	fmt.Println("Date:", result.Days[0].Datetime)
	fmt.Println("Temperature:", result.Days[0].Temp)
	fmt.Println("Max Temperature:", result.Days[0].Tempmax)
	fmt.Println("Min Temperature:", result.Days[0].Tempmin)
	fmt.Println("Conditions:", result.Days[0].Conditions)
	fmt.Println("Description:", result.Days[0].Description)
}
