package main

type ProtocolType uint8

const (
	PROTOCOL_TYPE_HTTP_RUNNER = 1
	PROTOCOL_TYPE_HTTPS_RUNNER = 2
)

type RunnerEndpoint struct {
	url string
	protocol ProtocolType
}

type Job struct {
	//universal unique id for a job
	id string
	//command
	cmd string
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

type NodeStatus uint8

const (
	NODE_STATUS_ACTIVE = 1
	NODE_STATUS_NEEDS_RECOVERY = 2
)

type Node struct {
	ipv4 string
	cluster_id string
	status NodeStatus
}

