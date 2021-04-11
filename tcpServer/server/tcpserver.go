package main


import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"tcpdemo/proto"
)

var clientMap map[string]net.Conn = make(map[string]net.Conn)

func process(conn net.Conn) {
	defer conn.Close()
	fmt.Println(conn.RemoteAddr())
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		msg_str := strings.Split(string(msg), "|")

		fmt.Println(msg_str[0])
		switch msg_str[0] {

		case "stop":
			fmt.Println(msg_str,msg_str[0],msg_str[1])
			fmt.Println(conn.RemoteAddr(), "的用户名是", msg_str[1])
			for user, message := range clientMap {
				//向除自己之外的用户发送加入聊天室的消息
				fmt.Println(user,msg_str[1],"<<<<<<")
				if user == msg_str[1] {
					data, err := proto.Encode("stop")
					if err != nil {
						fmt.Println("encode msg failed, err:", err)
						return
					}
					//fmt.Println("<<<<<<<<ok",data)
					message.Write(data)
				}
			}

		case "start":
			for user, message := range clientMap {
				if user == msg_str[1] {
					data, err := proto.Encode("start")
					if err != nil {
						fmt.Println("encode msg failed, err:", err)
						return
					}
					message.Write(data)
				}
			}

		default:
			if err == io.EOF {
				return
			}

			if err != nil {
				fmt.Println("decode msg failed, err:",conn.RemoteAddr(), err)
				return
			}

			fmt.Println("客户端已链接，主机为",conn.RemoteAddr(), msg)
			clientMap[msg_str[0]] = conn


			}




	}
}


func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}

}


