package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
)
var Myip struct{
	Ipshow string `json:"origin"`
}

func main() {
	// TODO: read info from "http://httpbin.org/ip" and write origin IP to stdout

	resp, err := http.Get("http://httpbin.org/ip")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	response,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}


	err = json.Unmarshal(response, &Myip)
	if err !=nil {
		panic(err)
					}

	fmt.Println(Myip.Ipshow)



	}



