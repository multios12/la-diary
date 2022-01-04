echo off

rem build front page
call yarn build

rem build server app
cd srv
go build -o ../dist
cd ..