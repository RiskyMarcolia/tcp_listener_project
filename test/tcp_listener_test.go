package test

import (
	"net"
	"testing"
	"time"

	"github.com/RiskyMarcolia/tcp_listener_project/app"
)

func TestTCPListener(t *testing.T) {
	host := "127.0.0.1"
	port := 12345

	go app.TCPListener(host, port)

	conn, err := net.DialTimeout("tcp", host+":12345", 5*time.Second)
	if err != nil {
		t.Errorf("Could not connect to server: %v", err)
		return
	}
	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	message := []byte("Hello, TCP!")
	_, err = conn.Write(message)
	if err != nil {
		t.Errorf("Error writing to server: %v", err)
	}

	response := make([]byte, 1024)
	_, err = conn.Read(response)
	if err != nil {
		t.Errorf("Error reading from server: %v", err)
	}

	if string(response) != string(message) {
		t.Errorf("Unexpected response. Expected: %s, Got: %s", string(message), string(response))
	}
}
