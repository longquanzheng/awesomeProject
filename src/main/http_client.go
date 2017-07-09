package main

import (
	"io"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/alexflint/go-filemutex"
	"log"
	"fmt"
	"encoding/json"
	"os/exec"
	"strconv"
	"sync"
	"bytes"
	"io/ioutil"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}


type HttpClient struct {
	port int
	jobs   map[string]bytes.Buffer
	//lock for multi process
	lock sync.Mutex
}

type PostData struct {
	CMD string
}

func (client HttpClient) Run()  {
	//exclusively running in a host/vm/container
	m, err := filemutex.New("/tmp/httpclient.lock")
	if err != nil{
		log.Println(err.Error())
		return
	}
	//lock will be released after process died
	//TODO it doesn't work for running multi processes
	m.Lock()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/run/{execution_id}", client.execute).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(client.port), router))
}

func (client HttpClient) ExecuteJob(job Job) (string, string) {
	log.Println("execute job"+job.id)
	client.lock.Lock()
	value, exists := client.jobs[job.id]
	if exists{
		//read output and return
		client.lock.Unlock()
		return "Running", value.String()
	}else{
		//start a thread running the job
		client.jobs[job.id] = bytes.Buffer{}
		client.lock.Unlock()
		go client.doExecuteJob(client.jobs[job.id], job.id, job.cmd)
		return "Started", ""
	}
}

func (client HttpClient)doExecuteJob(buffer bytes.Buffer, id string, cmd string) {
	cmdArgs := strings.Fields(cmd)

	cmd_obj := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	stdout, _ := cmd_obj.StdoutPipe()
	stderr, _ := cmd_obj.StderrPipe()
	cmd_obj.Start()

	someBytes := make([]byte, 10)
	for {
		_, err := stdout.Read(someBytes)
		if err != nil {
			break
		}
		client.lock.Lock()
		buffer.Write(someBytes)
		client.lock.Unlock()
	}

	for {
		_, err := stderr.Read(someBytes)
		if err != nil {
			break
		}
		client.lock.Lock()
		buffer.Write(someBytes)
		client.lock.Unlock()
	}

	cmd_obj.Wait()
}

func checkError(error error, resp http.ResponseWriter) bool {
	if error != nil {
		log.Println(error.Error())
		http.Error(resp, error.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

func (client HttpClient) execute(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["execution_id"]

	log.Println("Responsing to /hello request")
	log.Println(req.UserAgent())

	body, error := ioutil.ReadAll(req.Body)
	if checkError(error, resp) { return }
	log.Println(string(body))

	var d PostData
	error = json.Unmarshal(body, &d)
	if checkError(error, resp) { return }

	//run...
	job := Job{id, d.CMD, "", true, "", "", nil }
	status, out := client.ExecuteJob(job)

	//output
	resp.WriteHeader(http.StatusOK)
	fmt.Fprintf(resp, "{ \"id\": \"%s\", \"status\":\"%s\", \"result\": \"%s\", " , id, status, out)
}