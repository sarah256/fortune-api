package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/datfiles", GetFortune).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}

func GetFortune(w http.ResponseWriter, r *http.Request) {}
