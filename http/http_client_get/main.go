package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main(){
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
