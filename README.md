# small-tools
环境（Linux x64,Win x64...），可对服务器磁盘和CPU使用率指定阀值监控及支持自定义脚本(shell,cmd)并进行邮件告警
## 目录：
启动执行 sh start.sh

停止执行 sh stop.sh

config/conf.yml 配置文件，各项参数详见配置文件

sh/* 可存放监控脚本或配置文件指定路径

## 目前功能：
1. 指定多系统卷对应目录或磁盘使用率监控

2. 当前服务器CPU使用率监控

3. 脚本监控灵活扩展：

  * a.脚本监控-服务器僵尸进程监控
  
  * b.脚本监控-服务进程存在监控
  
  * c.脚本监控-服务健康检查
  
  * d.脚本自定义：注意，脚本需要赋予执行权限chmod +x xxx.sh
  
    开发要求：脚本中exit 0为成功不进行捕捉，其他值为失败，此时echo输出为邮件内容
    
4. 支持YML配置

5. 支持CRON表达式定时调度

6. 支持报警邮件发送
