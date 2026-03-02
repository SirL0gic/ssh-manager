package main

import (
	"fmt"
	"time"
)

type SSHProfile struct{
	Name string
	Host string 
	Port int
	Username string 
	Password string
	Keypath string
	isActive bool
}

func connectionString(profile SSHProfile) string {
	return fmt.Sprintf("ssh %s@%s -p %d", profile.Username, profile.Host, profile.Port)
}

func main() {
	version := "SSH Manager v0.0.1"
	date_time := time.Now().Format("2006-01-02 15:04:05")

	fmt.Println(version)
	fmt.Println("Date & Time:", date_time)

	server1 := SSHProfile {
		Name: "server-1",
		Host: "1.1.1.1",
		Port: 22,
		Username: "root",
		Password: "secret123",
		Keypath: "",
		isActive: true,
	}

	fmt.Println(server1)
	fmt.Println(connectionString(server1))
}
