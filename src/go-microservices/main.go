package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	// import detail package
	details "github.com/sikgyu/go-microservices/details"
	// you need to install the package
	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// print some information
	log.Println("Checking application health")
	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}
	// encoding the response "w"
	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// print http success status
	log.Println("Serving the homepage")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and running")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching the details")
	hostname, err := details.GetHostname()
	// if error is not nil, create an exception (panic)
	if err != nil {
		panic(err)
	}
	// print hostname and IP address if there's no error
	IP, _ := details.GetIP()
	fmt.Println(hostname, IP)
}

func main() {

	// create a new router
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/details", detailsHandler)

	log.Fatal(http.ListenAndServe(":80", r))
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, you've requested: %s with token: %s\n", r.URL.Path, r.URL.Query().Get("token"))
// }

// func main() {
// 	http.HandleFunc("/", rootHandler)

// 	// create file server
// 	fs := http.FileServer(http.Dir("static/"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	// listening on port 80
// 	log.Println("Web Server has started")
// 	http.ListenAndServe(":80", nil)
// }
