package main

import (
	"fmt"
	"os"
	"bytes"
)

/************************************************************************************************
 * Main entry for different programs: Job sheduler/Job monitor/Management daemon/Runner clients
 *************************************************************************************************/

func main() {
	arg := os.Args[1]
	fmt.Println(arg)

	if arg == "client" {
		fmt.Println("Running client...")
		var client Client = HttpClient{ port: 8080, jobs: make(map[string]bytes.Buffer)  }
		client.Run()
	}

	//start scheduler
	if arg == "scheduler" {
		fmt.Println("Running scheduler...")
		//test long running command: wget speedtest.ftp.otenet.gr/files/test1Gb.db
	}

	//TODO monitor
	//TODO apid
}
