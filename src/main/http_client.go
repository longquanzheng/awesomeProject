package main

import (
	"io"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"fmt"
	"encoding/json"
	"os/exec"
	"strings"
	"bytes"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}


type HttpClient struct {
	threads int
	port int
}

func (client HttpClient) Run(port int)  {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/run/{execution_id}", execute).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func execute(res http.ResponseWriter, req *http.Request) {
	log.Println("Responsing to /hello request")
	log.Println(req.UserAgent())

	vars := mux.Vars(req)
	name := vars["name"]

	decoder := json.NewDecoder(req.Body)
	job := new (Job)
	error := decoder.Decode(&job)
	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}
	//run...
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
	

	outgoingJSON, err := json.Marshal(cmd.Stdout)
	if err != nil {
		log.Println(error.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, string(outgoingJSON))

	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, "Hello:", name)
}