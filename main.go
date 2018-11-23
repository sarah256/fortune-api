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
	fmt.Println(a, "th datfile")
	fmt.Println(len(datfiles), "total datfiles")

	// Select random file
	randFile := datfiles[a]

	fmt.Println("datfile name:", randFile)

	file, err := os.Open(randFile)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    defer file.Close()

    reader := bufio.NewReader(file)
    scanner := bufio.NewScanner(reader)

    quotes := 0

    for scanner.Scan() {
    	if scanner.Text() == "%" {
    		quotes++
    	}
    }

	file2, err2 := os.Open(randFile)

    if err2 != nil {
        fmt.Println(err2)
        os.Exit(1)
    }

    defer file2.Close()

    reader2 := bufio.NewReader(file2)
    scanner2 := bufio.NewScanner(reader2)

	rand.Seed(time.Now().UnixNano())
	b := rand.Intn(quotes)
	fmt.Println(b, "th quote")
	fmt.Println(quotes, "total quotes")

    count := 1
    quote := ""

    for scanner2.Scan() {
    	if scanner2.Text() != "%" && count == b {
            quote = quote + scanner2.Text() + "\n"
    	} else if (scanner2.Text() == "%" && (count > b)) {
	    	break
    	} else if scanner2.Text() == "%" {
    		count++
    	}
    }

    fmt.Println(quote)

    // router := mux.NewRouter()
    // router.HandleFunc("/people", GetPeople).Methods("GET")
    // router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    // log.Fatal(http.ListenAndServe(":8000", router))
}