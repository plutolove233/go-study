// Package distributed
/*
@Coding : utf-8
@time : 2022/8/5 9:18
@Author : yizhigopher
@Software : GoLand
*/
package main

import (
	"errors"
	"net"
	"sync"
	"time"
)

const (
	bitLenTime      = 39
	bitLenSequence  = 8
	bitLenMachineID = 63 - bitLenTime - bitLenSequence
)

type Settings struct {
	StartTime      time.Time
	MachineID      func() (uint16, error)
	CheckMachineID func(uint16) bool
}

type Sonyflake struct {
	mutex       *sync.Mutex
	startTime   int64
	elapsedTime int64 // 自设置的startTime开始以来，总共经过了多少个单位时间
	sequence    uint16
	machineID   uint16
}

func NewSonyflake(st Settings) *Sonyflake {
	sf := new(Sonyflake)
	sf.mutex = new(sync.Mutex)
	sf.sequence = uint16(1<<bitLenSequence - 1)

	if st.StartTime.After(time.Now()) {
		return nil
	}
	if st.StartTime.IsZero() {
		sf.startTime = toSonyflakeTime(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC))
	} else {
		sf.startTime = toSonyflakeTime(st.StartTime)
	}
	var err error
	if st.MachineID == nil { // 如果不使用自定义的获取实例机器id号的函数的话，那么默认machineID就是私有IP地址的低16位
		sf.machineID, err = lower16BitPrivateIP()
	} else {
		sf.machineID, err = st.MachineID()
	}
	if err != nil || (st.CheckMachineID != nil && !st.CheckMachineID(sf.machineID)) {
		return nil
	}

	return sf
}

func (sf *Sonyflake) NextID() (uint64, error) {
	const maskSequence = uint16(1<<bitLenSequence - 1) // 掩码号

	sf.mutex.Lock() // 上锁
	defer sf.mutex.Unlock()

	current := currentElapsedTime(sf.startTime) // 自算法启动以来经过了多少个单位时间, 即当前时间和设置的startTime时间差值换算以后等于多少个单位时间
	if sf.elapsedTime < current {               // 即距离上一次生成id的时候已经经过了10ms以上了，那么就重置序列号
		sf.elapsedTime = current
		sf.sequence = 0
	} else { // sf.elapsedTime >= current						// 如果elapsedTime与current相等的话，就意味着还在同一个单位时间内，即在10 ms内又需要生成一次id; 如果是elapsedTime大于current的情况的话，就说明发生了时间回拨的问题
		sf.sequence = (sf.sequence + 1) & maskSequence // 序列号+1
		if sf.sequence == 0 {                          // 在单位时间内达到了最大可生成序列号的限制
			sf.elapsedTime++                     // 将elapsedTime设为下一个周期
			overtime := sf.elapsedTime - current // 计算时长，睡眠等待
			time.Sleep(sleepTime((overtime)))
			//
			// 需要详细说明一下，发生时间回拨的时候。
			// 继续上一周期的sequence进行id生成，当sequence达到上限的时候，
			// 才会睡眠，直到系统时钟走到回拨前的elapsedTime所代表时间点。
			//
			// 比如：sequence是101的时候，时间回拨了3秒。如果在这3秒内，
			// 不断的生成id，让sequence达到了255，假设生成154次id花费的时间是8毫秒，
			// 所以就会睡眠2992毫秒；但如果在这3秒内，sequence只到了200，那么就不会发生睡眠
			//
			// 巧妙的是，就算发生了时间回拨也不会生成重复id，因为在时间回拨区间内，
			// elapsedTime是固定的，能决定id的值就是sequence了。如果sequence溢出
			// 就睡眠使elapsedTime进入下一周期，如果sequence没溢出的话，系统时钟
			// 这时也恢复正常了，elapsedTime自然也被刷新了，那更不会生成重复id了。

		}
	}

	return sf.toID()
}

const sonyflakeTimeUnit = 1e7

func toSonyflakeTime(t time.Time) int64 {
	return t.UTC().UnixNano() / sonyflakeTimeUnit
}

func currentElapsedTime(startTime int64) int64 {
	return toSonyflakeTime(time.Now()) - startTime
}

func sleepTime(overtime int64) time.Duration {
	return time.Duration(overtime)*10*time.Millisecond -
		time.Duration(time.Now().UTC().UnixNano()%sonyflakeTimeUnit)*time.Nanosecond
}

func (s *Sonyflake) toID() (uint64, error) {
	if s.elapsedTime >= 1<<bitLenTime {
		return 0, errors.New("over the time limit")
	}
	return uint64(s.elapsedTime)<<(bitLenSequence+bitLenMachineID) |
		uint64(s.sequence)<<bitLenMachineID |
		uint64(s.machineID), nil
}

func privateIPv4() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if isPrivateIPv4(ip) {
			return ip, nil
		}
	}

	return nil, errors.New("no private ip address")
}

func isPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}

func lower16BitPrivateIP() (uint16, error) {
	ip, err := privateIPv4()
	if err != nil {
		return 0, err
	}

	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}

func Decompose(id uint64) map[string]uint64 {
	const maskSequence = uint64((1<<bitLenSequence - 1) << bitLenMachineID)
	const maskMachineID = uint64(1<<bitLenMachineID - 1)

	msb := id >> 63
	time := id >> (bitLenSequence + bitLenMachineID)
	sequence := id & maskSequence >> bitLenMachineID
	machineID := id & maskMachineID
	return map[string]uint64{
		"id":         id,
		"msb":        msb, // msb我没搞懂是什么的缩写
		"time":       time,
		"sequence":   sequence,
		"machine-id": machineID,
	}
}
