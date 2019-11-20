#!/bin/bash

export PORT=8080
export MYSQL_URL="root@tcp(127.0.0.1:3306)/user_mgr?parseTime=true&sql_mode=ansi"
export SECRET_KEY=kjdsdaiidfk

go run ./server/server.go
