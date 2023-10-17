package models

import (
	"distributed-system/raft/raftrpc"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Role string

const (
	Leader    = 0 // 领导
	Follower  = 1 // 追随者
	Candidate = 2 // 候选人
)

type ApplyMsg struct {
	CommandValid bool
	Command      interface{}
	CommandIndex int

	// For 2D:
	SnapshotValid bool
	Snapshot      []byte
	SnapshotTerm  int
	SnapshotIndex int
}

type Raft struct {
	mu            sync.Mutex           // Lock to protect shared access to this peer's state
	peers         []*raftrpc.ClientEnd // RPC end points of all peers
	persister     *Persister           // Object to hold this peer's persisted state
	me            int                  // this peer's index into peers[]
	dead          int32                // set by Kill()=
	lastRecv      time.Time            // 上一次收到其他结点发送消息的时间
	role          int8                 // 当前结点的状态
	currentTerm   int                  // 当前任期
	currentLeader int                  // current leader's index
	votedFor      int                  // 投票给谁
	voteCount     int                  // 当前任期(term)中获得的总票数
	votedChan     chan bool            // 投票通道
	heartbeatChan chan bool
}

func NewRaft(me int, peers []*raftrpc.ClientEnd) *Raft {
	return &Raft{
		peers:         peers,
		me:            me,
		role:          Follower, // 初始时身份是跟随者
		currentTerm:   0,
		currentLeader: -1,
		votedFor:      -1,
		voteCount:     0,
		votedChan:     make(chan bool),
		heartbeatChan: make(chan bool),
	}
}

func (rf *Raft) Start() {
	for {
		switch rf.role {
		case Leader:
			log.Println("send heartbeat message...")
			rf.heartbeat()
		case Follower:
			select {
			case <-rf.heartbeatChan:
				log.Printf("follower %d received heartbeat\n", rf.me)
			case <-time.After(rf.electionTimeout()):
				rf.mu.Lock()
				rf.role = Candidate
				rf.mu.Unlock()
				log.Println("role have been changed: Follower => Candidate")
			}
		case Candidate:
			rf.mu.Lock()
			rf.currentTerm += 1
			rf.votedFor = rf.me
			rf.voteCount = 1
			rf.mu.Unlock()
			rf.Elect()
		}
	}
}

// 检查过程是否被删除
func (rf *Raft) killed() bool {
	z := atomic.LoadInt32(&rf.dead)
	return z == 1
}

// 设置一次信息交换的延时时间，如果在这段时间内没有信息的交换，便需要进行一次选举
func (rf *Raft) electionTimeout() time.Duration {
	return time.Duration(4+rand.Int31n(4)) * time.Second
}

// 选举过程，按论文Section5.2描述的过程来实现
func (rf *Raft) Elect() bool {

	for i, client := range rf.peers {
		if rf.me == i {
			continue
		}
		args := raftrpc.RequestVoteArgs{
			Term:        rf.currentTerm,
			CandidateId: rf.me,
		}
		reply := raftrpc.RequestVoteReply{
			VoteGranted: false,
		}
		go func(client *raftrpc.ClientEnd) {
			ok := client.Call("Raft.RequestVote", &args, &reply)
			if !ok {
				return
			}
			if reply.VoteGranted {
				rf.mu.Lock()
				rf.voteCount += 1
				if rf.voteCount > len(rf.peers)/2+1 {
					rf.votedChan <- true
				}
				rf.mu.Unlock()
			}
			log.Printf("Get vote reply from %s, reply=%+v\n", client.Address, reply)
		}(client)
	}

	select {
	case <-time.After(rf.electionTimeout()):
		log.Println("election timeout, candidate => follower")
		rf.mu.Lock()
		// reset raft statue
		rf.voteCount = 0
		rf.votedFor = -1
		rf.role = Follower
		rf.mu.Unlock()
		return false
	case <-rf.votedChan:
		rf.mu.Lock()
		rf.role = Leader
		rf.currentLeader = rf.me
		rf.mu.Unlock()
		log.Println("leader elected, candiate => leader")
		return true
	}
}

func (rf *Raft) heartbeatInterval() time.Duration {
	return time.Second * 1
}

func (rf *Raft) heartbeat() {

	for i, peer := range rf.peers {
		if i == rf.me {
			continue
		}
		reply := raftrpc.AppendEntitiesReply{}
		// Leaders send periodic heartbeats
		// (AppendEntries RPCs that carry no log entries)
		// to all followers in order to maintain their authority.
		peer.Call("Raft.RequestAppendEntities", &raftrpc.AppendEntitiesArgs{
			Term:     rf.currentTerm,
			LeaderId: rf.me,
		}, &reply)
		rf.lastRecv = time.Now()
	}
	time.Sleep(rf.heartbeatInterval())

}

func (rf *Raft) ticker() {
	for rf.killed() == false {

		// Your code here (2A)
		// Check if a leader election should be started.

		// pause for a random amount of time between 50 and 350
		// milliseconds.
		ms := 50 + (rand.Int63() % 300)
		time.Sleep(time.Duration(ms) * time.Millisecond)
	}
}

// restore previously persisted state.
func (rf *Raft) readPersist(data []byte) {
	if data == nil || len(data) < 1 { // bootstrap without any state?
		return
	}
	// Your code here (2C).
	// Example:
	// r := bytes.NewBuffer(data)
	// d := labgob.NewDecoder(r)
	// var xxx
	// var yyy
	// if d.Decode(&xxx) != nil ||
	//    d.Decode(&yyy) != nil {
	//   error...
	// } else {
	//   rf.xxx = xxx
	//   rf.yyy = yyy
	// }
}

func Make(peers []*raftrpc.ClientEnd, me int,
	persister *Persister, applyCh chan ApplyMsg) *Raft {
	rf := &Raft{}
	rf.peers = peers
	rf.persister = persister
	rf.me = me

	// Your initialization code here (2A, 2B, 2C).

	// initialize from state persisted before a crash
	rf.readPersist(persister.ReadRaftState())

	// start ticker goroutine to start elections
	go rf.ticker()

	return rf
}
