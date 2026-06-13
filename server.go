package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func StartandListenServer() {
	http.HandleFunc("/", SendMusicinBytesToClient)
	http.ListenAndServe(":3000", nil)
}

func SendMusicinBytesToClient(w http.ResponseWriter, r *http.Request) {
	//check request header
	rangeHeader := r.Header.Get("Range")
	if rangeHeader == "" {
		log.Println("No range header found")
		return
	} else {
		log.Printf("The requested bytes : %v", rangeHeader)
		fmt.Printf("Type of the bytes : %T\n", rangeHeader)
	}
	//open song
	file, err := os.Open("./audio/alexguz-funk-amp-breakbeat-541097.mp3")
	if err != nil {
		log.Fatalf("Failed to open a file : %v ", err)
	}
	defer file.Close()
	data, err := os.ReadFile("./audio/alexguz-funk-amp-breakbeat-541097.mp3")
	if err != nil {
		log.Fatalf("Failed to read a file : %v ", err)
	}
	log.Printf("Data of the file : %v", data[0:10])
	log.Printf("Length: %v", len(data))

	//find start
	//read bytes
	//send bytes as response along with 206
}
