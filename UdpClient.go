package main

// import the necessary packages
import (
	"fmt"
	"net"
)

func main(){
	end_addr := "10.250.1.65:8090" // End point (server) address for the connection. IP: 10.250.1.65 Port: 8090
	s, _ := net.ResolveUDPAddr("udp4", end_addr) // A UDP end point address is created using ResolveUPDAddr
	c, err := net.DialUDP("udp4", nil, s) // Creates a connection to the destination address
	if err != nil {
		fmt.Println(err)
		return
	} 
	fmt.Printf("Connected to UDP Server %s\n", c.RemoteAddr().String())
	text := "My Name is Soma, My roll no is 170010037, What is the sum of numbers in my roll numbers?"
	data := []byte(text+"\n") // Convert the input text into the byte string format for communication
	_, err = c.Write(data) // Send the message to the server
	if err != nil {
		fmt.Println(err)
		return	
	}
	buffer := make([]byte, 1024)
	n, err := c.Read(buffer) // Read the reply from the server in byte array format
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("SERVER RESPONSE: %s\n", buffer[0:n]) // Print out the response to the console.
}