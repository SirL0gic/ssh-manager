package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

type SSHProfile struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Keypath  string `json:"keypath"`
	IsActive bool   `json:"isActive"`
}

func connectionString(profile SSHProfile) string {
	return fmt.Sprintf("ssh %s@%s -p %d", profile.Username, profile.Host, profile.Port)
}

func loadProfiles() []SSHProfile {
	data, err := os.ReadFile("profile.json")
	if err != nil {
		log.Fatal("Error reading profile.json:", err)
	}

	var profiles []SSHProfile
	err = json.Unmarshal(data, &profiles)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
	}

	return profiles
}

func connectSSH(profile SSHProfile) {
	config := &ssh.ClientConfig{
		User: profile.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(profile.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	address := fmt.Sprintf("%s:%d", profile.Host, profile.Port)
	fmt.Println("Connecting to", address, "...")

	client, err := ssh.Dial("tcp", address, config)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer client.Close()

	// Start an interactive session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session:", err)
	}
	defer session.Close()

	// Connect your terminal to the remote server
	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	// Request a terminal
	session.RequestPty("xterm", 80, 40, ssh.TerminalModes{})

	// Start shell
	session.Shell()
	session.Wait()
}

func main() {
	version := "SSH Manager v0.0.1"
	date_time := time.Now().Format("2006-01-02 15:04:05")

	fmt.Println(version)
	fmt.Println("Date & Time:", date_time)
	fmt.Println()

	profiles := loadProfiles()

	for i, profile := range profiles {
		number := i + 1
		name := profile.Name
		fmt.Printf("%d. %s", number, name)
	}

	fmt.Println()
	fmt.Print("Select a profile: ")

	var choice int
	fmt.Scan(&choice)

	if choice < 1 || choice > len(profiles) {
		log.Fatal("Invalid choice")
	}

	selected := profiles[choice-1]
	connectSSH(selected)
}