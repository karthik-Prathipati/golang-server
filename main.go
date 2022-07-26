package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hi welcome to the server side of our program")

	http.HandleFunc("/bms", BmsHandler)
	http.ListenAndServe(":8080", nil)
}

type bms struct {
	Height string `json:height`
	Weight string `json:weight`
	Name   string `json:name`
}

func BmsHandler(res http.ResponseWriter, req *http.Request) {

	var height, weight string
	var name string
	switch req.Method {
	case "GET":

		fmt.Fprintf(res, "%v has a weight %v and height %v ", name, weight, height)

	case "POST":

		err := req.ParseForm()
		if err != nil {
			fmt.Fprintf(res, "error from parsing the form:= %v\n", err)
		}

		fmt.Fprintf(res, "Responce from the form after parsing: %v\n", req.PostForm)

		user := bms{}
		err2 := json.NewDecoder(req.Body).Decode(&user)
		if err2 != nil {
			fmt.Println(err2)
		}
		// w := req.FormValue("weight")
		// n := req.FormValue("name")

		fmt.Println("The user is ", user)
	}
}
