//UDP-client.go
package main

import (
	"fmt"
	"net"
)

func main() {
	//adres tanımlanması
	serverAddr := net.UDPAddr{
		Port: 12345,
		IP:   net.ParseIP("127.0.0.1"),
	}
       //adrese bağlantı açılması
	conn, err := net.DialUDP("udp", nil, &serverAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
      // veri okunması
	go func() {
		buffer := make([]byte, 1024)
		for {
			n, _, err := conn.ReadFromUDP(buffer)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Received: %s\n", string(buffer[:n]))
		}
	}()
       //kullanıcıdan mesaj alınması
	for {
		var msg string
		fmt.Print("Mesaj girin: ")
		fmt.Scanln(&msg)

		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
		}
	}
}
