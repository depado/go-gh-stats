package main

import (
	"encoding/json"
	"net/http"
)

// Fetch the url given in parameter and unmarshall the JSON to the given out struct
func fetchURL(url string, out interface{}) (err error) {
	client := new(http.Client)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(out)
	return
}
