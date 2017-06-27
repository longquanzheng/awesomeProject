package main

import "github.com/samuel/go-zookeeper/zk"

type Job 
type JobManage interface {
	AddJob()
}

type JobView interface {

}

type JobModel interface {
	JobView
	JobManage
}

