package models

// 任务状态描述
const (
	TaskStatusReady   = "ready"
	TaskStatusQueue   = "queue"
	TaskStatusRunning = "running"
	TaskStatusError   = "error"
	TaskStatusFinish  = "finish"
)

// 任务阶段描述
const (
	MapPhrase    = "map"    // map phrase
	ReducePhrase = "reduce" // reduce phrase
)

type Task struct {
	// 任务阶段描述
	TaskPhrase string
	// map任务个数
	NMap int
	// reduce任务个数
	NReduce int
	// 任务id
	TaskID   int
	FileName string
	IsDone   bool
}
