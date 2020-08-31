
### 1、删除nginx，–purge包括配置文件

```Shell
sudo apt-get --purge remove nginx
```

### 2、自动移除全部不使用的软件包
```Shell
sudo apt-get autoremove
```
### 3、罗列出与nginx相关的软件
```Shell
dpkg --get-selections|grep nginx

执行结果
stephen@stephen-OptiPlex-390:~$ dpkg --get-selections|grep nginx
nginx                       install
nginx-common                    install
nginx-core                  install
```
### 4、删除1.3查询出与nginx有关的软件
```Shell
sudo apt-get --purge remove nginx
sudo apt-get --purge remove nginx-common
sudo apt-get --purge remove nginx-core
```
### 5、查看nginx正在运行的进程，如果有就kill掉
```Shell
ps -ef |grep nginx
```
看下nginx还有没有启动,一般执行完1后，nginx还是启动着的，如下：
```sh
stephen@stephen-OptiPlex-390:~$ ps -ef |grep nginx
root      7875  2317  0 15:02 ?        00:00:00 nginx: master process /usr/sbin/nginx
www-data  7876  7875  0 15:02 ?        00:00:00 nginx: worker process
www-data  7877  7875  0 15:02 ?        00:00:00 nginx: worker process
www-data  7878  7875  0 15:02 ?        00:00:00 nginx: worker process
www-data  7879  7875  0 15:02 ?        00:00:00 nginx: worker process
stephen   8321  3510  0 15:20 pts/0    00:00:00 grep --color=auto nginx
```

### kill nginx进程
```Shell
sudo kill  -9  7875 7876 7877 7879
```
### 全局查找与nginx相关的文件
```Shell
sudo  find  /  -name  nginx*
```
### 依依删除4列出的所有文件
```Shell
sudo rm -rf file
```
## 这样就彻底删除nginx了

# 再次重装
```Shell
sudo apt-get update
sudo apt-get install nginx
```
> 之后打开ip地址就可以访问到 ***Welcome to nginx!***


# 基本命令的使用***********
### 启动nginx：/etc/nginx/
```Shell
sudo nginx -c /etc/nginx/nginx.conf
```
### 停止nginx
```Shell
nginx -s stop
```
### 重新启动-(更改配制文件)
```Shell
nginx -s reload
```
###### * 如有不正确的欢迎评论指出