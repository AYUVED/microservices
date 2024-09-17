#!/bin/bash

go build ./cmd/main.go
go mod tidy
git add .
git commit -m "Pushing to git" 
git push origin main