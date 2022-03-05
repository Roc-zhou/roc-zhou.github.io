# Mac下使用docker安装nginx并打开html
### 拉去nginx 镜像
```shell
# 拉取latest
docker pull nginx
```
```shell
Using default tag: latest
latest: Pulling from library/nginx
33847f680f63: Pull complete
dbb907d5159d: Pull complete
8a268f30c42a: Pull complete
b10cf527a02d: Pull complete
c90b090c213b: Pull complete
1f41b2f2bf94: Pull complete
Digest: sha256:8f335768880da6baf72b70c701002b45f4932acae8d574dedfddaf967fc3ac90
Status: Downloaded newer image for nginx:latest
docker.io/library/nginx:latest
```
### 查看本地镜像
```shell
docker images
```
![image](https://user-images.githubusercontent.com/33691840/126733809-3374cfbe-0e2b-44eb-ac0b-916e931e6c7a.png)

#### 运行下nginx

```shell
docker run \
  -d \
  -p 127.0.0.1:8080:80 \
  --rm \
  --name nginx-demo \
  nginx
```

- -d：在后台运行
- -p ：容器的80端口映射到127.0.0.1:8080
- --rm：容器停止运行后，自动删除容器文件
- --name：容器的名字为 nginx-demo

![image](https://user-images.githubusercontent.com/33691840/126734051-36836068-02e0-42f9-a17f-45f598296b8e.png)
访问 [http://127.0.0.1:8080/](http://127.0.0.1:8080/) 就可以看到 nginx欢迎界面了

### 映射网页目录和配置文件
#### 映射网页目录
在桌面新建一个nginx目录，里面放一个 html文件夹。在html文件夹下 创建一个 html文件
```shell
docker container run \
  -d \
  -p 127.0.0.1:8080:80 \
  --rm \
  --name nginx-demo \
  --volume "$PWD/html":/usr/share/nginx/html \
  nginx
```
$PWD 在当前nginx目录下执行
然后 `docker ps`  查看

![image](https://user-images.githubusercontent.com/33691840/126734948-fc1f2e2e-6574-442f-97bd-5754071c85f4.png)

访问 [http://127.0.0.1:8080/](http://127.0.0.1:8080/) 就可以看到自己的html界面了

> 怎么知道nginx 的位置的呢？
```shell
# 进入容器
docker exec -it 容器id｜容器名字 /bin/bash
# 查看nginx位置
whereis nginx
```
![image](https://user-images.githubusercontent.com/33691840/127086352-af7b7945-a785-47f1-aea1-bdc85da001dd.png)

#### 拷贝配置
```shell
docker container cp nginx-demo:/etc/nginx .
```
把nginx-demo容器的/etc/nginx拷贝到当前目录。注意 后面的 .
当前目录 多了一个 nginx 目录 执行 `mv nginx conf` 修改为 conf
![image](https://user-images.githubusercontent.com/33691840/126735410-31f7ab4c-f8d9-4de5-a35e-c1039703faf8.png)

终止容器
```shell
docker container stop nginx-demo
```
![image](https://user-images.githubusercontent.com/33691840/126735471-faf5d924-27d2-4f89-b61d-f48e8842d56f.png)

#### 映射配置目录
```shell
docker container run \
  --rm \
  --name nginx-demo \
  --volume "$PWD/html":/usr/share/nginx/html \
  --volume "$PWD/conf":/etc/nginx \
  -p 127.0.0.1:8080:80 \
  -d \
  nginx
```
![image](https://user-images.githubusercontent.com/33691840/126742067-5da0328e-540b-4d56-be3f-deab147846b5.png)
此时在访问 [http://127.0.0.1:8080/](http://127.0.0.1:8080/) 就可以看到 自己的nginx配置了