package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type MyHandler struct{}

func DoSomething(query string, limit int) (string, error){
	return "some", nil
}
func (data *MyHandler) ServeHTTP(resp http.ResponseWriter, req * http.Request){
	if req.URL.Path == "/dump"{
		args := req.URL.Query()
		query := args.Get("query")
		limit, err := strconv.Atoi(args.Get("limit"))
		if err != nil{
			resp.WriteHeader(http.StatusBadRequest)
		}

		results, err := DoSomething(query, limit)
		if err != nil {
			resp.WriteHeader(http.StatusNotFound)
			return
		}

		resp.Header().Set("Content-Type","application/json; charset=utf-8")
		resp.WriteHeader(http.StatusOK)

		err = json.NewEncoder(resp).Encode(results)
		if err != nil{
			resp.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func main(){
	handler := &MyHandler{}
	_, err := run("8080",handler)
	if  err != nil {
		fmt.Printf("Cannot run %v", err)
		os.Exit(1)
	}

}

func run(port string, handler * MyHandler)  (*http.Server, error) {

	server := http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: handler,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &server, server.ListenAndServe()
}

func stop(server *http.Server) {
	//server.Shutdown()
}