官方文档
- [husky-github](https://github.com/typicode/husky)
- [husky-doc](https://typicode.github.io/husky/#/)
- [lint-staged](https://github.com/okonet/lint-staged)


安装依赖
```shell
yarn add husky lint-staged -D
```
在package.json 中添加一条 preinstall 脚本
```shell
{
  "script":{
    "prepare": "husky install"
  }
}
```
prepare 为 package.json 中的“scripts”属性中的脚本
当执行 `install` 的时候 它会在 `prepublish` 之后、`prepublishOnly`之前执行
[官方文档](https://docs.npmjs.com/cli/v8/using-npm/scripts)

手动执行一次 `yarn run prepare` 就会得到一个 `.husky` 目录

为git仓库添加钩子 `pre-commit` 执行 
> husky 支持所有git 所有钩子 https://git-scm.com/docs/githooks
```shell
npx husky add .husky/pre-commit "npx --no-install lint-staged"
```
.husky目录下生成一个pre-commit的脚本
```shell
#!/bin/sh 
. "$(dirname "$0")/_/husky.sh"

npx --no-install lint-staged
```
配置 `lint-staged` , 在 `package.json` 中添加下面的配置
```json
{
  "lint-staged": {
    "*.{js,vue,ts,jsx,tsx}": [
      xxx
      # "prettier --write",
      # "eslint --fix"
    ],
    "*.{html,css,less,scss,md}": [
      xxx
      # "prettier --write"
    ]
  }
}
```
这样提交代码之后，commit 的时候就回去检查放在暂存区的文件是否符合规范等等
