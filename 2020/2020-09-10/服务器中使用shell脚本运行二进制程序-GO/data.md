```shell
#!/bin/bash

APP_HOME=$(pwd)
APP_NAME=main

#使用说明，用来提示输入参数
check() {
	echo "use : sh startup.sh [start|stop|restart|status]"
	exit 1
}

#检查程序是否运行
is_exist() {
	pid=`ps -ef|grep $APP_NAME|grep -v grep|awk '{print $2}'`
	#存在 0 不存在 1
	if [ -z "${pid}" ]; then
		return "1"
	else
		return "0"
	fi
}

#获取各种状态
server_status() {
	cpu=`ps --no-heading --pid=$pid -o pcpu`
	rss=`ps --no-heading --pid=$pid -o rss`
	echo `date +'%Y-%m-%d %H:%M:%S'`
	echo "cpu：${cpu}"
	echo "rss：${rss} KB"
}

#启动
start() {
	is_exist
	if [ $? -eq "0" ]; then
		server_status
		echo "${APP_NAME} 已经运行，pid = ${pid}"
	else
		nohup ./$APP_NAME >/dev/null 2>&1 &
		echo "${APP_NAME} 服务启动成功！"
	fi
}

#停止
stop() {
	is_exist
	if [ $? -eq "0" ]; then
		kill -9 $pid
		echo "${APP_NAME} 服务关闭！"
	else
		echo "${APP_NAME} 服务已关闭！"
	fi
}

#状态
status() {
	is_exist
	if [ $? -eq "0" ]; then
		server_status
		echo "${APP_NAME} 服务正在运行，pid = ${pid}"
	else
		echo "${APP_NAME} 服务已关闭！"
	fi
}

#重启
restart() {
	stop
	start
}

case "$1" in
	"start" ) start
	;;
	"stop" ) stop
	;;
	"status" ) status
	;;
	"restart" ) restart
	;;
	*) check
	;;
esac

```

