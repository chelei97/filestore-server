package main

import (
	"fmt"
	"net"
	"time"
)

func main(){
	fmt.Println("Client Test ... start")
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp4", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("dial ", err)
		return
	}

	for {
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			fmt.Println("Write ", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read ", err)
			return
		}

		fmt.Println(buf[:cnt], cnt)

		time.Sleep(1 * time.Second)
	}
}
