package raftrpc

type AppendEntitiesArgs struct {
	Term         int // leader's term
	LeaderId     int
	PrevLogIndex int
	PrevLogTerm  int // term of prevLogIndex entry
	// Entries []LogEntry
	LeaderCommit int // leader's commit index
}

type AppendEntitiesReply struct {
	Term    int  // current term, for leader to update itself
	Success bool // true if follower contained entry matching prevLogIndex and prevLogTerm
}

type RequestVoteArgs struct {
	Term         int
	CandidateId  int
	LastLogIndex int
	LastLogTerm  int
}

type RequestVoteReply struct {
	Term        int  // current term, for candidate to update itself
	VoteGranted bool //true means candidate received vote
}
