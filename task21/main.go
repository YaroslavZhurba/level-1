package main

import (
	"fmt"
)

//Interface with ConnectWithTCP
type TCPConnection interface {
	ConnectWithTCP()
}

// Struct LinuxServer
// Implements TCPConnection interface
type LinuxServer struct {
	name string
}

// Method of interface TCPConnection
func (server *LinuxServer) ConnectWithTCP() {
	fmt.Println("Server name =",server.name + ". Conneted using TCP")
}

// Struct WindowsServer
// It does not implement TCPConnection interface
// But it has method with similar functionality
type WindowsServer struct {
	name string
}

func (server *WindowsServer) ConnectWithUDP() {
	fmt.Println("Server name =",server.name + ". Conneted using UDP")
}

// Using adapter to implement TCPConnection interface
type WindowsServerAdapter struct {
	WindowsServer
}

func (server *WindowsServerAdapter) ConnectWithTCP() {
	server.ConnectWithUDP()
}

// Client struct
type Client struct {
	name string
}

// Clients's method which using TCPConnection interface
// Make an adapter to use here struct with similar functionality
// but withiout TCPConnection interface implementation
func (c *Client) ConnectToTCPServer(connection TCPConnection) {
	connection.ConnectWithTCP()
}


func main() {
	client := &Client{"Client-1"}

	linServer := &LinuxServer{"Linux-Server-1"}
	winServer := &WindowsServer{"Windows-Server-1"}
	tcpAdapter := &WindowsServerAdapter{*winServer}

	// normal usage of client
	client.ConnectToTCPServer(linServer)

	// Case where uses WindowsServerAdapter
	// to make it work
	client.ConnectToTCPServer(tcpAdapter)
}