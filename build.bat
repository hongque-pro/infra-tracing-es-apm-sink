@echo off
echo Compile for windows ...
del /f /s /q %~dp0\output\*.*
pushd %~dp0
go build -ldflags="-w -s" -o "./output/tracing-sink-win.exe" -i ./main

echo.
set /p linux=compile for linux? [y/n]:
if "%linux%"=="y" (
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
echo Compile for linux amd64 ...
go build -ldflags="-w -s" -o "./output/tracing-sink-linux" -i ./main
)

echo.
set /p linux=Compile succeed ! press any key to exit .
