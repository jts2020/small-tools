res=`curl -s http://127.0.0.1:17806/app/health`
if [ "{\"status\":\"up\"}" == "$res" ];then
    exit 0;
else
    echo "App no health.."
    exit 99
fi