echo off

rem build front page
call yarn build

rem build server app
cd srv
set GOOS=linux
set GOARCH=amd64
go build -o ../dist 
cd ..