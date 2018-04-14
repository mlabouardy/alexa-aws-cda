#!/bin/bash

echo "Building Binary"
GOOS=linux GOARCH=amd64 go build -o main *.go

echo "Creating deployment package"
zip deployment.zip main

echo "Cleaning up"
rm main