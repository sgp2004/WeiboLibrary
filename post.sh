curl -u $1:$2 -d "source=${3}&status=${4}" "http://api.weibo.com/2/statuses/update.json"

