package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	toggleLights()
}

type state struct {
	on bool
}

func toggleLights() {
	client := &http.Client{}

	resp, err := client.Get("http://192.168.1.19/api/sUyXfdlBLeJGHkE2uewJdEs0HQHqycTNXTQPWVfJ/lights/1")
	if err != nil {
		log.Fatalln(err)
	}
	// Why is this needed
	defer resp.Body.Close()

	//store response in a variable
	var dat map[string]map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &dat)

	currentState := dat["state"]["on"]

	var converted string
	converted = fmt.Sprint(currentState)

	fmt.Println("converted: " + converted)
	newState := state{}
	if converted == "true" {
		fmt.Println("State was true, now false")
		newState.on = false
		fmt.Println("Turning off Light!!!!")
	} else if converted == "false" {
		newState.on = true
		fmt.Println("Turning on Light!!!")
	}

	// Preparing Post Body
	postBody, _ := json.Marshal(map[string]bool{
		"on": newState.on,
	})
	responseBody := bytes.NewBuffer(postBody)

	// Send the post requests
	req, err := http.NewRequest(http.MethodPut, "http://192.168.1.19/api/sUyXfdlBLeJGHkE2uewJdEs0HQHqycTNXTQPWVfJ/lights/1/state", responseBody)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err = client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp)
}

func getLights() {

	client := &http.Client{}

	resp, err := client.Get("http://192.168.1.19/api/sUyXfdlBLeJGHkE2uewJdEs0HQHqycTNXTQPWVfJ/lights")

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)

}
