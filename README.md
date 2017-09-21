# Mowos

![goreportcard](https://goreportcard.com/badge/github.com/mbndr/mowos)

**ATTENTION**  
This is very very very NOT working code!  
It's recently under development!  
Things may appear or disappear!  
All text below this is coming from the future and just tells how it should work and be!

> Goal of this project is to create a tool which helps monitoring small infrastructures with a file based configuration.

## Explanation / How it should be
Mowos (**MO**nitoring **W**ith**O**ut a **S**erver) is a software which helps you to keep an eye on devices in a local network. I started this project because I wanted to have a quick overview if all necessary services are running on my Raspberry PI home server.

### Agent
On every device you want to monitor an executable `mowos-agent` has to be installed. This agent is used to gather all the data from the host.

### Monitor
An executable called `mowos-monitor` runs on the device you're currently working on (e.g. laptop, pc). A built-in web server serves a page with all the current information of your hosts. It's responsible for gathering data from the agents and updating your local web view.

## History
Mowos is only for small infrastructures for home usage. If you want to modify the host configuration on a web frontend, render graphs depending on the history of gathered data or more great features, this is probably not the software you're searching for.

## Encryption
The communication between the agents and the monitor is currently not encrypted because I wanted to minimize complexity at the beginning of the project. It's not priority number one because Mowos was written to monitor only small infrastructure members in a local network.

Encryption should be definitely be implemented in the future because basically everyone could send a request TCP packet to a device running an `mowos-agent` and would get data.
