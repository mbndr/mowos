# Mowos

[![Go Report Card](https://goreportcard.com/badge/github.com/mbndr/mowos)](https://goreportcard.com/report/github.com/mbndr/mowos)

**ATTENTION**  
This is very early in development and not ready for production yet!

> Goal of this project is to create a tool which helps monitoring small infrastructures with a file based configuration.

## Explanation / How it should be
Mowos (**MO**nitoring **W**ith**O**ut a **S**erver) is a software which helps you to keep an eye on devices in a local network. I started this project because I wanted to have a quick overview if all necessary services are running on my Raspberry PI home server.

### Agent
On every device you want to monitor an executable `mowos-agent` has to be installed. This agent is used to gather all the data from the host.

### Gatherer
An executable called `mowos-gatherer` runs on the device you're currently working on (e.g. laptop, pc). A built-in web server serves a page with all the current information of your hosts. It's responsible for gathering data from the agents and updating your local web view.

## History
Mowos is only for small infrastructures for home usage. If you want to modify the host configuration on a web frontend, render graphs depending on the history of gathered data or more great features, this is probably not the software you're searching for.

## Encryption
coming soon