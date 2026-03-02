package main

import (
	"fmt"
	"time"
	"log"
	"os"
	"encoding/json"
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
	data,err := os.ReadFile("profiles.json")
	if err != nil {
		log.Fatal("Error reading profiles.json", err)
	}

	var profiles []SSHProfile
	
	err = json.Unmarshal(data, &profiles)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
	}

	return profiles
}

func main() {
	version := "SSH Manager v0.0.1"
	date_time := time.Now().Format("2006-01-02 15:04:05")

	fmt.Println(version)
	fmt.Println("Date & Time:", date_time)

}
