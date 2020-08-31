# Mac 安装 mongoDB

> Mac 安装mongodb 的方式可以选择 Homebrew 方式，这里是手动下载安装包 执行命令方式 
> 推荐一个 mongoDB 的一个客户端可视化工具 [NoSQLBooster for MongoDB](https://nosqlbooster.com/downloads)

### 1、下载 mongoDB 到电脑上
[mongoDB 下载链接](https://www.mongodb.com/download-center/community)
![image.png](https://i.loli.net/2019/09/22/HUgYERcWLpDBF9N.png)
下载到桌面 或者自己的文件夹中

### 2、解压压缩包
```sh
移动下载的包(这里选择自己下载包 所在的路径，我的是放在桌面上)
sudo mv /桌面/mongodb-macos-x86_64-4.2.0 /usr/local
解压
tar -zxvf mongodb-macos-x86_64-4.2.0.tgz
更换 文件夹名称
sudo mv mongodb-macos-x86_64-4.2.0 mongodb
添加到环境变量
export PATH=/usr/local/mongodb/bin:$PATH
```
### 3、创建数据库文件夹 data
```sh
cd /
mkdir data/db
```
### 4、启动 mongo
```sh
cd /usr/local/mongodb/bin
sudo ./mongod
```
![image.png](https://i.loli.net/2019/09/22/3Gdyur6Sc8CHNoB.png)
启动成功！
浏览器打开 http://localhost:27017/ 会看到 ` It looks like you are trying to access MongoDB over HTTP on the native driver port.` 说明mongo 正在运行，大功告成！
