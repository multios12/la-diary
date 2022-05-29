#!/bin/sh

yarn build
cp ./build/* ./srv/static/ -R
find ./build -delete

cd srv
GOOS=linux
GOARCH=amd64
go build -ldflags="-s -w" -trimpath -o ../dist/
cd ..