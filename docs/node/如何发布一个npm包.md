![cover](https://user-images.githubusercontent.com/33691840/121772414-da74d980-cba7-11eb-8f2f-b09cb90a2136.jpg)
## 安装nodeJs 
[https://nodejs.org/en/](https://nodejs.org/en/)

## 写一个最简单的npm包
一个npm包最少包含两个文件，一个 `package.json` 、一个 `index.js` （入口文件）
### 1、创建项目
```shell
npm init
```
输入这个包的一些基本信息、关键字等等。

新建一个 `index.js` 文件
```javascript
// index.js
console.log('hello npm')
```
这样一个最简单的npm 包就完成了、下面我们去发布
### 2、注册npm账号
[https://www.npmjs.com](https://www.npmjs.com/) 

### 3、添加账户 
```shell
npm adduser 
填入自己的npm账户名、密码和邮箱即可
```

### 4、发布npm包
```shell
npm publish
```
ok，大功告成！这样我们自己的npm包就已经发布了
> 注意：我们修改了 npm 包后，要发布的时候 一定要修改自己的版本号，不然会 error，更新不成功！
