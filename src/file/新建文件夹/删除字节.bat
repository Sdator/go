@echo off
for /r %%i in (*.bin) do dx20 -f "%%i"
pause
exit