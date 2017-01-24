package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// Title contains a simple title request
type Title struct {
	Title string
}

const (
	defaultPort   = "8080"
	titleCaseTool = "bin/convert_titlecase.py"
)

func titleCase(input []byte) ([]byte, error) {
	cmd := exec.Command(titleCaseTool)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	if _, err := stdin.Write(input); err != nil {
		return nil, err
	}
	if err := stdin.Close(); err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(stdout)
	if err != nil {
		return nil, err
	}
	if err := cmd.Wait(); err != nil {
		return nil, err
	}
	return result, nil
}

func titleCasePost(w http.ResponseWriter, r *http.Request) {
	var t Title
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := titleCase([]byte(t.Title))
	if err != nil {
		log.Printf("Error in titleCase(): %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Write(result)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	http.HandleFunc("/", titleCasePost)
	log.Printf("Running server at %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
