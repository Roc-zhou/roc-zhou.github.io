#### 1、全局安装 pm2
> 首先要安装node 

```sh
npm install -g pm2
```
#### 2、pm2 -v

```sh
假如 安装完pm2 后 使用命令 pm2 -v 发现会报错
Command 'pm2' not found, did you mean:
这是因为没有配置全局环境变量
```
#### 3、配置全局环境变量

```sh
cd /usr/local/bin 进入 (这是我的环境变量所在的位置)
ll 查看所有

drwxr-xr-x  2 root root 4096 Apr 30 03:14 ./
drwxr-xr-x 11 root root 4096 Mar  4 09:51 ../
lrwxrwxrwx  1 root root   25 Mar  4 10:04 node -> /etc/node/nodeJs/bin/node*
lrwxrwxrwx  1 root root   24 Mar  4 10:04 npm -> /etc/node/nodeJs/bin/npm*
lrwxrwxrwx  1 root root   37 Apr 30 03:14 pm2 -> /etc/node/nodeJs/lib/node_modules/pm2/

会发现已经有 pm2 存在了。但是这里的软连接是错误的，需要重新配置软连接

删除 pm2 软连接
mv /usr/local/bin/pm2 /tmp/

重新配置
ln -s /etc/node/nodeJs/bin/pm2 /usr/local/bin

（/etc/node/nodeJs/bin/pm2）是自己的node安装目录，node下面会有pm2 

执行命令 pm2 -v
大功告成！
```



