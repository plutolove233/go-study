package main

import (
	"distributed-system/map-reduce/models"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"sync"
	"time"
)

const MAX_TASK_RUNTIME = time.Duration(10 * time.Second)

type TaskState struct {
	Status    string
	StartTime time.Time
}

type Master struct {
	TaskChan   chan models.Task
	TaskPhrase string // 任务状态 map/reduce
	NMap       int
	NReduce    int
	Files      []string    // 处理的文件列表
	State      []TaskState // 每个任务的状态
	Mutex      sync.Mutex
	IsDone     bool
}

func NewMaster(files []string, nreduce int) *Master {
	m := &Master{
		TaskChan:   make(chan models.Task, 10),
		TaskPhrase: models.MapPhrase,
		NMap:       len(files),
		NReduce:    nreduce,
		Files:      files,
		State:      make([]TaskState, len(files)),
		IsDone:     false,
	}
	for index := range m.State {
		m.State[index].Status = models.TaskStatusReady
	}

	m.Server()

	return m
}

/*
启动Master结点，用于接收来自worker结点的rpc请求
*/
func (m *Master) Server() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("start master node failed, error=", err)
		}
	}()
	if err := rpc.Register(m); err != nil {
		panic(err)
	}

	l, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
	rpc.HandleHTTP()
	go http.Serve(l, nil)
}

func (m *Master) HandleTaskReq(req *models.ReqTaskArgs, reply *models.ReqTaskReply) error {
	fmt.Println("开始接收请求...")
	if !req.WorkerStatus {
		return errors.New("当前结点已下线")
	}

	// 从master的任务队列中取出任务，并交给worker结点工作
	task, ok := <-m.TaskChan
	if ok {
		reply.Task = task
		// 主节点调整任务状态
		m.State[task.TaskID].Status = models.TaskStatusRunning
		m.State[task.TaskID].StartTime = time.Now()
	} else {
		reply.TaskDone = true
	}
	return nil
}

func (m *Master) HandleTaskReport(req *models.ReportTaskArgs, reply *models.ReportTaskReply) error {
	fmt.Println("开始处理任务报告...")

	if !req.WorkerStatus {
		reply.MasterAck = false
		return errors.New("工作结点已下线")
	}

	if req.IsDone {
		// 标志任务顺利完成，调整master结点中任务的相关信息
		m.State[req.TaskIndex].Status = models.TaskStatusFinish
	} else {
		// 未完成任务却发送报告——任务出现问题
		m.State[req.TaskIndex].Status = models.TaskStatusError
	}

	reply.MasterAck = true
	return nil
}

func (m *Master) addTask(taskindex int) {
	m.State[taskindex].Status = models.TaskStatusQueue
	task := models.Task{
		FileName:   "",
		NMap:       len(m.Files),
		NReduce:    m.NReduce,
		TaskID:     taskindex,
		TaskPhrase: m.TaskPhrase,
		IsDone:     false,
	}
	if m.TaskPhrase == models.MapPhrase {
		task.FileName = m.Files[taskindex]
	}
	m.TaskChan <- task
}

func (m *Master) checkTask(taskindex int) {
	timeDuration := time.Now().Sub(m.State[taskindex].StartTime)
	if timeDuration > MAX_TASK_RUNTIME {
		m.addTask(taskindex)
	}
}

func (m *Master) Done() bool {
	ret := false
	finished := true

	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	for key, ts := range m.State {
		switch ts.Status {
		case models.TaskStatusReady:
			finished = false
			m.addTask(key)
		case models.TaskStatusQueue:
			finished = false
		case models.TaskStatusRunning:
			finished = false
			m.checkTask(key)
		case models.TaskStatusFinish:
		case models.TaskStatusError:
			finished = false
			m.addTask(key)
		}
	}

	if finished {
		if m.TaskPhrase == models.MapPhrase {
			m.initReduceTask()
		} else {
			m.IsDone = true
			close(m.TaskChan)
		}
	} else {
		m.IsDone = false
	}

	ret = m.IsDone

	return ret
}

func (m *Master) initReduceTask() {
	fmt.Println("开始进行reduce任务")
	m.TaskPhrase = models.ReducePhrase
	m.IsDone = false
	m.State = make([]TaskState, m.NReduce)
	for index := range m.State {
		m.State[index].Status = models.TaskStatusReady
	}
}
