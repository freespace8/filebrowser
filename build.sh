#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"
upx -9 filebrowser