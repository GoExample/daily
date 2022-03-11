package main

func main() {
	// Example 1
	//PortMultiplexing1()

	// Example 2
	go PortMultiplexing2("127.0.0.1:1234")
	go PortMultiplexing2("0.0.0.0:1234")
	select {}
}
