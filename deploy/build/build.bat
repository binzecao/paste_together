@echo off

set GOPROXY=https://goproxy.io
set CGO_ENABLED=0
set GOOS=windows
set CURRENT_PATH="%cd%"
set PROJECT_ROOT_PATH=%CURRENT_PATH%\..\..
set BUILD_TARGET_PATH=%PROJECT_ROOT_PATH%\bin\paste_together.exe

cd "%PROJECT_ROOT_PATH%\src\aaa.com\paste_together"
go build -tags netgo -a -o "%BUILD_TARGET_PATH%" aaa.com/paste_together

xcopy "%PROJECT_ROOT_PATH%\src\aaa.com\paste_together\template" "%PROJECT_ROOT_PATH%\bin\template" /y /e /i

cd "%CURRENT_PATH%"

echo "Build finish..."

rem pause
