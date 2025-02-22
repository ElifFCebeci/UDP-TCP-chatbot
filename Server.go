//UDP-Server.go
package main

import (
	"fmt"
	"net"
	"strings"
)
//UDP oluşturulması
func main() {
	addr := net.UDPAddr{
		Port: 12345,
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
    // verilerin buffer a yazılması
	buffer := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			continue
		}
		message := string(buffer[:n])
		fmt.Printf("Received: %s from %s\n", message, clientAddr)

		response := processMessage(message)
		_, err = conn.WriteToUDP([]byte(response), clientAddr)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func processMessage(msg string) string {
	// Basit mesaj işleme mantığı
	if strings.TrimSpace(msg) == "Merhaba" {
		return "Merhaba! Size nasıl yardımcı olabilirim?"
	}
	return "Anlamadım, lütfen tekrar edin."
}
