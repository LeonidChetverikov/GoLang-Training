package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
)
type myip struct{
	ipshow string
}

func main() {
	// TODO: read info from "http://httpbin.org/ip" and write origin IP to stdout

	resp, err := http.Get("http://httpbin.org/ip")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	myip2:=new(myip)
	response,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(response), &myip2)
	if err !=nil {
		panic(err)
					}

	fmt.Println(myip2)



	}



