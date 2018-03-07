#!/bin/sh

### Set ENV
Out="srv"
Options="-a -installsuffix cgo -o"

### Build
echo -e "go build ${Options} ${Out}"
	 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${Options} ${Out} ./
