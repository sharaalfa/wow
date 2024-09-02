package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"wow/pkg/config"
	"wow/pkg/generator"
)

func main() {
	config.InitConfig()

	cfg := config.GetConfig()

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port))

	if err != nil {
		log.Fatalf("Error connection to server: %v", err)
		return
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Failed to close client: %v", err)
		}
	}(conn)

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatalf("Error reading challenge: %v", err)
		return
	}

	challenge := strings.TrimSpace(string(buffer[:n]))
	nonce := solveChallenge(challenge)

	_, err = conn.Write([]byte(nonce))
	if err != nil {
		log.Fatalf("Error writing nonce: %v", err)
		return
	}

	quoteBuffer := make([]byte, 1024)
	n, err = conn.Read(quoteBuffer)
	if err != nil {
		log.Fatalf("Error reading quote: %v", err)
		return
	}

	quote := strings.TrimSpace(string(quoteBuffer[:n]))
	log.Printf("Received quote: %s", quote)

}

func solveChallenge(challenge string) string {
	parts := strings.Split(challenge, ":")
	if len(parts) != 2 {
		log.Fatal("Invalid challenge format")
	}

	data := parts[0]
	difficulty, err := strconv.Atoi(parts[1])

	if err != nil {
		log.Fatalf("Invalid difficulty: %v", err)
	}

	for i := 0; ; i++ {
		nonce := strconv.Itoa(i)
		hash := sha256.Sum256([]byte(data + nonce))
		if generator.CheckHashPrefix(hash[:], difficulty) {
			return nonce
		}
	}
}
