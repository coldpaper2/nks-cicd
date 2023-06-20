package main

import (

	"fmt"
	"log"
	"net/http"

)


func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Jenkins & GitHub WebHook TEST !!!! ")

}

func main() {

	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8888", nil))

}






