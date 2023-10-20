package models

import (
	"distributed-system/raft/raftrpc"
	"log"
	"net/http"
	"net/rpc"
	"time"
)

// run a goroutine so that can serve rpc request from other clientend
func (rf *Raft) RunRPCService(address string) {
	rpc.Register(rf)
	rpc.HandleHTTP()

	go func() {
		err := http.ListenAndServe(address, nil)
		if err != nil {
			log.Fatalln("start rpc service failed, err=", err.Error())
		}
	}()

	log.Println("rpc service is listening on ", address)
}

func (rf *Raft) RequestVote(args *raftrpc.RequestVoteArgs, reply *raftrpc.RequestVoteReply) error {
	reply.VoteGranted = false
	rf.lastRecv = time.Now()
	if args.Term < rf.currentTerm {
		reply.Term = rf.currentTerm
		reply.VoteGranted = false
		return nil
	}
	if args.Term > rf.currentTerm {
		rf.mu.Lock()
		rf.currentTerm = args.Term
		rf.role = Follower
		rf.votedFor = -1
		rf.mu.Unlock()
		return nil
	}
	if rf.votedFor == -1 {
		rf.votedFor = args.CandidateId
		reply.VoteGranted = true
	}
	return nil
}

func (rf *Raft) AppendEntities(args *raftrpc.AppendEntitiesArgs, reply *raftrpc.AppendEntitiesReply) error {
	if args.Term < rf.currentTerm {
		reply.Term = rf.currentTerm
		reply.Success = true
	}
	if args.Term > rf.currentTerm {
		rf.mu.Lock()
		rf.currentTerm = args.Term
		rf.role = Follower
		log.Println("change to follower")
		rf.votedFor = -1
		rf.currentLeader = args.LeaderId
		rf.mu.Unlock()
	}
	rf.mu.Lock()
	rf.lastRecv = time.Now()
	rf.mu.Unlock()
	reply.Term = rf.currentTerm
	rf.heartbeatChan <- true
	return nil
}
