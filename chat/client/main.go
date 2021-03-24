//netcat.go
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	ip := os.Args[1]
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn) //接收程序
	mustCopy(conn, os.Stdin)     //输入程序
	conn.Close()
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println(err)
	}
}
