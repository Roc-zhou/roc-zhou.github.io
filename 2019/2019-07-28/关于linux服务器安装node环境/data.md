### 1、下载 node文件 [node](https://nodejs.org/zh-cn/)

### 2、下载下来的tar文件上传到服务器并且解压，然后通过建立软连接变为全局；
1、上传服务器可以是自己任意路径
2、解压上传（解压后的文件我这边将名字改为了nodejs，这个地方自己随意，只要在建立软连接的时候写正确就可以）
3、解压版本` tar -xvf node-v10.14.1-linux-x64.tar.xz`
4、更换个名字 `mv node-v10.14.1-linux-x64  nodejs `
5、确认一下nodejs下bin目录是否有node 和npm文件，如果有执行软连接，如果没有重新下载执行上边步骤

### 3、 建立软连接，变为全局
```sh
1、ln -s /app/software/nodejs/bin/npm /usr/local/bin/ 
2、ln -s /app/software/nodejs/bin/node /usr/local/bin/ 
前面的是 自己node路径
```
### 4、最后一步检验nodejs是否已变为全局
```sh
node -v 
v10.14.1
```