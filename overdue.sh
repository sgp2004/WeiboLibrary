#!/bin/bash
. /etc/profile
. ~/.bash_profile
cd /data0/WeiboLibrary
go run overdue_info_poster.go
