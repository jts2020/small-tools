res=`ps |grep -e '^[zZ]'`
if [ -n "$res" ];then
    echo $res
    exit 99
fi
exit 0