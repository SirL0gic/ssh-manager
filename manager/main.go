package main

import (
	"fmt"
	"time"
)

func main() {
	version := "SSH Manager v0.0.1"
	date_time := time.Now().Format("2006-01-02 15:04:05")
	
	fmt.Println(version)
	fmt.Println("Date & Time:", date_time)
}
