@echo off

echo Generating app.json
go build github.com/Tnze/CoolQ-Golang-SDK/tools/cqcfg
go generate
IF ERRORLEVEL 1 pause
