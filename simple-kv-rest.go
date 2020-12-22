package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"strings"
)

var projectStatus map[string]string

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		parameters := strings.Split(r.URL.Path, "/")
		projectName, projectHealth := parameters[1], parameters[2]
		projectStatus[projectName] = projectHealth
	}

	projects, err := json.Marshal(projectStatus)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(projects))
}


func handleRequests() {
	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	projectStatus = make(map[string]string, 50)
	fmt.Println("Started SimpleKVServer")
	handleRequests()
}
