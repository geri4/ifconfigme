package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func getClientAddress(w http.ResponseWriter, r *http.Request) {
	var realIP string
	if len(r.Header["X-Real-Ip"]) != 0 {
		realIP = r.Header["X-Real-Ip"][0]
	} else {
		realIP = strings.Split(r.RemoteAddr, ":")[0]
	}
	fmt.Fprintf(w, "%s\n", realIP)
}

func main() {
	http.HandleFunc("/", getClientAddress) //set router
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
