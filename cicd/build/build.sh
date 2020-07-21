#!/usr/bin/bash

GOPROXY=https://goproxy.io
CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
CURRENT_PATH=$(pwd)
PROJECT_ROOT_PATH=${CURRENT_PATH}/../..
BUILD_TARGET_PATH=${PROJECT_ROOT_PATH}/bin/paste_together

cd ${PROJECT_ROOT_PATH}/src/aaa.com/paste_together
go build -tags netgo -a -o ${BUILD_TARGET_PATH} aaa.com/paste_together
#go build -tags netgo -a -o ${BUILD_TARGET_PATH} main.go

cp -r ${PROJECT_ROOT_PATH}/src/aaa.com/paste_together/template ${PROJECT_ROOT_PATH}/bin/

cd ${CURRENT_PATH}

echo "Build finish..."
