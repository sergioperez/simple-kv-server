package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"strings"
	"strconv"
	"sync"
	"time"
	"os"
)

var mutex sync.Mutex
var projectStatus map[string]string
var dataTimestamp map[string]time.Time
var usingTimestamp bool

// Check erasable values every minute
func cleanupLoop(maxAgeSeconds float64) {
	for {
		time.Sleep(60 * time.Second)
		t := time.Now()
		for project, creationTime := range dataTimestamp {
			mutex.Lock()
			if t.Sub(creationTime).Seconds() >= maxAgeSeconds {
				delete(projectStatus, project)
				delete(dataTimestamp, project)
			}
			mutex.Unlock()
		}
	}
}

func setKeyHandler(w http.ResponseWriter, r *http.Request) {
	// Set key value
	parameters := strings.Split(r.URL.Path, "/")
	if len(parameters) <= 2 {
		fmt.Println(len(parameters))
		return
	}
	projectName, projectHealth := parameters[1], parameters[2]

	if usingTimestamp {
		mutex.Lock()
		projectStatus[projectName] = projectHealth
		dataTimestamp[projectName] = time.Now()
		mutex.Unlock()
	} else {
		projectStatus[projectName] = projectHealth
	}

}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		setKeyHandler(w, r)
	}

	// Return list of values
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

	maxKeyAgeSeconds, _ := strconv.Atoi(os.Getenv("MAX_KEY_AGE_SECONDS"))
	if maxKeyAgeSeconds > 0 {
		usingTimestamp = true
		dataTimestamp = make(map[string]time.Time, 50)
		go cleanupLoop(float64(maxKeyAgeSeconds))
	}

	fmt.Println("Started SimpleKVServer")
	handleRequests()
}
