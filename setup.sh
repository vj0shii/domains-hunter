#!/bin/bash

# Update the system
sudo apt update
sudo apt upgrade -y

# Install Go
sudo apt install -y golang

# Install required tools
sudo apt install -y subfinder amass httpx nuclei pandoc

# Install additional tools
sudo apt install -y python3 python3-pip
pip3 install dirsearch
go get -u github.com/projectdiscovery/aquatone/cmd/aquatone
go get -u github.com/xm1k3/wfuzz

# Clone dirsearch repository
git clone https://github.com/maurosoria/dirsearch.git

# Install subfinder
GO111MODULE=on go get -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder

# Install amass
GO111MODULE=on go get -v github.com/OWASP/Amass/v3/...

# Install httpx
GO111MODULE=on go get -v github.com/projectdiscovery/httpx/cmd/httpx

# Install nuclei
GO111MODULE=on go get -v github.com/projectdiscovery/nuclei/v2/cmd/nuclei

# Install subfinder's dependencies
GO111MODULE=on go get -v github.com/haccer/subjack
GO111MODULE=on go get -v github.com/subfinder/goaltdns

# Clone nuclei templates
git clone https://github.com/projectdiscovery/nuclei-templates.git

# Compile the Go script
go build script.go

echo "Setup completed!"
