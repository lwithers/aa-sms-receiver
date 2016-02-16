/*
aa-sms-receiver is a simple server that receives SMS messages from A&A
(https://aa.net.uk/) and logs them.
*/
package main

import (
	"flag"
	"log"
	"net/http"
	"sync"
)

var (
	bindAddr = flag.String("bind", ":8123", "Address to listen on")

	lastMessage     string
	lastMessageLock sync.Mutex
)

func main() {
	http.HandleFunc("/sms", IncomingSmsHandler)
	http.HandleFunc("/last-message", GetLastMessageHandler)
	log.Fatal(http.ListenAndServe(*bindAddr, nil))
}

func IncomingSmsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("/sms: rejecting method %q", r.Method)
		http.Error(w, "Method not allowed",
			http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Printf("/sms: failed to parse form: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	from, ok := r.PostForm["oa"]
	if !ok || len(from) != 1 {
		log.Printf("/sms: expecting one \"oa\" field, got %d",
			len(from))
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	msgText, ok := r.PostForm["ud"]
	if !ok || len(from) != 1 {
		log.Printf("/sms: expecting one \"ud\" field, got %d",
			len(msgText))
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	log.Printf("Message from %.30s: %.80q", from[0], msgText[0])
	lastMessageLock.Lock()
	lastMessage = msgText[0]
	lastMessageLock.Unlock()

	w.Write([]byte("OK"))
}

func GetLastMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("/last-message: rejecting method %q", r.Method)
		http.Error(w, "Method not allowed",
			http.StatusMethodNotAllowed)
		return
	}

	lastMessageLock.Lock()
	msg := lastMessage
	lastMessage = ""
	lastMessageLock.Unlock()

	w.Write([]byte(msg))
}
