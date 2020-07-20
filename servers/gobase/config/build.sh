#!/bin/bash

set -e

# Copy .env file
if [ ! -f "/app/${APP_NAME}/.env" ]; then
	echo "[build.sh: Copy .env for $APP_NAME]"
	cp /app/$APP_NAME/.env.example /app/$APP_NAME/.env
fi

#initilized go module
if [ ! -f "/app/${APP_NAME}/go.mod" ]; then
	echo "[build.sh: go mod init for $APP_NAME]"
	go mod init $APP_NAME
fi

#echo "[build.sh: Building binary for $APP_NAME]"
#cd $BUILDPATH && go build -o /servicebin
#echo "[build.sh: launching binary for $APP_NAME]"

# Run compiled service
go run /app/$APP_NAME/main.go "$@"