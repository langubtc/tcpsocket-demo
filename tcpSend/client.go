package main

import (
	"fmt"
	"net"
	"tcpclient/proto"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}

	defer conn.Close()

	msg := `stop|hello`
	data, err := proto.Encode(msg)
	if err != nil {
		fmt.Println("encode msg failed, err:", err)
		return
	}
	conn.Write(data)

	select {

	}

}
