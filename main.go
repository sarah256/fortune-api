package main

import (
    // "github.com/gorilla/mux"
    // "log"
    // "net/http"
	"math/rand"
	"time"
	"fmt"
    "path/filepath"
    "os"
    "strings"
    "bufio"
)

func main() {
    var files []string
    var datfiles []string

    root := "datfiles"
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })
    if err != nil {
        panic(err)
    }
    for _, file := range files {
    	if !strings.Contains(file, "CMakeLists.txt") && !strings.Contains(file, "off") && !strings.Contains(file, "README.md") {
        	datfiles = append(datfiles, file)
    	}
    }

	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(len(datfiles))
	fmt.Println(a)
	fmt.Println(len(datfiles))

	// Select random file
	randFile := datfiles[a]

	file, err := os.Open(randFile)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    defer file.Close()

    reader := bufio.NewReader(file)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        // fmt.Println(scanner.Text())
        // TODO: process each quote here
    }

    // router := mux.NewRouter()
    // router.HandleFunc("/people", GetPeople).Methods("GET")
    // router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    // log.Fatal(http.ListenAndServe(":8000", router))
}