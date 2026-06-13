package main

import (
	"io"
	"log"
	"net/http"
)

func ClientStart() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:3000/", nil)
	if err != nil {
		log.Fatalf("Err to generate a request: %v", err)
	}
	req.Header.Add("Range", "0-1111")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Err connecting to server : %v", err)
	}
	resp.Body.Close()
	log.Printf("Recieved response status : %v", string(resp.Status))
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read the resp.body : %v", err)
	}
	log.Println("Body data : ", string(body))
}
