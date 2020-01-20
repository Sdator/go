@echo off
pushd %~dp0
set "用户数据目录=.vscode\Code"
rem code --user-data-dir %用户数据目录% .
code .
exit

