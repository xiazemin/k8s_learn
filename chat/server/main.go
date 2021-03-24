//chatroom.go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//一个聊天服务器demo
func main() {
	ip := os.Args[1]
	listener, err := net.Listen("tcp", ip)
	if err != nil {
		fmt.Println("listener start fail!")
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept fail!")
			continue
		}
		go handleConn(conn)
	}
}

//广播器
type client chan<- string //对外发送消息的通道
var (
	entering = make(chan client) //新客户进入 用来传递该用户的chan string
	leaving  = make(chan client) //新用户离开
	messages = make(chan string) //客户消息广播
)

func broadcaster() {
	clients := make(map[client]bool) //所有连接的客户端	每个cli是一个chan string，就是handleConn里的ch
	for {
		select {
		case msg := <-messages:
			//把所有接收的消息广播给所有客户
			fmt.Println(msg)
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//每个客户端维持的连接
func handleConn(conn net.Conn) {
	ch := make(chan string)   //创建客户端ch 用于从select接收消息并传给remote
	go clientWriter(conn, ch) //发送至客户端

	who := conn.RemoteAddr().String()
	ch <- "Hello your IP=" + who + "\n Welcome come to zzp chat room!"
	entering <- ch //ch注册到clients
	ch <- "in put you name:"
	input := bufio.NewScanner(conn) //接收remote消息
	if input.Scan() {
		who = input.Text()
	}
	messages <- who + " has arrived!" //通知其他client
	for input.Scan() {
		str := input.Text()
		if str == "\n" {
			continue
		}
		messages <- who + ": " + str
	}

	leaving <- ch //注销
	messages <- who + " has left!"
	conn.Close()
}

//发送消息到客户端
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
