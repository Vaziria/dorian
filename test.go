package main

import "fmt"
import "net/http"
import (
	"os"
)


func hanlderIndex(w http.ResponseWriter, r *http.Request){
	var message = "welcode"

	w.Write([]byte(message))
}


func testRequest(w http.ResponseWriter, r *http.Request){
	var message = "welcode"

	w.Write([]byte(message))
}






func main(){

	var port = os.Getenv("port")
	if port == "" {
		port = "9000"
	}

	var address string = "0.0.0.0:" + port

	http.HandleFunc("/", hanlderIndex)

	fmt.Println(address)

	http.ListenAndServe(address, nil)


}

