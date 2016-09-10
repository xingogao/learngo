@echo off
if "%1" equ "" (
    gofmt --help
    goto :end
)
gofmt "%1"
:end
pause>nul
exit