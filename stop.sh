pid = `ps -ef|grep small-tools|grep -v grep|awk 'print $2'`
echo $pid
kill -9 $pid