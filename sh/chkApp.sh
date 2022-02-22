res=`ps -ef|grep AppName|grep -v grep`
if [ -n "$res" ];then
    exit 0
else
    echo "AppName non-existent.."
    exit 99
fi