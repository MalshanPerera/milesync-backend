package main

import "jira-for-peasents/server"

func main() {
	server := server.NewServer()
	if server != nil {
		server.Start()
	}
}
