@echo off
set GOPATH=%cd%
echo Starting server...
go run httpserv.go
pause > nul