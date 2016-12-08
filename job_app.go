package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
)

type SecretMessage struct {
	Message string
}

func main() {
	http.HandleFunc("/secret_message.json", secretMessage)

	const PORT string = ":8080"
	fmt.Printf("Listening on port %q\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func secretMessage(w http.ResponseWriter, r *http.Request) { 
	decoder := json.NewDecoder(r.Body)
	
	var accessCode SecretMessage
	err := decoder.Decode(&accessCode)

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	
	log.Println(accessCode.Message)

	var msg string
	if accessCode.Message == "JCJRMG" {
		msg = "These violent delights have violent ends."
	} else {
		msg = "That's not the secret password."
	}

	secretMessage := SecretMessage{msg}

	js, err := json.Marshal(secretMessage)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
