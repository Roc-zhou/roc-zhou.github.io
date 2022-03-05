## nginx 配制
```shell
# 开启gzip
gzip on;
# 启用gzip压缩的最小文件，小于设置值的文件将不会压缩
gzip_min_length 1k;
# 设置压缩所需要的缓冲区大小
gzip_buffers 16 64k;
# 设置gzip压缩针对的HTTP协议版本
gzip_http_version 1.1;
# gzip 压缩级别，1-9，数字越大压缩的越好，也越占用CPU时间
gzip_comp_level 3;
gzip_types text/plain application/x-javascript application/javascript text/javascript text/css application/xml application/x-httpd-php image/jpeg image/gif image/png;
# 是否在http header中添加Vary: Accept-Encoding，建议开启    	
gzip_vary on;
```

将上面的配制添加到 server 上面重启 nginx

```she
nginx -s reload
```

## vue 配制gzip 
> (vue cli 2X) vue-cli3.x 配制 差不多相似

### 1、安装 compression-webpack-plugin 
```sh
# 安装最新版本 会出错，看网友有的是安装 1.1.12 版本 确实 可以成功！感谢
npm install --save-dev compression-webpack-plugin@1.1.12    
```
### 2、修改productionGzip 为true   /config/index.js 下
![image.png](../../../images/9419407-d1c6aa44faf20a87.png)
### 3、修改 build/webpack.prod.conf.js 文件 为 filename
![image.png](../../../images/9419407-2415ecc086de1ace.png)

### 4、npm run build  更新服务器文件
最后访问下站点可以看到 gzip 成功
![image.png](../../../images/9419407-ae1abd9a278b7d11.png)