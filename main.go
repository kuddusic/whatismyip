package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
)

type response struct {
	IP            string `json:"ip"`
	UserAgent     string `json:"user-agent"`
	XForwardedFor string `json:"x-forwarded-for"`
}

func extractIP(remoteAddr string) string {
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return remoteAddr
	}
	return host
}

func metadataHandler(w http.ResponseWriter, r *http.Request) {
	resp := response{
		IP:            extractIP(r.RemoteAddr),
		UserAgent:     r.Header.Get("User-Agent"),
		XForwardedFor: r.Header.Get("X-Forwarded-For"),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", metadataHandler)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
