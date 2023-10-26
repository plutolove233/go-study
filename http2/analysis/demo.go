package main

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"golang.org/x/net/http2/hpack"
	"time"
)

const _key = "../cert/localhost-key.pem"
const _crt = "../cert/localhost.pem"

const (
	CONN_PREFACE   = 24
	FRAME_HEAD_LEN = 9
)

type FrameType uint8

const (
	FrameData         FrameType = 0x0
	FrameHeaders      FrameType = 0x1
	FramePriority     FrameType = 0x2
	FrameRSTStream    FrameType = 0x3
	FrameSettings     FrameType = 0x4
	FramePushPromise  FrameType = 0x5
	FramePing         FrameType = 0x6
	FrameGoAway       FrameType = 0x7
	FrameWindowUpdate FrameType = 0x8
	FrameContinuation FrameType = 0x9
)

func main() {
	certs := []tls.Certificate{}
	crt, err := tls.LoadX509KeyPair(_crt, _key)
	if err != nil {
		fmt.Println("load err")
	}
	certs = append(certs, crt)
	tlsConfig := &tls.Config{}
	tlsConfig.Certificates = certs
	tlsConfig.Time = time.Now
	tlsConfig.Rand = rand.Reader
	tlsConfig.NextProtos = append(tlsConfig.NextProtos, "h1")
	tlsConfig.NextProtos = append(tlsConfig.NextProtos, "h2")
	tlsConfig.NextProtos = append(tlsConfig.NextProtos, "h2c")

	lis, err := tls.Listen("tcp", "0.0.0.0:5555", tlsConfig)
	if err != nil {
		fmt.Println("listen err")
	}

	conn, _ := lis.Accept()
	defer conn.Close()
	fmt.Println(conn.RemoteAddr().String())
	connectionPreface := make([]byte, CONN_PREFACE)
	conn.Read(connectionPreface)
	fmt.Printf("connection-preface=\n%s", string(connectionPreface))
	if string(connectionPreface) == "PRI * HTTP/2.0\r\n\r\nSM\r\n\r\n" {
		// we should send setting frame
		setting_frame := []byte{0, 0, 0, 4, 0, 0, 0, 0, 0}
		conn.Write(setting_frame)
	}
	for {
		buf := make([]byte, FRAME_HEAD_LEN)
		conn.Read(buf)
		length := (uint32(buf[0])<<16 | uint32(buf[1])<<8 | uint32(buf[2]))
		frameType := FrameType(buf[3])
		switch frameType {
		case FrameSettings:
			// just read
			payload := make([]byte, length)
			conn.Read(payload)
			fmt.Println("Setting Frame PayLoad=", payload)
		case FrameHeaders:
			payload := make([]byte, length)
			conn.Read(payload)
			decoder := hpack.NewDecoder(4096, nil)
			full, _ := decoder.DecodeFull(payload)
			fmt.Println("Head Frame PayLoad=", full)

			resp := []byte{0, 0, 1, 0, 0, 0, 0, 0, 0, 2}
			conn.Write(resp)
		case FrameData:
			payload := make([]byte, length)
			conn.Read(payload)
			fmt.Println("Data Frame PayLoad=", payload)
		case FramePriority:
			payload := make([]byte, length)
			conn.Read(payload)
			fmt.Println("Priority Frame PayLoad=", payload)
		case FrameRSTStream:
			payload := make([]byte, length)
			conn.Read(payload)
			fmt.Println("RST Frame PayLoad=", payload)
		case FramePushPromise:
			payload := make([]byte, length)
			conn.Read(payload)
			fmt.Println("Push Promise Frame PayLoad=", payload)
		case FramePing:
			payload := make([]byte, length)
			conn.Read(payload)
			fmt.Println("Ping Frame PayLoad=", payload)
		case FrameGoAway:
			payload := make([]byte, length)
			conn.Read(payload)
			fmt.Println("Go Away Frame PayLoad=", payload)
		case FrameWindowUpdate:
			payload := make([]byte, length)
			conn.Read(payload)
			fmt.Println("Window Update Frame PayLoad=", payload)
		case FrameContinuation:
			payload := make([]byte, length)
			conn.Read(payload)
			fmt.Println("Continuation Frame PayLoad=", payload)
		}
	}

	//for {
	//	rBuf := bufio.NewReader(conn)
	//	frame, err := http2.ReadFrameHeader(rBuf)
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	// Decode the frame
	//	switch frame.Header().Type {
	//	case http2.FrameHeaders:
	//		fmt.Println("StreamID", frame.Header())
	//		var headPayload []byte
	//		length := frame.Header().Length
	//		headPayload = make([]byte, length)
	//
	//		_, err := io.ReadFull(conn, headPayload[:length])
	//		if err != nil {
	//			fmt.Println("read header payload failed, err=", err)
	//			return
	//		}
	//		fmt.Println(string(headPayload))
	//
	//	case http2.FrameData:
	//		fmt.Println("data", frame.Header())
	//	}
	//}
}
