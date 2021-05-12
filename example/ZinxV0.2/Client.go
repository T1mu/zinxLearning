package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("[Client Start...]")
	conn, err := net.Dial("tcp", "127.0.0.1:9573")
	if err != nil {
		fmt.Println("[Dial Error]:", err)
	}
	// 循环写入数据
	for {
		_, err = conn.Write([]byte("[Hello Zinxv0.2...]\n"))
		if err != nil {
			return
		}
		// 读数据
		buff := make([]byte, 512)
		_, err := conn.Read(buff)
		if err != nil {
			return
		}
		fmt.Printf("%s",buff)
		time.Sleep(1 * time.Second)
	}

}
