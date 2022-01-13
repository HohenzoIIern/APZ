package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func TimeResponseHandler(w http.ResponseWriter, req *http.Request) {
	tm := TimeS{Time: time.Now().Format(time.RFC3339)}
	res, _ := json.Marshal(tm)
	fmt.Fprintf(w, "%s\n", res)
}

type TimeS struct {
	Time string `json:"time"`
}

func main() {
	log.Printf("HTTP server started")
	http.HandleFunc("/time", TimeResponseHandler)
	http.ListenAndServe(":8795", nil)
}
