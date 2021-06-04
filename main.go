package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello world")
	})

	http.ListenAndServe(":9090", nil)

	//fmt.Println("")
	//var i int
	//var err error

	//i, err = AddRequirement("Build in go", "The build environment must use Go", "R100", 100)
	//if err == nil {
	//	println("Success: ", i)
	//}
}
