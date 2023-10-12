package models

import (
	"math/rand"
	"sync"
	"time"
)

type Role string

const (
	Leader    = "leader"    // 领导
	Follower  = "follower"  // 追随者
	Candidate = "Candidate" // 候选人
)

type Raft struct {
	mu          sync.Mutex   // Lock to protect shared access to this peer's state
	peers       []*ClientEnd // RPC end points of all peers
	persister   *Persister   // Object to hold this peer's persisted state
	me          int          // this peer's index into peers[]
	dead        int32        // set by Kill()=
	lastRecv    time.Time
	role        Role
	currentTerm int
	votedFor    int
}

type Persister struct {
	mu        sync.Mutex
	raftstate []byte
	snapshot  []byte
}

// 设置一次信息交换的延时时间，如果在这段时间内没有信息的交换，便需要进行一次选举
func (r *Raft) electionTimeout() time.Duration {
	return time.Duration(150+rand.Int31n(150)) * time.Millisecond
}
