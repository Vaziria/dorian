package main

import "fmt"
import "net/http"
import (
	"os"
	"log"
	"github.com/imroc/req"
)


func hanlderIndex(w http.ResponseWriter, r *http.Request){
	var message = "welcode"

	w.Write([]byte(message))
}


func testRequest(w http.ResponseWriter, r *http.Request){
	var message = "welcode"

	w.Write([]byte(message))
}


func handlingTest(w http.ResponseWriter, requ *http.Request){

	var r = req.New()
	var respon, err = r.Get("http://www.mocky.io/v2/5e4f810830000090002268f8")

	if err != nil {
		log.Fatal(err)
	}

	var pesan string = string(respon.Body)

	w.Write([]byte(pesan))

}




func main(){

	var port = os.Getenv("port")
	if port == "" {
		port = "9000"
	}

	var address string = "0.0.0.0:" + port

	http.HandleFunc("/", hanlderIndex)
	http.HandleFunc("/blues", handlingTest)

	fmt.Println(address)

	http.ListenAndServe(address, nil)


}

