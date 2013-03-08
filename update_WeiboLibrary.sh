#!/bin/bash
. /etc/profile
. ~/.bash_profile
cd /data0/WeiboLibrary
#update files
env GIT_SSL_NO_VERIFY=true git pull
#kill process and restart it
id=`ps -ef|grep WeiboLibraryServer|awk '{print $2}'`
echo $id
kill -9 $id
id2=`ps -ef|grep go-build|awk '{print $2}'`
echo $id2
kill -9 $id2
nohup go run WeiboLibraryServer.go  > /tmp/weibo_library.log 2>&1 &
