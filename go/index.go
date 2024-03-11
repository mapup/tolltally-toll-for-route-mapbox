package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	MAPBOX_API_KEY     string = os.Getenv("MAPBOX_API_KEY")
	TOLLGURU_API_KEY   string = os.Getenv("TOLLGURU_API_KEY")
)

const (
	MAPBOX_API_URL            = "https://api.mapbox.com/directions/v5/mapbox/driving"
	MAPBOX_GEOCODE_API_URL    = "https://api.mapbox.com/geocoding/v5/mapbox.places"
	TOLLGURU_API_URL  		  = "https://apis.tollguru.com/toll/v2"
	POLYLINE_ENDPOINT 		  = "complete-polyline-from-mapping-service"

	source      string = "Philadelphia, PA"
	destination string = "New York, NY"
)

// Explore https://tollguru.com/toll-api-docs to get the best of all the parameters that tollguru has to offer
var requestParams = map[string]interface{}{
	"vehicle": map[string]interface{}{
		"type": "2AxlesAuto",
	},
	// Visit https://en.wikipedia.org/wiki/Unix_time to know the time format
	"departure_time": "2021-01-05T09:46:08Z",
}

func main() {
	url := fmt.Sprintf("%s/%s.json?access_token=%s", MAPBOX_GEOCODE_API_URL, source, MAPBOX_API_KEY)
	spaceClient := http.Client{
		Timeout: time.Second * 200, // Timeout after 200 seconds
	}
	// fmt.Println(records)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	var result map[string]interface{}

	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		log.Fatal(result)
	}

	Origin_longitute := result["features"].([]interface{})[0].(map[string]interface{})["geometry"].(map[string]interface{})["coordinates"].([]interface{})[0]
	Origin_latitude := result["features"].([]interface{})[0].(map[string]interface{})["geometry"].(map[string]interface{})["coordinates"].([]interface{})[1]

	url_add := fmt.Sprintf("%s/%s.json?access_token=%s", MAPBOX_GEOCODE_API_URL, destination, MAPBOX_API_KEY)

	// fmt.Println(records)
	req_add, err := http.NewRequest(http.MethodGet, url_add, nil)
	if err != nil {
		log.Fatal(err)
	}

	req_add.Header.Set("User-Agent", "spacecount-tutorial")

	res_add, getErr := spaceClient.Do(req_add)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res_add.Body != nil {
		defer res_add.Body.Close()
	}

	body_add, readErr := ioutil.ReadAll(res_add.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	var result_add map[string]interface{}

	jsonerr := json.Unmarshal(body_add, &result_add)
	if jsonerr != nil {
		log.Fatal(result_add)
	}

	Destination_longitute := result_add["features"].([]interface{})[0].(map[string]interface{})["geometry"].(map[string]interface{})["coordinates"].([]interface{})[0]
	Destination_latitute := result_add["features"].([]interface{})[0].(map[string]interface{})["geometry"].(map[string]interface{})["coordinates"].([]interface{})[1]

	url_mapbox := fmt.Sprintf("%s/%v,%v;%v,%v?geometries=polyline&access_token=%s&overview=full", MAPBOX_API_URL, Origin_longitute, Origin_latitude, Destination_longitute, Destination_latitute, MAPBOX_API_KEY)
	spaceClient1 := http.Client{
		Timeout: time.Second * 15, // Timeout after 15 seconds
	}

	req1, err := http.NewRequest(http.MethodGet, url_mapbox, nil)
	if err != nil {
		log.Fatal(err)
	}

	req1.Header.Set("User-Agent", "spacecount-tutorial")

	res1, getErr := spaceClient1.Do(req1)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res1.Body != nil {
		defer res1.Body.Close()
	}

	body1, readErr := ioutil.ReadAll(res1.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	var result_mapbox map[string]interface{}

	jsonEr := json.Unmarshal(body1, &result_mapbox)
	if jsonEr != nil {
		log.Fatal(result)
	}

	polyline := result_mapbox["routes"].([]interface{})[0].(map[string]interface{})["geometry"].(string)

	// Tollguru API request
	url_tollguru := fmt.Sprintf("%s/%s", TOLLGURU_API_URL, POLYLINE_ENDPOINT)

	params := map[string]interface{}{
		"source":   "mapbox",
		"polyline": polyline,
	}

	for k, v := range requestParams {
		params[k] = v
	}

	requestBody, err := json.Marshal(params)

	request, err := http.NewRequest("POST", url_tollguru, bytes.NewBuffer(requestBody))
	request.Header.Set("x-api-key", TOLLGURU_API_KEY)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		log.Fatal(err)
	}

	var cost map[string]interface{}
	json_err := json.Unmarshal([]byte(body), &cost)
	if json_err != nil {
		log.Fatal(result)
	}

	toll := cost["route"].(map[string]interface{})["costs"].(map[string]interface{})["tag"]
	fmt.Printf("\nThe toll rate for Source Longitude: %v,  Source Latitude: %v, Destination Longitude: %v, Destination Latitude: %v is %v\n", Origin_longitute, Origin_latitude, Destination_longitute, Destination_latitute, toll)

}
