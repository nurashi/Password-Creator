package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func generatePassword(length int, characters string) string {
	password := ""
	for i := 0; i < length; i++ {
		index := rand.Intn(len(characters))
		password += string(characters[index])
	}
	return password
}

func passwordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	lengthStr := r.URL.Query().Get("length")
	length, err := strconv.Atoi(lengthStr)
	if err != nil || length <= 0 {
		http.Error(w, "Invalid 'length' parameter. It must be a positive number.", http.StatusBadRequest)
		return
	}

	baseOfPassword := "wewmkwem123n12k3nlasnlsqjcgjnasd@@35"
	rand.Seed(time.Now().UnixNano())
	password := generatePassword(length, baseOfPassword)

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, password)
}

func main() {
	http.HandleFunc("/generate-password", passwordHandler)

	fmt.Println("Server is running on http://localhost:1111")
	if err := http.ListenAndServe(":1111", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
