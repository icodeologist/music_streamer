package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func StartandListenServer() {
	http.HandleFunc("/", SendMusicinBytesToClient)
	http.ListenAndServe(":3000", nil)
}

func SendMusicinBytesToClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "audio/mpeg")
	//check request header
	var start int
	var end int

	rangeHeader := r.Header.Get("Range")
	if rangeHeader == "" {
		log.Println("No range header found")
		return
	} else {
		log.Printf("The requested bytes : %v", rangeHeader)
		fmt.Printf("Type of the bytes : %T\n", rangeHeader)
		ans := strings.Split(rangeHeader, "-")
		start, _ = strconv.Atoi(ans[0])
		end, _ = strconv.Atoi(ans[1])
		fmt.Println("start : ", start)
		fmt.Println("end : ", end)

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
	log.Printf("Data of the file : %v", data[start:end])
	log.Printf("Length: %v", len(data))
	fmt.Printf("%T\n", data)
	w.Write(data[start:end])
	//find start
	//read bytes
	//send bytes as response along with 206
}
