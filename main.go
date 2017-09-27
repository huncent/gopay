package main

import (
	"net/http"
	"log"
	"io/ioutil"
)

func main()  {
	http.HandleFunc("/notify", func(writer http.ResponseWriter, request *http.Request) {
		body ,_:=ioutil.ReadAll(request.Body)
		log.Println(string(body))
	})
	if err :=http.ListenAndServe(":8080",nil);err!=nil{
		log.Println(err.Error())
	}
}