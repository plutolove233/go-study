package models

type ReqTaskArgs struct {
	// 工作结点状态
	WorkerStatus bool
}

type ReqTaskReply struct {
	Task     Task
	TaskDone bool
}

type ReportTaskArgs struct {
	// 标记工作结点状态
	WorkerStatus bool
	TaskIndex    int
	// 标记任务是否完成
	IsDone bool
}

type ReportTaskReply struct {
	// master响应是否处理成功
	MasterAck bool
}
