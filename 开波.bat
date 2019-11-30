@echo off
pushd %~dp0
set "用户数据目录=.vscode\Code"
code --user-data-dir %用户数据目录% .

exit

