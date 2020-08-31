### 1、安装nodeJs （不多说）
### 2、写一个最简单的npm包
一个npm包最少包含两个文件，一个 `package.json` 、一个 `index.js` （入口文件）
我们生成项目
```sh
npm init
为什么不使用 -y, 我们要写一下这个包的一些基本信息，一些关键字等等
```
新建一个 `index.js` 文件
```js
console.log('hello npm')
```
这样一个最简单的npm 包就完成了、下面我们去发布
1、注册npm账号
[https://www.npmjs.com](https://www.npmjs.com/) 
2、添加账户 
```sh
npm adduser 
填入自己的npm账户名、密码和邮箱即可
```
3、发布npm包
```sh
npm publish
```
ok，大功告成！这样我们自己的npm包就已经发布了
> 注意：我们修改了 npm 包后，要发布的时候 一定要修改自己的版本号，不然会 error，更新不成功！
