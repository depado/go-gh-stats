package main

import (
	"encoding/json"
	"net/http"
)

func fetchUserData(usr string) (usra userAPI, err error) {
	client := new(http.Client)
	req, err := http.NewRequest("GET", "https://api.github.com/users/"+usr, nil)
	if err != nil {
		return
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&usra)
	if err != nil {
		return
	}
	return
}

func fetchReposData(usr string) (repos allRepos, err error) {
	client := new(http.Client)
	req, err := http.NewRequest("GET", "https://api.github.com/users/"+usr+"/repos", nil)
	if err != nil {
		return
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&repos)
	if err != nil {
		return
	}
	return
}
