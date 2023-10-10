package main

import (
	"distributed-system/map-reduce/models"
	"distributed-system/map-reduce/utils"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"log"
	"net/rpc"
	"os"
	"sort"
)

// 远程rpc调用
func call(rpcName string, args any, reply any) bool {
	c, err := rpc.DialHTTP("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalln("dial master failed, err=", err.Error())
		return false
	}
	defer c.Close()

	err = c.Call(rpcName, args, reply)
	if err != nil {
		log.Println("call rpc service failed, err=", err.Error())
		return false
	}
	return true
}

func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

func reqTask() models.ReqTaskReply {
	args := models.ReqTaskArgs{
		WorkerStatus: true,
	}
	reply := models.ReqTaskReply{}
	ok := call("Master.HandleTaskReq", &args, &reply)
	if !ok {
		log.Println("request task failed...")
	}
	return reply
}

func reportTask(taskid int, isDone bool) models.ReportTaskReply {
	args := models.ReportTaskArgs{
		WorkerStatus: true,
		TaskIndex:    taskid,
		IsDone:       isDone,
	}
	reply := models.ReportTaskReply{}

	if ok := call("Master.HandleTaskReport", &args, &reply); !ok {
		log.Println("report task info failed...")
	}
	return reply
}

func doTask(mapf utils.MapFunc, reducef utils.ReduceFunc, task models.Task) error {
	switch task.TaskPhrase {
	case models.MapPhrase:
		err := DoMapTask(mapf, task.FileName, task.TaskID, task.NReduce)
		return err
	case models.ReducePhrase:
		err := DoReduceTask(reducef, task.FileName, task.TaskID, task.NMap)
		return err
	default:
		log.Println("task phrase exception!")
		return errors.New("task phrase exception")
	}
}

func DoMapTask(mapf utils.MapFunc, filename string, mapTaskIndex int, nreduce int) error {
	log.Println("start mapping...")
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln("open specific file failed, error=", err.Error())
		return err
	}

	kva := mapf(filename, string(content))

	for i := 0; i < nreduce; i++ {
		intermediateName := fmt.Sprintf("mr-%d-%d", mapTaskIndex, i)
		f, _ := os.Create(intermediateName)
		jsonEncoder := json.NewEncoder(f)
		for _, kv := range kva {
			if ihash(kv.Key)%nreduce == i {
				jsonEncoder.Encode(kv)
			}
		}
		f.Close()
	}

	return nil
}

func DoReduceTask(reducef utils.ReduceFunc, filename string, reduceTaskIndex int, nmap int) error {
	log.Println("start reducing...")
	res := make(map[string][]string)
	for i := 0; i < nmap; i++ {
		intermediateName := fmt.Sprintf("mr-%d-%d", i, reduceTaskIndex)
		f, err := os.Open(intermediateName)
		if err != nil {
			log.Fatalln("open intermediate file failed, err=", err.Error())
			return err
		}
		decoder := json.NewDecoder(f)
		for {
			var kv models.KeyValue
			err := decoder.Decode(&kv)
			if err != nil {
				break
			}

			_, ok := res[kv.Key]
			if !ok {
				res[kv.Key] = make([]string, 0)
			}
			res[kv.Key] = append(res[kv.Key], kv.Value)
		}
		f.Close()
	}

	keys := make([]string, 0)
	for k := range res {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	outputFileName := fmt.Sprintf("mr-out-%d", reduceTaskIndex)
	outputFile, _ := os.Create(outputFileName)

	for _, v := range keys {
		cnt := reducef(v, res[v])
		fmt.Fprintf(outputFile, "%s %s\n", v, cnt)
	}
	outputFile.Close()
	return nil
}

func Worker(mapf utils.MapFunc, reducef utils.ReduceFunc) {
	for {
		// 请求任务
		reply := models.ReqTaskReply{}
		reply = reqTask()
		fmt.Printf("%+v\n", reply)
		if reply.TaskDone {
			break
		}
		// 执行任务
		err := doTask(mapf, reducef, reply.Task)
		if err != nil {
			reportTask(reply.Task.TaskID, false)
			continue
		}
		// 反馈任务
		reportTask(reply.Task.TaskID, true)
	}
}
