#!/bin/sh

### Set ENV
Out="logger"
Options="-a -installsuffix cgo -o"

### Build
echo -e "go build ${Options} ${Out}"
	 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${Options} ${Out} ./log.go
