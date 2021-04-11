package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os/exec"
	"tcpclient/proto"
	"time"
)
func doWork(conn net.Conn){

	for {
		msga := `hello`
		data, err := proto.Encode(msga)
		if err != nil {
			fmt.Println("encode msg failed, err:<<<<<<<<<", err)
			break
		}
		conn.Write(data)
		reader := bufio.NewReader(conn)

		msg, err := proto.Decode(reader)
		if err != nil {
			fmt.Println("encode msg failed, err:------", err)
			break
		}

		if msg == "stop"{
			cmd := exec.Command("df","-h")
			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatalf("cmd.Run() failed with %s\n", err)
			}
			fmt.Printf("combined out:\n%s\n", string(out))
		} else if msg == "start"{
			cmd := exec.Command("uname","-a")
			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatalf("cmd.Run() failed with %s\n", err)
			}
			fmt.Printf("combined out:\n%s\n", string(out))
		}

	}

}

func main() {
	for {
		conn, err := net.Dial("tcp", "127.0.0.1:30000")
		if err != nil {
			fmt.Println("服务器连接失败...", err)
		}else {
			fmt.Println("连接成功....")
			defer conn.Close()
			doWork(conn)
		}
		time.Sleep(3 *time.Second) //断线间隔3秒重连
		fmt.Println("连接断开，正在重连...")
	}
}
