package main

type ProtocolType uint8

const (
	HTTP_RUNNER = 1
	HTTPS_RUNNER = 1
)

type RunnerEndpoint struct {
	url string
	protocol ProtocolType
}

type Job struct {
	//universal unique id for a job
	id string
	//owner id
	owner string
	//is enabled
	enabled bool
	//description for a job
	desc string
	//cron expression: http://crontab.org/
	cron string
	//runner endpoint information
	runners []RunnerEndpoint
}


