package main

import (
	"fmt"
	"io"
	"net/http"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	// check if the request method is POST (Webhook sends data this way)
	if r.Method == http.MethodPost {
		//read the request body
		// data := make([]byte, r.ContentLength)
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		// you can do anything with your data here
		// here I simply print
		fmt.Println("recieved webhook data:", string(data))

		//send the response back to the sender
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Webhook received successfully!")
		w.Write([]byte("\nreceived the message: " + string(data)))
	} else {
		// response with error for unsupported mehtods
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)

	//start the server on port 8080
	fmt.Println("webhook receiver server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
