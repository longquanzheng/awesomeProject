package main

type Object interface {}

type JobManagement interface {
	GetJob(id string) Job
	AddJob(job Job) bool
	UpdateJob(job Job, param Object) bool
}

/************************************************************************************************
 * A scheduler's responsibility:
 *	equally distribute the jobs(according to workload) into different nodes, and schedule triggers as
 *	requesting execution to the associated Job Runner Cluster. Update the result as the execution finishes.
*************************************************************************************************/

/************************************************************************************************
 * A monitor's responsibility:
 *	Proactively monitor triggering process is working or not in other nodes.
 *	Each node is responsible for monitoring another, so that all of them forms a monitoring loop.
 *	In a case that one node detects another is not trigger execution on time for more than a threshold,
 *	it would tell Cluster Manager to check the health of that another node. Remove unhealthy node from
 *	cluster if necessary
 *************************************************************************************************/

