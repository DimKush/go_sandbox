package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func getRequestClient(){
	client := http.Client{}

	reqArgs := url.Values{}
	reqArgs.Add("query" , "go go Johny")
	reqArgs.Add("limit", "5")

	urlAddress := "https://google.com"
	reqUrl, err := url.Parse(urlAddress)
	if err != nil {
		log.Fatalf("Error during parse url: %s", urlAddress)
	}

	reqUrl.Path = "/search"
	reqUrl.RawQuery= reqArgs.Encode()

	req, err := http.NewRequest("GET", reqUrl.String(),nil)
	if err != nil {
		log.Fatalf("Error during create new request: %s", err)
	}
	fmt.Println(reqUrl)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error during responce: %s", err)
	}

	fmt.Printf("urlAddress %s responced with code : %d \n", urlAddress, resp.StatusCode)
}

type addRequest struct {
	Id 		 int `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

func postRequestClient(){
	client := &http.Client{}

	addReq := &addRequest {
		Id : 123,
		Title : "for loop",
		Text : "text",
	}

	var body bytes.Buffer
	json.NewEncoder(body).Encode(addReq)
	
	req , err := http.NewRequest("POST", "https://google.com", body)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	resp, err := client.Do(req)
}

func main(){

}
