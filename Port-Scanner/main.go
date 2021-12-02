package main

import (
	"fmt"
	"port/port"
)

func main() {
	fmt.Println("Ports are scanning!")
	results := port.InitialScan("localhost")
	fmt.Println(results)
}
