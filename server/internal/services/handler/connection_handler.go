package handler

import (
	"log"
	"net"
	"strings"
	"time"
	"wow/server/internal/services/generator"
	"wow/server/internal/services/quote"
	"wow/server/internal/services/verifier"
	//"wow/server/internal/services/verifier"
)

type ConnectionHandler struct {
	Difficulty   int
	timeout      time.Duration
	QuoteService *quote.Service
}

func NewConnectionHandler(difficulty int, timeout time.Duration, service *quote.Service) *ConnectionHandler {
	return &ConnectionHandler{
		Difficulty:   difficulty,
		timeout:      timeout,
		QuoteService: service,
	}
}

func (ch *ConnectionHandler) HandleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("Failed to close connection: ", err.Error())
		}
	}(conn)

	challenge := generator.GenerateChallenge(ch.Difficulty)
	if err := ch.sendChallenge(conn, challenge); err != nil {
		log.Println("Error sending challenge: ", err.Error())
		return
	}

	nonce, err := ch.readOnce(conn)
	if err != nil {
		log.Println("Error reading nonce: ", err.Error())
		return
	}

	if verifier.VerifySolution(challenge, nonce) {
		if err := ch.sendQuote(conn); err != nil {
			log.Println("Error sending quote: ", err.Error())
		}
	} else {
		if err := ch.sendInvalidSolution(conn); err != nil {
			log.Println("Error sending invalid solution: ", err.Error())
		}
	}

}

func (ch *ConnectionHandler) sendChallenge(conn net.Conn, challenge string) error {
	err := conn.SetWriteDeadline(time.Now().Add(ch.timeout))
	if err != nil {
		return err
	}
	_, err = conn.Write([]byte(challenge + "\n"))
	if err != nil {
		return err
	}
	return err
}

func (ch *ConnectionHandler) readOnce(conn net.Conn) (string, error) {
	buffer := make([]byte, 1024)
	err := conn.SetReadDeadline(time.Now().Add(ch.timeout))
	if err != nil {
		return "", err
	}
	n, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(buffer[:n])), nil
}

func (ch *ConnectionHandler) sendQuote(conn net.Conn) error {
	randomQuote, err := ch.QuoteService.GetRandomQuote()
	if err != nil {
		return err
	}
	err = conn.SetWriteDeadline(time.Now().Add(ch.timeout))
	if err != nil {
		return err
	}
	_, err = conn.Write([]byte(randomQuote + "\n"))

	return err

}

func (ch *ConnectionHandler) sendInvalidSolution(conn net.Conn) error {
	err := conn.SetWriteDeadline(time.Now().Add(ch.timeout))
	if err != nil {
		return err
	}
	_, err = conn.Write([]byte("Invalid solution\n"))
	return err
}
