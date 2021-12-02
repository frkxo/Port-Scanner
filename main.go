package main

import (
	"fmt"
	"github.com/frkxo/Port-Scanner/port"
)

func main() {
	fmt.Println("Ports are scanning!")
	results := port.InitialScan("localhost")
	fmt.Println(results)
}
