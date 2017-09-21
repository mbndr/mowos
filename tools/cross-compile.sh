#!/bin/bash
# ONLY FOR LOCAL TESTING ATM

go build -o /vagrant/mowos-agent cmd/agent.go
go build -o /vagrant/mowos-monitor cmd/monitor.go
