### 解决vue 入口文件 index.html 文件缓存问题
vue-cli打包出来的 js css 文件 本身是经过 hash 过的，所以文件不会存在缓存问题，浏览器缓存的是 服务器端的 index.html 这个文件，所以我只要不缓存 index.html 文件 nginx 配制如下
```js
location = /index.html {
    add_header Cache-Control "no-cache, no-store";
}
```