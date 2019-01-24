// This code runs a webserver and then prints output on top 
// Cool!
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Println(http.ListenAndServe("localhost:8001", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v\n", time.Now())
}
